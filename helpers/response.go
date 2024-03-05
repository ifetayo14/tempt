package helpers

type JSONResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}
