package registration_csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/mail"
	"reflect"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
	core_errors "github.com/qandoni/keeneyePractice/internal/core/errors"
)

func Parse(reader io.Reader) ([]Row, error) {
	csvReader := csv.NewReader(reader)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv readAll: %w", err)
	}

	if len(records) <= 1 {
		return nil, fmt.Errorf("csv len: %w", core_errors.ErrEmptyFile)
	}

	rows := make([]Row, 0, len(records)-1)
	expectedHeader := []string{"fio", "email", "phone_number", "role", "group_name"}

	if !reflect.DeepEqual(records[0], expectedHeader) {
		return nil, fmt.Errorf("invalid csv header: %w", core_errors.ErrInvalidArgument)
	}

	for i := 1; i < len(records); i++ {
		record := records[i]

		if len(record) != 5 {
			return nil, fmt.Errorf(
				"line %d: expected 5 colums: %w",
				i+1,
				core_errors.ErrInvalidArgument,
			)
		}
		_, err := mail.ParseAddress(record[1])
		if err != nil {
			return nil, fmt.Errorf(
				"line %d: invalid email %q: %w",
				i+1,
				record[1],
				core_errors.ErrInvalidArgument,
			)
		}

		role := enum.Role(record[3])
		switch role {
		case enum.RoleStudent, enum.RoleTeacher:
		default:
			return nil, fmt.Errorf(
				"line %d: %q: %w",
				i+1,
				record[3],
				core_errors.ErrInvalidArgument,
			)
		}

		rows = append(rows, Row{
			FIO:         record[0],
			Email:       record[1],
			PhoneNumber: record[2],
			Role:        role,
			Group:       record[4],
		})
	}
	return rows, nil
}
