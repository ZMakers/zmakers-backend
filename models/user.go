package models

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
	"zmakers-backend/service"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `gorm:"-"`
	Hash     string    `json:"hash"`
	Token    string    `json:"token"`
}

func (u *User) CreateBefore() error {
	if u.Id == uuid.Nil {
		u.Id = uuid.New()
	}

	hash := crypto.Keccak256([]byte(u.Password))
	u.Hash = common.Bytes2Hex(hash)
	return nil
}

func (u *User) CreatUser() error {
	err := u.CreateBefore()
	if err != nil {
		return err
	}
	result := service.Db.Create(u)
	if result.Error != nil || result.RowsAffected <= 0 {
		return result.Error
	}
	return nil
}

func (u *User) SearchUser() *User {
	hash := common.Bytes2Hex(crypto.Keccak256([]byte(u.Password)))
	result := service.Db.Where("name = ? && hash = ?", u.Name, hash).Find(u)
	if result.RowsAffected > 0 {
		return u
	}
	return nil
}

func (u *User) UpdateUserToken() *User {
	res := service.Db.Select("token").Updates(u)
	if res.RowsAffected > 0 || res.Error != nil {
		return u
	}
	return nil
}

var mySigningKey = []byte("ZMAKERS")

func (u *User) GenerateToken() error {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: u.Name,
		Subject: u.Password,
		NotBefore: jwt.NewNumericDate(time.Now().Add(-1*time.Hour)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1*time.Hour)),
	})

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return err
	}
	u.Token = ss
	return nil
}

func (u *User) VerifyToken(token string)  error {
	if token == "" {
		return errors.New("empty token")
	}

	_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("ZMAKERS"), nil
	})

	if _token.Valid {
		return nil
	}
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("Not Even A Token ")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return errors.New(" Either expired or too early ")
		}
	}
	return errors.New(" Couldn't handle this token ")
}

