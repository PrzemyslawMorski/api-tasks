package dtos

type CreateTaskDto struct {
	Title string `json:"title"`
}

type DeleteTaskDto struct {
	Id int `json:"id"`
}

type UpdateTaskDto struct {
	NewTitle string `json:"newTitle"`
}

type ReadTaskDto struct {
	Id int `json:"id"`
}

type ErrorResponseDto struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
