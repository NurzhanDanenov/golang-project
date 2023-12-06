package repository

import (
	"fmt"
	"restapi/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UploadImagePostgres struct {
	db *sqlx.DB
}

func NewUploadImagePostgres(db *sqlx.DB) *UploadImagePostgres {
	return &UploadImagePostgres{db: db}
}

func (r *UploadImagePostgres) UploadImage(userId int, image entity.Image) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var imageId int
	queryImage := fmt.Sprintf("INSERT INTO %s (image) values($1) RETURNING id", ImageTable)
	row := tx.QueryRow(queryImage, image.Image)
	if err := row.Scan(&imageId); err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	//if err := row.Scan(&id); err != nil {
	//	return 0, err
	//}
	createUsersImagesQuery := fmt.Sprintf("INSERT INTO %s (user_id, image_id) VALUES ($1, $2)", UsersImagesTable)
	_, err = tx.Exec(createUsersImagesQuery, userId, imageId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return imageId, tx.Commit()
}

func (r *UploadImagePostgres) GetAll(userId int) ([]entity.Image, error) {
	var images []entity.Image
	query := fmt.Sprintf("SELECT ti.id, ti.image FROM %s ti INNER JOIN %s ui on ti.id = ui.image_id WHERE ui.user_id = $1", ImageTable, UsersImagesTable)
	err := r.db.Select(&images, query, userId)

	return images, err
}

func (r *UploadImagePostgres) GetById(userId, imageId int) (entity.Image, error) {
	var image entity.Image
	query := fmt.Sprintf("SELECT ti.id, ti.image FROM %s ti INNER JOIN %s ui on ti.id = ui.image_id WHERE ui.user_id = $1 AND ui.image_id = $2", ImageTable, UsersImagesTable)
	err := r.db.Get(&image, query, userId, imageId)

	return image, err
}

func (r *UploadImagePostgres) Delete(userId, imageId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s ui WHERE ti.id = ui.image_id AND ui.user_id = $1 AND ui.image_id = $2", ImageTable, UsersImagesTable)
	_, err := r.db.Exec(query, userId, imageId)

	return err
}
