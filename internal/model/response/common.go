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

type GetHealthCheckMemoryResp struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"totalAlloc"`
	Sys        uint64 `json:"sys"`
	HeapAlloc  uint64 `json:"heapAlloc"`
	HeapSys    uint64 `json:"heapSys"`
}

type GetHealthCheckServiceStatusResp struct {
	Name string `json:"name"`
	IsUp bool   `json:"isUp"`
}

type GetHealthCheckResp struct {
	Status          string                            `json:"status"`
	Environtment    string                            `json:"environtment"`
	Version         string                            `json:"version"`
	GoVersion       string                            `json:"goVersion"`
	GoRoutine       int                               `json:"goRoutine"`
	Memory          GetHealthCheckMemoryResp          `json:"memory"`
	ServiceStatuses []GetHealthCheckServiceStatusResp `json:"serviceStatuses"`
}
