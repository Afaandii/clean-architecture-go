package usecase

import (
	"clean-architecture-go/internal/domain/entity"
	"clean-architecture-go/internal/domain/repository"
	"errors"
)

// Fungsi: aplikasi/business rules. Menggunakan repository interface, bukan implementasi konkret.
var (
	ErrNotFound = errors.New("category not found")
)

type CategoryUsecase interface {
	GetAll() ([]entity.Category, error)
	GetByID(id int) (*entity.Category, error)
	Create(name string, description string) (*entity.Category, error)
	Update(id int, name string, description string) (*entity.Category, error)
	Delete(id int) error
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(r repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo: r}
}

func (u *categoryUsecase) GetAll() ([]entity.Category, error) {
	return u.repo.FindAll()
}

func (u *categoryUsecase) GetByID(id int) (*entity.Category, error) {
	c, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, ErrNotFound
	}
	return c, nil
}

func (u *categoryUsecase) Create(name string, description string) (*entity.Category, error) {
	c := &entity.Category{Name: name, Description: description}
	if err := u.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (u *categoryUsecase) Update(id int, name string, description string) (*entity.Category, error) {
	c := &entity.Category{ID: id, Name: name, Description: description}
	if err := u.repo.Update(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (u *categoryUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
