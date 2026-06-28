package core_auth

import "github.com/qandoni/keeneyePractice/internal/core/enum"

type AuthInfo struct {
	UserID int
	Role   enum.Role
}
