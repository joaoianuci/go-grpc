package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) FindAll() ([]*Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]*Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	return courses, nil
}

func (c *Course) FindByID(id string) (*Course, error) {
	row := c.db.QueryRow("SELECT * FROM courses WHERE id = ?", id)

	var course Course
	err := row.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *Course) Create(name string, description string, categoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)", id, name, description, categoryID)
	if err != nil {
		return nil, err
	}

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) FindByLessonID(lessonID string) (*Course, error) {
	var course Course
	err := c.db.QueryRow("SELECT id, name, description, category_id FROM courses WHERE id = (SELECT course_id FROM lessons WHERE id = $1)", lessonID).Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
	if err != nil {
		return nil, err
	}

	return &course, nil
}
