package dto

type TaskDTO struct {
    ID          uint   `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}