package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var category Category
	err := c.db.QueryRow("SELECT id, name, description FROM categories WHERE id = (SELECT category_id FROM courses WHERE id = $1)", courseID).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (c *Category) Update(id string, name string, description string) (Category, error) {
	_, err := c.db.Exec("UPDATE categories SET name = $1, description = $2 WHERE id = $3", name, description, id)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) Delete(id string) (bool, error) {
	_, err := c.db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	return true, nil
}
