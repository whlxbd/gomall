package middleware

import (
    "context"
    "github.com/golang-jwt/jwt"
    "os"
    "fmt"
)

const (
    RoleKeyName = "user_role"
    RoleAdmin   = "admin"
)


// AuthError 自定义权限错误
type AuthError struct {
    Message string
}

func (e *AuthError) Error() string {
    return e.Message
}

// CheckAdminPermission 检查管理员权限
func CheckAdminPermission(ctx context.Context) error {
    tokenValue := ctx.Value("token")
    if tokenValue == nil {
        return &AuthError{Message: "no token found"}
    }

    token, ok := tokenValue.(string)
    if !ok {
        return &AuthError{Message: "invalid token type"}
    }

    claims := jwt.MapClaims{}
    _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SCRETE")), nil
    })
    if err != nil {
        return &AuthError{Message: fmt.Sprintf("parse token failed: %v", err)}
    }

    if role, ok := claims["role"].(string); !ok || role != RoleAdmin {
        return &AuthError{Message: "user is not admin"}
    }
    
    return nil
}