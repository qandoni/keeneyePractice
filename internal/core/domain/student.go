package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type Student struct {
	ID          int
	Version     int
	UserID      int
	GroupID     int
	FIO         string
	PhoneNumber string
}

func NewStudent(
	id int,
	version int,
	userID int,
	groupID int,
	fio string,
	phoneNumber string,
) Student {
	return Student{
		ID:          id,
		Version:     version,
		UserID:      userID,
		GroupID:     groupID,
		FIO:         fio,
		PhoneNumber: phoneNumber,
	}
}

func NewStudentUninitialized(
	userID int,
	groupID int,
	fio string,
	phoneNumber string,
) Student {
	return NewStudent(
		UninitializedID,
		UninitializedVersion,
		userID,
		groupID,
		fio,
		phoneNumber,
	)
}

func (s *Student) Validate() error {
	fioLen := len([]rune(s.FIO))
	if fioLen < 3 || fioLen > 100 {
		return fmt.Errorf("invalid `FIO` len: %d: %w", fioLen, core_errors.ErrInvalidArgument)
	}

	phoneNumberLen := len([]rune(s.PhoneNumber))
	if phoneNumberLen < 10 || phoneNumberLen > 15 {
		return fmt.Errorf("invalid `PhoneNumber` len: %d: %w",
			phoneNumberLen,
			core_errors.ErrInvalidArgument,
		)
	}
	re := regexp.MustCompile(`^\+[0-9]+$`)
	if !re.MatchString(s.PhoneNumber) {
		return fmt.Errorf(
			"invalid `PhoneNumber` format: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

type StudentPatch struct {
	FIO         Nullable[string]
	PhoneNumber Nullable[string]
}

func NewStudentPatch(
	fio Nullable[string],
	phoneNumber Nullable[string],
) StudentPatch {
	return StudentPatch{
		FIO:         fio,
		PhoneNumber: phoneNumber,
	}
}

func (p *StudentPatch) Validate() error {
	if p.FIO.Set && p.FIO.Value == nil {
		return fmt.Errorf("`FIO` cant be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (s *Student) ApplyPatch(patch StudentPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate student patch: %w", err)
	}

	tmp := *s
	if patch.FIO.Set {
		tmp.FIO = *patch.FIO.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = *patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched student: %w", err)
	}

	*s = tmp

	return nil
}
