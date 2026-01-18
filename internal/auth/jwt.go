package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey string
	ExpiresIn time.Duration
}

// UserClaims 用户JWT声明
type UserClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// NewJWTConfig 创建JWT配置
func NewJWTConfig(secretKey string, expiresIn time.Duration) *JWTConfig {
	return &JWTConfig{
		SecretKey: secretKey,
		ExpiresIn: expiresIn,
	}
}

// GenerateToken 生成访问令牌（Access Token）
func (j *JWTConfig) GenerateToken(userID int64, email, name string) (string, error) {
	// 创建声明
	claims := UserClaims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken 生成刷新令牌（Refresh Token）
// 刷新令牌使用更长的过期时间
func (j *JWTConfig) GenerateRefreshToken(userID int64, email, name string, refreshExpiresIn time.Duration) (string, error) {
	// 创建声明
	claims := UserClaims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析JWT令牌
func (j *JWTConfig) ParseToken(tokenString string) (*UserClaims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌有效性
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// IsValidToken 检查令牌是否有效
func (j *JWTConfig) IsValidToken(tokenString string) bool {
	_, err := j.ParseToken(tokenString)
	return err == nil
}
