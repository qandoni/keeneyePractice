package students_postgres_repository

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type StudentModel struct {
	ID          int
	Version     int
	UserID      int
	GroupID     int
	FIO         string
	PhoneNumber string
}

func studentDomainsFromModels(students []StudentModel) []domain.Student {
	studentDomains := make([]domain.Student, len(students))

	for i, student := range students {
		studentDomains[i] = domain.NewStudent(
			student.ID,
			student.Version,
			student.UserID,
			student.GroupID,
			student.FIO,
			student.PhoneNumber,
		)
	}
	return studentDomains
}
