package usecase

import (
	"context"

	"github.com/janainamai/learning-go/8-uow/internal/entity"
	"github.com/janainamai/learning-go/8-uow/internal/repository"
)

type (
	InputUseCase struct {
		CategoryName     string
		CourseName       string
		CourseCategoryID int
	}

	AddCourseUseCase struct {
		CourseRepository   repository.CourseRepositoryInterface
		CategoryRepository repository.CategoryRepository
	}
)

func NewAddCourseUseCase(courseRepository repository.CourseRepositoryInterface, categoryRepository repository.CategoryRepository) *AddCourseUseCase {
	return &AddCourseUseCase{
		CourseRepository:   courseRepository,
		CategoryRepository: categoryRepository,
	}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
