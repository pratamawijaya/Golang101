package service

import (
	"context"
	"database/sql"
	"restapi/helper"
	"restapi/model/domain"
	"restapi/model/web"
	"restapi/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (s *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	s.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}
