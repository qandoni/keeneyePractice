package teachers_postgres_repository

import "github.com/qandoni/keeneyePractice/internal/core/domain"

type TeacherModel struct {
	ID          int
	Version     int
	UserID      int
	FIO         string
	PhoneNumber string
}

func teachersDomainsFromModels(teachers []TeacherModel) []domain.Teacher {
	teacherDomains := make([]domain.Teacher, len(teachers))
	for i, teacher := range teachers {
		teacherDomains[i] = domain.NewTeacher(
			teacher.ID,
			teacher.Version,
			teacher.UserID,
			teacher.FIO,
			teacher.PhoneNumber,
		)
	}
	return teacherDomains
}
