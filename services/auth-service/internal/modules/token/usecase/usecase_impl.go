// Code generated by candi v1.8.17.

package usecase

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"time"

	"monorepo/services/auth-service/internal/modules/token/domain"
	"monorepo/services/auth-service/pkg/shared"
	shareddomain "monorepo/services/auth-service/pkg/shared/domain"
	"monorepo/services/auth-service/pkg/shared/repository"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"pkg.agungdp.dev/candi/candihelper"
	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/codebase/factory/dependency"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

const (
	// TokenKey const
	TokenKey = "18608c7d-b319-0xc000165c80-0xc0000da000-11478e4e2650"
)

var (
	// ErrTokenFormat var
	ErrTokenFormat = errors.New("Invalid token format")
	// ErrTokenExpired var
	ErrTokenExpired = errors.New("Token is expired")
)

type tokenUsecaseImpl struct {
	cache interfaces.Cache

	repoMongo  *repository.RepoMongo
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// NewTokenUsecase usecase impl constructor
func NewTokenUsecase(deps dependency.Dependency) TokenUsecase {
	return &tokenUsecaseImpl{
		cache:      deps.GetRedisPool().Cache(),
		repoMongo:  repository.GetSharedRepoMongo(),
		publicKey:  deps.GetKey().PublicKey(),
		privateKey: deps.GetKey().PrivateKey(),
	}
}

// Generate token
func (uc *tokenUsecaseImpl) Generate(ctx context.Context, payload *domain.Claim) (resp domain.ResponseGenerateToken, err error) {
	trace := tracer.StartTrace(ctx, "TokenUsecase:Generate")
	defer trace.Finish()
	ctx = trace.Context()

	now := time.Now()

	savedToken := shareddomain.Token{DeviceID: payload.DeviceID, UserID: payload.User.ID}
	if err := uc.repoMongo.TokenRepo.Find(ctx, &savedToken); err == nil &&
		candihelper.PtrToBool(savedToken.IsActive) &&
		savedToken.ExpiredAt.After(now) {
		return domain.ResponseGenerateToken{
			Token:        savedToken.Token,
			RefreshToken: savedToken.RefreshToken,
			Claim:        savedToken.Claims,
		}, nil
	}

	ageDuration := shared.GetEnv().JWTAccessTokenAge

	exp := now.Add(ageDuration)
	payload.Id = uuid.NewString()

	var key interface{}
	var token = new(jwt.Token)
	if payload.Alg == domain.HS256 {
		token = jwt.New(jwt.SigningMethodHS256)
		key = []byte(TokenKey)
	} else {
		token = jwt.New(jwt.SigningMethodRS256)
		key = uc.privateKey
	}
	claims := jwt.MapClaims{
		"iss":  "mooc",
		"exp":  exp.Unix(),
		"iat":  now.Unix(),
		"did":  payload.DeviceID,
		"aud":  payload.Audience,
		"jti":  payload.Id,
		"sub":  payload.User.ID,
		"user": payload.User,
	}
	token.Claims = claims

	tokenString, err := token.SignedString(key)
	if err != nil {
		return resp, err
	}

	redisKey := candihelper.BuildRedisPubSubKeyTopic(
		domain.RedisTokenExpiredKeyConst,
		domain.RedisTokenExpiredKey{
			DeviceID: payload.DeviceID, UserID: payload.User.ID,
		},
	)
	uc.cache.Set(ctx, redisKey, candihelper.ToBytes(claims), ageDuration)

	refreshTokenHS := jwt.New(jwt.SigningMethodHS256)
	refreshTokenHS.Claims = jwt.MapClaims{
		"exp": now.Add(shared.GetEnv().JWTRefreshTokenAge).Unix(),
	}
	refreshTokenString, err := refreshTokenHS.SignedString([]byte(TokenKey))
	if err != nil {
		return resp, err
	}
	savedToken.RefreshToken = refreshTokenString

	savedToken.Token = tokenString
	savedToken.DeviceID = payload.DeviceID
	savedToken.UserID = payload.User.ID
	savedToken.ExpiredAt = exp
	savedToken.IsActive = candihelper.ToBoolPtr(true)
	savedToken.Claims = claims
	uc.repoMongo.TokenRepo.Save(ctx, &savedToken)

	resp.Token = tokenString
	resp.RefreshToken = savedToken.RefreshToken
	resp.Claim = claims
	return
}

// Refresh token
func (uc *tokenUsecaseImpl) Refresh(ctx context.Context, token, refreshToken string) (resp domain.ResponseGenerateToken, err error) {
	trace := tracer.StartTrace(ctx, "TokenUsecase:Refresh")
	defer trace.Finish()
	ctx = trace.Context()

	now := time.Now()

	refreshTokenParse, err := jwt.ParseWithClaims(refreshToken, &domain.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenKey), nil
	})
	switch ve := err.(type) {
	case *jwt.ValidationError:
		if ve.Errors == jwt.ValidationErrorExpired {
			err = errors.New("Refresh token is expired")
		} else {
			err = errors.New("Refresh token invalid")
		}
	}
	if err != nil {
		return
	}
	if !refreshTokenParse.Valid {
		return resp, errors.New("Refresh token invalid")
	}

	savedToken := shareddomain.Token{Token: token, RefreshToken: refreshToken}
	if err := uc.repoMongo.TokenRepo.Find(ctx, &savedToken); err != nil {
		return resp, errors.New("Token not found")
	}
	if savedToken.Token != token || savedToken.RefreshToken != refreshToken {
		return resp, errors.New("Token not found")
	}

	ageDuration := shared.GetEnv().JWTAccessTokenAge
	savedToken.Claims["exp"] = now.Add(ageDuration).Unix()
	savedToken.Claims["jti"] = uuid.NewString()

	var jwtToken = new(jwt.Token)
	jwtToken = jwt.New(jwt.SigningMethodRS256)
	jwtToken.Claims = jwt.MapClaims(savedToken.Claims)

	token, err = jwtToken.SignedString(uc.privateKey)
	if err != nil {
		return resp, err
	}

	refreshTokenHS := jwt.New(jwt.SigningMethodHS256)
	refreshTokenHS.Claims = jwt.MapClaims{
		"exp": now.Add(shared.GetEnv().JWTRefreshTokenAge).Unix(),
	}
	refreshTokenString, err := refreshTokenHS.SignedString([]byte(TokenKey))
	if err != nil {
		return resp, err
	}
	savedToken.RefreshToken = refreshTokenString
	savedToken.Token = token
	savedToken.IsActive = candihelper.ToBoolPtr(true)
	uc.repoMongo.TokenRepo.Save(ctx, &savedToken)

	redisKey := candihelper.BuildRedisPubSubKeyTopic(
		domain.RedisTokenExpiredKeyConst,
		domain.RedisTokenExpiredKey{
			DeviceID: savedToken.DeviceID, UserID: savedToken.UserID,
		},
	)
	uc.cache.Set(ctx, redisKey, candihelper.ToBytes(savedToken.Claims), ageDuration)

	resp.Token = token
	resp.RefreshToken = savedToken.RefreshToken
	resp.Claim = savedToken.Claims
	return
}

// Validate token
func (uc *tokenUsecaseImpl) Validate(ctx context.Context, tokenString string) (claim *domain.Claim, err error) {
	trace := tracer.StartTrace(ctx, "TokenUsecase:Validate")
	defer trace.Finish()
	ctx = trace.Context()

	tokenParse, err := jwt.ParseWithClaims(tokenString, &domain.Claim{}, func(token *jwt.Token) (interface{}, error) {
		checkAlg, _ := candishared.GetValueFromContext(ctx, candishared.ContextKey("tokenAlg")).(string)
		if checkAlg == domain.HS256 {
			return []byte(TokenKey), nil
		}
		return uc.publicKey, nil
	})

	switch ve := err.(type) {
	case *jwt.ValidationError:
		if ve.Errors == jwt.ValidationErrorExpired {
			err = ErrTokenExpired
		} else {
			err = ErrTokenFormat
		}
	}

	if err != nil {
		return
	}

	if !tokenParse.Valid {
		return claim, ErrTokenFormat
	}

	claim, _ = tokenParse.Claims.(*domain.Claim)

	redisKey := candihelper.BuildRedisPubSubKeyTopic(
		domain.RedisTokenExpiredKeyConst,
		domain.RedisTokenExpiredKey{
			DeviceID: claim.DeviceID, UserID: claim.User.ID,
		},
	)
	redisValue, errRedis := uc.cache.Get(ctx, redisKey)
	if errRedis != nil {
		userToken := shareddomain.Token{DeviceID: claim.DeviceID, UserID: claim.User.ID}
		if err := uc.repoMongo.TokenRepo.Find(ctx, &userToken); err != nil {
			return nil, ErrTokenExpired
		}
		if (userToken.IsActive == nil) || (userToken.IsActive != nil && !*userToken.IsActive) {
			return nil, ErrTokenExpired
		}
	}
	var redisClaim domain.Claim
	json.Unmarshal(redisValue, &redisClaim)
	if redisClaim.Id != claim.Id {
		return nil, errors.New("Invalid token")
	}

	return
}

// Revoke token
func (uc *tokenUsecaseImpl) Revoke(ctx context.Context, token string) error {
	trace := tracer.StartTrace(ctx, "TokenUsecase:Revoke")
	defer trace.Finish()
	ctx = trace.Context()

	claims, err := uc.Validate(ctx, token)
	if err != nil {
		return err
	}

	redisKey := candihelper.BuildRedisPubSubKeyTopic(
		domain.RedisTokenExpiredKeyConst,
		domain.RedisTokenExpiredKey{
			DeviceID: claims.DeviceID, UserID: claims.User.ID,
		},
	)

	uc.cache.Delete(ctx, redisKey)
	return uc.RevokeByKey(ctx, claims.DeviceID, claims.User.ID)
}

// RevokeByKey token
func (uc *tokenUsecaseImpl) RevokeByKey(ctx context.Context, deviceID, userID string) error {
	trace := tracer.StartTrace(ctx, "TokenUsecase:RevokeByKey")
	defer trace.Finish()
	ctx = trace.Context()

	data := shareddomain.Token{DeviceID: deviceID, UserID: userID}
	if err := uc.repoMongo.TokenRepo.Find(ctx, &data); err != nil {
		return err
	}

	data.IsActive = candihelper.ToBoolPtr(false)
	return uc.repoMongo.TokenRepo.Save(ctx, &data)
}
