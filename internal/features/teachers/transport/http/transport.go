package teachers_transport_http

type TeachersHTTPHandler struct {
	teachersService TeachersService
}

type TeachersService interface{}

func NewTeachersHTTPHandler(
	teachersService TeachersService,
) *TeachersHTTPHandler {
	return &TeachersHTTPHandler{
		teachersService: teachersService,
	}
}
