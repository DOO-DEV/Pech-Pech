package usecase

import (
	"context"
	"errors"
	"github.com/doo-dev/pech-pech/infrastructure/cache"
	"github.com/doo-dev/pech-pech/infrastructure/mail"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/internal/modules/auth/presenter"
	authRepo "github.com/doo-dev/pech-pech/internal/modules/auth/repository"
	userRepo "github.com/doo-dev/pech-pech/internal/modules/users/repository"
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/helper"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
)

const (
	TokenExpirationTime = 4 * time.Hour
	SigningKey          = "hello"
)

type Config struct {
	JwtSigningKey                string        `koanf:"jwt_signing_key"`
	JwtPrefix                    string        `koanf:"jwt_prefix"`
	AccessTokenExpTimeoutInHours time.Duration `koanf:"access_token_exp_timeout_in_hours"`
	OtpLength                    int           `koanf:"otp_length"`
	OtpExpTimeoutInSeconds       time.Duration `koanf:"otp_exp_timeout_in_seconds"`
}

type AuthClaims struct {
	userID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthService struct {
	userRepo userRepo.UserRepository
	authRepo authRepo.AuthRepository
	mail     mail.IMail
	cfg      Config
}

func NewAuthService(cfg Config, uRepo userRepo.UserRepository, autRepo authRepo.AuthRepository, m mail.IMail) AuthService {
	return AuthService{
		userRepo: uRepo,
		authRepo: autRepo,
		mail:     m,
		cfg:      cfg,
	}
}

func (a AuthService) Register(ctx context.Context, dto *presenter.RegisterRequest) (*presenter.RegisterResponse, error) {
	fmtUsername := strings.ToLower(dto.Username)

	if user, err := a.userRepo.GetUserByIdOrUsername(ctx, fmtUsername); user != nil {
		// TODO - implement rich error to find inner cause of error
		// TODO - add log
		if user != nil {
			return nil, constants.ErrUserExisted
		}

		if !errors.Is(err, constants.ErrNoRecord) {
			return nil, err
		}
	}
	hashedPassword, err := helper.Encrypt(dto.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		ID:       uuid.New().String(),
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
		Email:    user.Email,
	}, nil
}

func (a AuthService) Login(ctx context.Context, dto *presenter.LoginRequest) (*presenter.LoginResponse, error) {
	user, err := a.userRepo.GetUserByIdOrUsername(ctx, dto.Username)
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
			Issuer:    userId,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * a.cfg.AccessTokenExpTimeoutInHours)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(a.cfg.JwtSigningKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
func (a AuthService) ParseToken(tokenStr string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.JwtSigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func (a AuthService) ForgetPassword(_ context.Context, email string) error {
	otp := helper.RandomNumber(a.cfg.OtpLength)

	newMail := &mail.Mail{
		To:      email,
		Subject: "Reset Pech-Pech password",
		Body:    "Your OTP code is: " + otp,
	}

	cache.SetCache(cache.MailOtpKey(email), otp, time.Second*a.cfg.OtpExpTimeoutInSeconds)
	return a.mail.SendingMail(newMail)
}

func (a AuthService) ResetPassword(ctx context.Context, dto *presenter.ResetPasswordRequest) error {
	cacheKey := cache.MailOtpKey(dto.Email)
	sentOtp := cache.GetCache(cacheKey)

	if sentOtp == nil {
		// TODO - log this
		return constants.ErrOtpExpired
	}

	if sentOtp.(string) != dto.Code {
		return constants.ErrOtpInvalid
	}

	cache.RemoveFromCache(cacheKey)

	if dto.Password != dto.ConfirmPassword {
		return constants.ErrPasswordAndConfirmPasswordMatch
	}

	hashedPass, err := helper.Encrypt(dto.Password)
	if err != nil {
		return constants.ErrHashPassword
	}

	return a.authRepo.UpdatePassword(ctx, dto.Email, hashedPass)
}
func (a AuthService) UpdatePassword(ctx context.Context, dto *presenter.UpdatePasswordRequest, email string) error {
	hashedPass, err := helper.Encrypt(dto.Password)
	if err != nil {
		return constants.ErrHashPassword
	}
	return a.authRepo.UpdatePassword(ctx, email, hashedPass)
}
