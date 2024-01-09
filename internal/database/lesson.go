package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Lesson struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CourseID    string
	Content     string
}

func NewLesson(db *sql.DB) *Lesson {
	return &Lesson{
		db: db,
	}
}

func (l *Lesson) FindAll() ([]*Lesson, error) {
	rows, err := l.db.Query("SELECT * FROM lessons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*Lesson
	for rows.Next() {
		var lesson Lesson
		err := rows.Scan(&lesson.ID, &lesson.Name, &lesson.Description, &lesson.CourseID, &lesson.Content)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, &lesson)
	}

	return lessons, nil
}

func (l *Lesson) FindByCourseID(courseID string) ([]*Lesson, error) {
	rows, err := l.db.Query("SELECT * FROM lessons WHERE course_id = ?", courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*Lesson
	for rows.Next() {
		var lesson Lesson
		err := rows.Scan(&lesson.ID, &lesson.Name, &lesson.Description, &lesson.CourseID, &lesson.Content)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, &lesson)
	}

	return lessons, nil
}

func (l *Lesson) FindByID(id string) (*Lesson, error) {
	row := l.db.QueryRow("SELECT * FROM lessons WHERE id = ?", id)

	var lesson Lesson
	err := row.Scan(&lesson.ID, &lesson.Name, &lesson.Description, &lesson.CourseID, &lesson.Content)
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (l *Lesson) Create(name, description, courseID, content string) (*Lesson, error) {
	id := uuid.New().String()
	_, err := l.db.Exec("INSERT INTO lessons (id, name, description, course_id, content) VALUES (?, ?, ?, ?, ?)", id, name, description, courseID, content)
	if err != nil {
		return nil, err
	}

	return l.FindByID(id)
}
