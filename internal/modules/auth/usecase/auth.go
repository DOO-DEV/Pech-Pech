package usecase

import (
	"context"
	"fmt"
	"github.com/doo-dev/pech-pech/infrastructure/cache"
	"github.com/doo-dev/pech-pech/infrastructure/mail"
	"github.com/doo-dev/pech-pech/internal/models"
	"github.com/doo-dev/pech-pech/internal/modules/auth/presenter"
	authRepo "github.com/doo-dev/pech-pech/internal/modules/auth/repository"
	userRepo "github.com/doo-dev/pech-pech/internal/modules/users/repository"
	"github.com/doo-dev/pech-pech/pkg/constants"
	"github.com/doo-dev/pech-pech/pkg/helper"
	"github.com/doo-dev/pech-pech/pkg/richerror"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
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
	const op = "authservice.Register"

	fmtUsername := strings.ToLower(dto.Username)

	if _, err := a.userRepo.GetUserByIdOrUsername(ctx, fmtUsername); err != nil {
		// TODO - add log
		return nil, richerror.New(op).WithError(err)
	}
	hashedPassword, err := helper.Encrypt(dto.Password)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	user := &models.User{
		ID:       uuid.New().String(),
		Username: fmtUsername,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	if err := a.authRepo.CreateUser(ctx, user); err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	return &presenter.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (a AuthService) Login(ctx context.Context, dto *presenter.LoginRequest) (*presenter.LoginResponse, error) {
	const op = "authservice.Login"

	user, err := a.userRepo.GetUserByIdOrUsername(ctx, dto.Username)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	if err := helper.Decrypt(dto.Password, user.Password); err != nil {
		return nil, richerror.New(op).WithError(err)
	}
	// We can read token from a in app memory cache but I'm not sure is it good
	token, err := a.createToken(user.ID, user.Username, user.Email)
	if err != nil {
		return nil, richerror.New(op).WithError(err)
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
	const op = "authservice.createToken"

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
		return "", richerror.New(op).WithError(err)
	}

	return tokenStr, nil
}
func (a AuthService) ParseToken(tokenStr string) (*AuthClaims, error) {
	const op = "authService.ParseToken"

	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.JwtSigningKey), nil
	})
	if err != nil {
		fmt.Println("parsing error", err)
		return nil, richerror.New(op).WithError(err)
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, richerror.New(op).WithError(err)
}

func (a AuthService) ForgetPassword(_ context.Context, email string) error {
	const op = "authservice.ForgetPassword"

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
	const op = "authservice.ResetPassword"

	cacheKey := cache.MailOtpKey(dto.Email)
	sentOtp := cache.GetCache(cacheKey)

	if sentOtp == nil {
		// TODO - log this
		richerror.New(op).WithMessage(constants.ErrMsgOtpExpired)
	}

	if sentOtp.(string) != dto.Code {
		return richerror.New(op).WithMessage(constants.ErrMsgOtpInvalid)
	}

	cache.RemoveFromCache(cacheKey)

	if dto.Password != dto.ConfirmPassword {
		return richerror.New(op).WithMessage(constants.ErrMsgPasswordAndConfirmPasswordMatch)
	}

	hashedPass, err := helper.Encrypt(dto.Password)
	if err != nil {
		return richerror.New(op).WithMessage(constants.ErrMsgHashPassword)
	}

	return a.authRepo.UpdatePassword(ctx, dto.Email, hashedPass)
}
func (a AuthService) UpdatePassword(ctx context.Context, dto *presenter.UpdatePasswordRequest, email string) error {
	const op = "authservice.UpdatePassword"

	hashedPass, err := helper.Encrypt(dto.Password)
	if err != nil {
		return richerror.New(op).WithMessage(constants.ErrMsgHashPassword)
	}
	return a.authRepo.UpdatePassword(ctx, email, hashedPass)
}
