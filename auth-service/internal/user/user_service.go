package user

import (
	"chat/util"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	secretKey = "secret"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	exists, _ := s.Repository.EmailExists(ctx, req.Email)
	if !exists {
		hashedPassword, err := util.HashPassword(req.Password)

		if err != nil {
			return nil, err
		}

		u := &User{
			UserName: req.UserName,
			Email:    req.Email,
			Password: hashedPassword,
		}

		r, err := s.Repository.CreateUser(ctx, u)
		if err != nil {
			return nil, err
		}

		res := &CreateUserRes{
			ID:       r.ID,
			UserName: r.UserName,
			Email:    r.Email,
		}

		return res, nil
	}

	return nil, util.ErrUserExists("user already exists")
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		{
			return &LoginUserRes{}, err
		}
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       u.ID,
		Username: u.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    u.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{accessToken: ss, UserName: u.UserName, ID: u.ID}, nil
}
