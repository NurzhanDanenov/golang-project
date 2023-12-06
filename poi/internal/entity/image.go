package entity

type Image struct {
	Id    int    `json:"id" db:"id"`
	Image string `json:"image" db:"image" binding:"required"`
}
