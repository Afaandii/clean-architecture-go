package repository

import (
	"clean-architecture-go/internal/domain/entity"
	repo "clean-architecture-go/internal/domain/repository"
	"database/sql"
	"errors"
)

// Fungsi: implementasi CategoryRepository menggunakan PostgreSQL. Hanya file ini tahu SQL.
type categoryPGRepository struct {
	db *sql.DB
}

func NewCategoryPGRepository(db *sql.DB) repo.CategoryRepository {
	return &categoryPGRepository{db: db}
}

func (r *categoryPGRepository) FindAll() ([]entity.Category, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM categories ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []entity.Category
	for rows.Next() {
		var c entity.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

func (r *categoryPGRepository) FindByID(id int) (*entity.Category, error) {
	var c entity.Category
	row := r.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id)
	if err := row.Scan(&c.ID, &c.Name, &c.Description); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *categoryPGRepository) Create(c *entity.Category) error {
	err := r.db.QueryRow("INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id", c.Name, c.Description).Scan(&c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *categoryPGRepository) Update(c *entity.Category) error {
	res, err := r.db.Exec("UPDATE categories SET name = $1, description = $2 WHERE id = $3", c.Name, c.Description, c.ID)
	if err != nil {
		return err
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

func (r *categoryPGRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}
