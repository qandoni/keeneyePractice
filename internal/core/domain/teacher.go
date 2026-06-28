package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

type Teacher struct {
	ID          int
	Version     int
	UserID      int
	FIO         string
	PhoneNumber string
}

func (t *Teacher) Validate() error {
	fioLen := len([]rune(t.FIO))
	if fioLen < 3 || fioLen > 100 {
		return fmt.Errorf("invalid `FIO` len: %d: %w", fioLen, core_errors.ErrInvalidArgument)
	}
	phoneNumberLen := len([]rune(t.PhoneNumber))
	if phoneNumberLen < 10 || phoneNumberLen > 15 {
		return fmt.Errorf("invalid `PhoneNumber` len: %d: %w",
			phoneNumberLen,
			core_errors.ErrInvalidArgument,
		)
	}
	re := regexp.MustCompile(`^\+[0-9]+$`)
	if !re.MatchString(t.PhoneNumber) {
		return fmt.Errorf("invalid `PhoneNumber` format: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

func NewTeacher(
	id int,
	version int,
	userID int,
	fio string,
	phoneNumber string,
) Teacher {
	return Teacher{
		ID:          id,
		Version:     version,
		UserID:      userID,
		FIO:         fio,
		PhoneNumber: phoneNumber,
	}
}

func NewTeacherUninitialized(
	userID int,
	fio string,
	phoneNumber string,
) Teacher {
	return Teacher{
		ID:          UninitializedID,
		Version:     UninitializedVersion,
		UserID:      userID,
		FIO:         fio,
		PhoneNumber: phoneNumber,
	}
}

type TeacherPatch struct {
	FIO         Nullable[string]
	PhoneNumber Nullable[string]
}

func NewTeacherPatch(
	fio Nullable[string],
	phoneNumber Nullable[string],
) TeacherPatch {
	return TeacherPatch{
		FIO:         fio,
		PhoneNumber: phoneNumber,
	}
}

func (p *TeacherPatch) Validate() error {
	if p.FIO.Set && p.FIO.Value == nil {
		return fmt.Errorf("`FIO` cant be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (s *Teacher) ApplyPatch(patch TeacherPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate teacher patch: %w", err)
	}

	tmp := *s
	if patch.FIO.Set {
		tmp.FIO = *patch.FIO.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = *patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched teacher: %w", err)
	}

	*s = tmp

	return nil
}
