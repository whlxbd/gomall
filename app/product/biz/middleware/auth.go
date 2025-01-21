package middleware

import (
    "context"
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
    role, ok := ctx.Value(RoleKeyName).(string)
    if !ok || role != RoleAdmin {
        return &AuthError{Message: "permission denied: admin required"}
    }
    return nil
}