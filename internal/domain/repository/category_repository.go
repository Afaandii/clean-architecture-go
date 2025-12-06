package repository

import "clean-architecture-go/internal/domain/entity"

// kontrak repository — domain hanya tahu interface ini
// Fungsi: mendefinisikan kontrak akses data. Implementasi ada di infrastructure.
type CategoryRepository interface {
	FindAll() ([]entity.Category, error)
	FindByID(id int) (*entity.Category, error)
	Create(c *entity.Category) error
	Update(c *entity.Category) error
	Delete(id int) error
}
