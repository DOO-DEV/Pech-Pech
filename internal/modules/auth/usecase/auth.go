package usecase

import (
	"context"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/internal/modules/auth/presenter"
	authRepo "github.com/doo-dev/pech-pech/internal/modules/auth/repository"
	userRepo "github.com/doo-dev/pech-pech/internal/modules/users/repository"
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/helper"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

const (
	TokenExpirationTime = 4 * time.Hour
	SigningKey          = "hello"
)

type AuthClaims struct {
	userID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthService struct {
	userRepo userRepo.UserRepository
	authRepo authRepo.AuthRepository
}

func NewAuthService(uRepo userRepo.UserRepository, autRepo authRepo.AuthRepository) AuthService {
	return AuthService{
		userRepo: uRepo,
		authRepo: autRepo,
	}
}

func (a AuthService) Register(ctx context.Context, dto *presenter.RegisterRequest) (*presenter.RegisterResponse, error) {
	fmtUsername := strings.ToLower(dto.Username)

	if _, err := a.userRepo.GetUserByIdOrUsername(ctx, fmtUsername); err != nil {
		// TODO - implement rich error to find inner cause of error
		// TODO - add log
		return nil, constants.ErrUserExisted
	}

	hashedPassword, err := helper.Encrypt(dto.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Username: fmtUsername,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	if err := a.authRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &presenter.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Password,
	}, nil
}

func (a AuthService) Login(ctx context.Context, dto *presenter.LoginRequest) (*presenter.LoginResponse, error) {
	user, err := a.userRepo.GetUserByIdOrUsername(ctx, dto.UsernameOrEmail)
	if err != nil {
		return nil, err
	}

	if err := helper.Decrypt(dto.Password, user.Password); err != nil {
		return nil, err
	}
	// We can read token from a in app memory cache but I'm not sure is it good
	token, err := a.createToken(user.ID, user.Username, user.Email)
	if err != nil {
		return nil, err
	}
	res := &presenter.LoginResponse{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return res, nil

}
func (a AuthService) createToken(userId string, username, email string) (string, error) {
	claims := &AuthClaims{
		userID:   userId,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: userId,
			// TODO - add exp timeout to config
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpirationTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// TODO - add sign key and sign method to config
	tokenStr, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
func (a AuthService) ParseToken(tokenStr string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func (a AuthService) ForgetPassword() {}
func (a AuthService) VerifyOtp()      {}
func (a AuthService) ResetPassword()  {}
func (a AuthService) UpdatePassword() {}
