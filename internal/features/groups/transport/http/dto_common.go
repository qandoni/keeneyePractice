package groups_http_transport

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type GroupsDTOResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func groupDTOFromDomain(group domain.Group) GroupsDTOResponse {
	return GroupsDTOResponse{
		ID:   group.ID,
		Name: group.Name,
	}
}

func groupsDTOFromDomains(groups []domain.Group) []GroupsDTOResponse {
	groupsDTO := make([]GroupsDTOResponse, len(groups))
	for i, group := range groups {
		groupsDTO[i] = groupDTOFromDomain(group)
	}
	return groupsDTO
}
