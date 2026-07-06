package registration_http_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	registration_contracts "github.com/qandoni/keeneyePractice/internal/features/registration_requests/contracts"
	registration_csv "github.com/qandoni/keeneyePractice/internal/features/registration_requests/csv"
)

func (h *RegistrationRequestsHTTPHandler) Import(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.Error(err).SetMeta("failed to form file by key 'file'")
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.Error(err).SetMeta("failed to open file")
		return
	}
	defer file.Close()

	csvRows, err := registration_csv.Parse(file)
	if err != nil {
		c.Error(err).SetMeta("failed to parse data from file")
		return
	}

	rows := make([]registration_contracts.ImportRow, 0, len(csvRows))

	for _, row := range csvRows {
		rows = append(rows, registration_contracts.ImportRow{
			FIO:         row.FIO,
			Email:       row.Email,
			PhoneNumber: row.PhoneNumber,
			Role:        row.Role,
			Group:       row.Group,
		})
	}

	input := registration_contracts.ImportInput{
		Rows: rows,
	}

	err = h.service.Import(c.Request.Context(), input)
	if err != nil {
		c.Error(err).SetMeta("failed to import data")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "registration requests imported",
	})
}
