package teachers_transport_http

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type TeacherDTOResponse struct {
	ID          int    `json:"id"`
	Version     int    `json:"version`
	UserID      int    `json:"user_id"`
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phone_number"`
}

func teacherDTOFromDomain(teacher domain.Teacher) TeacherDTOResponse {
	return TeacherDTOResponse{
		ID:          teacher.ID,
		Version:     teacher.Version,
		UserID:      teacher.UserID,
		FIO:         teacher.FIO,
		PhoneNumber: teacher.PhoneNumber,
	}
}

func teachersDTOFromDomains(teachers []domain.Teacher) []TeacherDTOResponse {
	teachersDTO := make([]TeacherDTOResponse, len(teachers))
	for i, teacher := range teachers {
		teachersDTO[i] = teacherDTOFromDomain(teacher)
	}
	return teachersDTO
}
