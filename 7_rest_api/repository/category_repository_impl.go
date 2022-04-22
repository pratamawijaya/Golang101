package repository

import (
	"restapi/helper"
	"restapi/model/domain"
	"database/sql"
	"context"
)

type CategoryRepositoryImpl struct {

}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	
	helper.PanicIfError(err)
	
	id, err := result.LastInsertId()
	
	helper.PanicIfError(err)

	category.Id = id

	return category
}