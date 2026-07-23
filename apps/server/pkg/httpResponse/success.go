package httpResponse

type Meta struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	Total     int64 `json:"total"`
	TotalPage int64 `json:"totalPage"`
}

type Success struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta        `json:"meta,omitempty"`
}
