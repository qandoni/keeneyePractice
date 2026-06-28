package users_http_transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
	core_http_types "github.com/qandoni/keeneyePractice/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	Login    core_http_types.Nullable[string] `json:"login"`
	Password core_http_types.Nullable[string] `json:"password"`
	Role     core_http_types.Nullable[string] `json:"role" validate:"required"`
}

type PatchUserResponse UserDTOResponse

func (r *PatchUserRequest) Validate() error {
	if r.Login.Set {
		if r.Login.Value == nil {
			return fmt.Errorf("`Login` can't be NULL")
		}
		loginLen := len([]rune(*r.Login.Value))
		if loginLen < 3 || loginLen > 100 {
			return fmt.Errorf("`Login` must be between 3 and 100 symbols")
		}
	}
	if r.Password.Set {
		if r.Password.Value != nil {
			passwordLen := len([]rune(*r.Password.Value))
			if passwordLen < 6 || passwordLen > 30 {
				return fmt.Errorf("`Password` must be between 6 and 30 symbols")
			}
		}
	}

	return nil
}

func (h *UsersHTTPHandler) PatchUser(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get int path value",
		)
		return
	}
	var request PatchUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}
	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to patch user",
		)
		return
	}
	response := PatchUserResponse(userDTOFromDomain(userDomain))
	c.JSON(http.StatusOK, response)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	var role domain.Nullable[enum.Role]

	if request.Role.Set {
		if request.Role.Value != nil {
			parsed := enum.Role(*request.Role.Value)
			role = domain.Nullable[enum.Role]{
				Value: &parsed,
				Set:   true,
			}
		} else {
			role = domain.Nullable[enum.Role]{
				Value: nil,
				Set:   true,
			}
		}
	}

	return domain.NewUserPatch(
		request.Login.ToDomain(),
		request.Password.ToDomain(),
		role,
	)
}
