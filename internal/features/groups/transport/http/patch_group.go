package groups_http_transport

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandoni/keeneyePractice/internal/core/domain"
	core_http_request "github.com/qandoni/keeneyePractice/internal/core/transport/http/request"
	core_http_response "github.com/qandoni/keeneyePractice/internal/core/transport/http/response"
	core_http_types "github.com/qandoni/keeneyePractice/internal/core/transport/http/types"
)

type PatchGroupRequest struct {
	Name core_http_types.Nullable[string] `json:"Name"`
}

func (r *PatchGroupRequest) Validate() error {
	if r.Name.Set {
		if r.Name.Value == nil {
			return fmt.Errorf("`Name` can't be NULL")
		}
		nameLen := len([]rune(*r.Name.Value))
		if nameLen < 3 || nameLen > 10 {
			return fmt.Errorf("`Name` must be between 3 and 10 symbols")
		}
	}
	return nil
}

type PatchGroupResponse GroupsDTOResponse

func (h *GroupsHTTPHandler) PatchGroup(c *gin.Context) {
	ctx := c.Request.Context()

	groupID, err := core_http_request.GetIntPathValue(c, "id")
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to get int path value",
		)
		return
	}
	var request PatchGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	groupPatch := groupPatchFromRequest(request)

	groupDomain, err := h.groupsService.PatchGroup(ctx, groupID, groupPatch)
	if err != nil {
		core_http_response.RespondError(
			c,
			err,
			"failed to patch group",
		)
		return
	}
	response := PatchGroupResponse(groupDTOFromDomain(groupDomain))
	c.JSON(http.StatusOK, response)
}

func groupPatchFromRequest(request PatchGroupRequest) domain.GroupPatch {
	return domain.NewGroupPatch(
		request.Name.ToDomain(),
	)
}
