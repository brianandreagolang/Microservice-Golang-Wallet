package server

import (
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"

	"github.com/Marlos-Rodriguez/go-postgres-wallet-back/user/internal/cache"
	internal "github.com/Marlos-Rodriguez/go-postgres-wallet-back/user/internal/storage"
	"github.com/Marlos-Rodriguez/go-postgres-wallet-back/user/models"
	"github.com/Marlos-Rodriguez/go-postgres-wallet-back/user/storage"
)

//Server User Server struct
type Server struct {
}

var storageService storage.IUserStorageService

//GetStorageService Start the storage service for GPRC server
func GetStorageService() {
	var DB *gorm.DB = internal.ConnectDB()
	var RDB *redis.Client = cache.NewRedisClient()

	storageService = storage.NewUserStorageService(DB, RDB)
}

//CloseDB Close both DB
func CloseDB() {
	storageService.CloseDB()
}

//CheckUser check if the user exits
func (s *Server) CheckUser(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	if len(request.ID) < 0 || request.ID == "" {
		return &UserResponse{Exits: false, Active: false}, errors.New("Must send a ID")
	}

	exits, isActive, err := storageService.CheckExistingUser(request.ID)

	if err != nil {
		return &UserResponse{Exits: false, Active: false}, err
	}

	return &UserResponse{Exits: exits, Active: isActive}, nil
}

//CheckRelation Check if exits a Relation
func (s *Server) CheckRelation(ctx context.Context, request *RelationRequest) (*RelationResponse, error) {
	if len(request.FromUsername) < 0 || request.FromUsername == "" && len(request.ToUsername) < 0 || request.ToUsername == "" {
		return &RelationResponse{Exits: false}, errors.New("Must send ID")
	}

	exits, err := storageService.CheckExistingRelation(request.FromUsername, request.ToUsername, false)

	if err != nil {
		return &RelationResponse{Exits: false}, err
	}

	return &RelationResponse{Exits: exits}, nil
}

//ChangeAvatar Change the avatar in DB
func (s *Server) ChangeAvatar(ctx context.Context, request *AvatarRequest) (*AvatarResponse, error) {
	if len(request.ID) < 0 || request.Url == "" {
		return &AvatarResponse{Sucess: false}, errors.New("Must send the avatar name")
	}

	var userDB *models.User = new(models.User)

	userID, err := uuid.Parse(request.ID)

	if err != nil {
		return &AvatarResponse{Sucess: false}, err
	}

	userDB.Profile.Avatar = request.Url
	userDB.Profile.UserID = userID

	if sucess, err := storageService.ModifyUser(userDB, "", ""); sucess == false || err != nil {
		return &AvatarResponse{Sucess: false}, err
	}

	return &AvatarResponse{Sucess: true}, nil
}

//CheckUsersTransactions Check if the transaction is valid
func (s *Server) CheckUsersTransactions(ctx context.Context, request *CheckTransactionRequest) (*TransactionResponse, error) {
	//Check Input
	if request.FromID == "" || request.ToID == "" || request.Amount <= 0 || request.Password == "" {
		return &TransactionResponse{
			Exits:   false,
			Actives: false,
			Enough:  false,
		}, errors.New("No valid Input")
	}
	//Get user info
	fromUser, err := storageService.GetUser(request.FromID)

	if fromUser == nil || err != nil {
		return &TransactionResponse{
			Exits:   false,
			Actives: false,
			Enough:  false,
		}, err
	}

	//Check Password
	if success, err := storageService.CheckPassword(request.FromID, request.Password); !success || err != nil {
		return &TransactionResponse{
			Exits:   false,
			Actives: false,
			Enough:  false,
		}, err
	}

	//Check Balance
	if fromUser.Balance <= 0 || fromUser.Balance < request.Amount {
		return &TransactionResponse{
			Exits:   false,
			Actives: false,
			Enough:  false,
		}, errors.New("User not have enough balance")
	}

	//Check Active user
	toUser, err := storageService.GetUser(request.ToID)

	if toUser == nil || err != nil {
		return &TransactionResponse{
			Exits:   false,
			Actives: false,
			Enough:  false,
		}, err
	}

	return &TransactionResponse{
		Exits:   true,
		Actives: true,
		Enough:  true,
	}, nil
}

//MakeTransaction Between Users
func (s *Server) MakeTransaction(ctx context.Context, request *TransactionRequest) (*NewTransactionResponse, error) {
	if request.FromID == "" || request.ToID == "" || request.Amount <= 0 {
		return &NewTransactionResponse{Sucess: false}, errors.New("No valid Input")
	}

	fromUser, err := storageService.GetUser(request.FromID)

	if fromUser == nil || err != nil {
		return &NewTransactionResponse{Sucess: false}, err
	}

	if fromUser.Balance <= 0 || fromUser.Balance < request.Amount {
		return &NewTransactionResponse{Sucess: false}, errors.New("User not have enough balance")
	}

	toUser, err := storageService.GetUser(request.ToID)

	if toUser == nil || err != nil {
		return &NewTransactionResponse{Sucess: false}, err
	}

	var fromUserDB *models.User = new(models.User)
	fromUserDB.Balance = (fromUser.Balance - float64(request.Amount))

	var toUserDB *models.User = new(models.User)
	toUserDB.Balance = (toUser.Balance - float64(request.Amount))

	if success, err := storageService.ModifyUser(fromUserDB, request.FromID, ""); !success || err != nil {
		return &NewTransactionResponse{Sucess: false}, err
	}
	if success, err := storageService.ModifyUser(toUserDB, request.ToID, ""); !success || err != nil {
		return &NewTransactionResponse{Sucess: false}, err
	}

	storageService.UpdateUserCache(fromUser.UserID.String())
	storageService.UpdateUserCache(toUser.UserID.String())

	return &NewTransactionResponse{Sucess: true}, nil
}
