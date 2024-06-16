package response

const (
	MessageOK = "ok"
)

type CustomError struct {
	Code       string
	Message    string
	StatusCode int
}

func (m CustomError) Error() string {
	return m.Message
}

func (m CustomError) ToResponse() BaseResponse {
	return BaseResponse{
		StatusCode: m.StatusCode,
		Message:    m.Message,
	}
}

type BaseResponse struct {
	StatusCode       int               `json:"-"`
	Message          string            `json:"message"`
	Data             any               `json:"data"`
	Error            any               `json:"error"`
	ValidationErrors []ValidationError `json:"validationErrors"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

func NewResponseOK() *BaseResponse {
	return &BaseResponse{
		Message: MessageOK,
	}
}
