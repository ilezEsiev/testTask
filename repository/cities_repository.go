package repository

import (
	"github.com/jmoiron/sqlx"
	"testTask/models"
)

type CitiesRepository struct {
	db *sqlx.DB
}

func NewCitiesRepository(db *sqlx.DB) *CitiesRepository {
	return &CitiesRepository{db: db}
}

func (r *CitiesRepository) Create(city models.City) error {
	query := `
		INSERT INTO cities (name, state) VALUES (?, ?)
	`
	_, err := r.db.Exec(query, city.Name, city.State)
	if err != nil {
		return err
	}

	return nil
}

func (r *CitiesRepository) Delete(id int) error {
	query := `
		DELETE FROM cities WHERE id = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CitiesRepository) Update(city models.City) error {
	query := `
		UPDATE cities SET name = ?, state = ? WHERE id = ?
	`

	_, err := r.db.Exec(query, city.Name, city.State, city.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CitiesRepository) List() ([]models.City, error) {
	var cities []models.City

	query := `
		SELECT * FROM cities
	`

	err := r.db.Select(&cities, query)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
