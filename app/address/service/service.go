package service

import (
	"fiber-crud/app/address/model"
	"fiber-crud/app/address/repository"
	"fiber-crud/pkg/config"
	"fiber-crud/pkg/enum"
	"fiber-crud/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Login(ctx *fiber.Ctx, authReq *model.Auth) (model.Response, error)
	Create(ctx *fiber.Ctx, auth *model.Auth) error
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return &service{repo}
}

func (s *service) Login(ctx *fiber.Ctx, authReq *model.Auth) (result model.Response, err error) {
	// Inquiry in db.
	auth, err := s.repo.Inquiry_Auth(authReq.UserName)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return result, fiber.NewError(fiber.StatusInternalServerError)
	}

	//## Validate password
	if !CheckPasswordHash(authReq.Password, auth) {
		return result, fiber.ErrUnauthorized
	}

	//## Check User in db was data.
	if auth.User.UserId == "" {
		return result, fiber.NewError(fiber.StatusUnauthorized, "User not found in database.")
	}

	result.Token, err = createToken(&auth.User)
	if err != nil {
		return result, fiber.NewError(fiber.StatusInternalServerError)
	}

	return result, nil
}

func (s *service) Create(ctx *fiber.Ctx, req *model.Auth) error {
	// Init value
	authId := uuid.New().String()
	userId := uuid.New().String()
	pass, err := HashPassword(req.Password + authId)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	auth := model.Auth{Id: authId, UserId: userId, UserName: req.UserName, Password: pass}
	user := model.User{UserId: userId, Name: req.UserName}

	// Call repository
	err = s.repo.Create_UserAndAuth(&auth, &user)
	if err != nil {
		utils.LogErrCtx(ctx, err.Error())
		return err
	}

	return nil
}

func createToken(user *model.User) (string, error) {
	// Get config
	_config := config.Server()
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	transactionId := uuid.New().String()
	claims[enum.USER_INFO] = utils.JsonSerialize(&utils.UserInfo{
		TransactionId: transactionId,
		UserId:        user.UserId,
	})
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(_config.Token_Expire)).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(_config.Secret_Key))
	if err != nil {
		return t, fiber.ErrUnauthorized
	}

	return t, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// func CheckPasswordHash(password string, hash string) bool {
func CheckPasswordHash(passReq string, auth model.Auth) bool {
	return bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(passReq+auth.Id)) == nil
}
