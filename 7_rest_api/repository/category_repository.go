package repository

import (
	"restapi/model/domain"
	"database/sql"
	"context"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category) 
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
	