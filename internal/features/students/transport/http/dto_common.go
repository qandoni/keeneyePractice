package students_transport_http

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type StudentDTOResponse struct {
	ID           int    `json:"id"`
	Version      int    `json:"version`
	FIO          string `json:"fio"`
	StudentGroup string `json:"student_group"`
	PhoneNumber  string `json:"phone_number"`
}

func studentDTOFromDomain(student domain.Student) StudentDTOResponse {
	return StudentDTOResponse{
		ID:           student.ID,
		Version:      student.Version,
		FIO:          student.FIO,
		StudentGroup: student.StudentGroup,
		PhoneNumber:  student.PhoneNumber,
	}
}
