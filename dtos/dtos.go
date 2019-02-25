package dtos

type CreateTaskDto struct {
	Title string `json:"title"`
}

type UpdateTaskDto struct {
	NewTitle string `json:"newTitle"`
}

type ErrorResponseDto struct {
	Code  int    `json:"code"`
	Error string `json:"message"`
}
