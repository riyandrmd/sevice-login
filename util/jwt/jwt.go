package jwt

import (
	"time"
	"users/authenticate"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func CreateToken(refID int) (*authenticate.TokenDetail, error) {
	td := &authenticate.TokenDetail{}
	timeNow := time.Now()
	td.AtExpires = timeNow.Add(time.Minute * 15).Unix()
	uidAt, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	uidRt, err := uuid.NewRandom()

	td.AccessUuid = uidAt.String()
	td.RefreshUuid = uidRt.String()

	acClaims := authenticate.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    authenticate.From,
			ExpiresAt: td.AtExpires,
			NotBefore: timeNow.Unix(),
			Id:        uidAt.String(),
			IssuedAt:  timeNow.Unix(),
		},
		UID: refID,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		acClaims,
	)

	signedAc, err := token.SignedString([]byte(authenticate.AuthSecret))
	if err != nil {
		return nil, err
	}
	td.AccessToken = signedAc

	rfClaims := authenticate.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    authenticate.From,
			NotBefore: timeNow.Unix(),
			Id:        uidRt.String(),
			IssuedAt:  timeNow.Unix(),
		},
	}

	tokenRf := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		rfClaims,
	)

	signedRf, err := tokenRf.SignedString([]byte(authenticate.RefreshAuthSecret))
	if err != nil {
		return nil, err
	}

	td.RefreshToken = signedRf
	return td, err
}
