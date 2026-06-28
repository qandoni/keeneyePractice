package students_transport_http

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type StudentDTOResponse struct {
	ID          int    `json:"id"`
	Version     int    `json:"version`
	UserID      int    `json:"user_id"`
	GroupID     int    `json:"group_id"`
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phone_number"`
}

func studentDTOFromDomain(student domain.Student) StudentDTOResponse {
	return StudentDTOResponse{
		ID:          student.ID,
		Version:     student.Version,
		UserID:      student.UserID,
		GroupID:     student.GroupID,
		FIO:         student.FIO,
		PhoneNumber: student.PhoneNumber,
	}
}

func studentsDTOFromDomains(students []domain.Student) []StudentDTOResponse {
	studentsDTO := make([]StudentDTOResponse, len(students))
	for i, student := range students {
		studentsDTO[i] = studentDTOFromDomain(student)
	}
	return studentsDTO
}
