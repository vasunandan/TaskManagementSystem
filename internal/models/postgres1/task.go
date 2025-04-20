package postgres1

type Task struct {
    ID          uint   `gorm:"primaryKey"`
    Title       string
    Description string
    Status      string
}