package registration_csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

func Parse(reader io.Reader) ([]Row, error) {
	csvReader := csv.NewReader(reader)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv readAll: %w", err)
	}

	if len(records) <= 1 {
		return nil, fmt.Errorf("csv is empty")
	}

	rows := make([]Row, 0, len(records)-1)

	for i := 1; i < len(records); i++ {
		record := records[i]

		if len(record) != 5 {
			return nil, fmt.Errorf(
				"line %d: expected 5 colums",
				i+1,
			)
		}

		role := enum.Role(record[3])
		switch role {
		case enum.RoleStudent, enum.RoleTeacher:
		default:
			return nil, fmt.Errorf(
				"line %d: invalid role %q",
				i+1,
				record[3],
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
