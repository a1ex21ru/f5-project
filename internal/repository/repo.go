package repository

import (
	"gorm.io/gorm"

	"f5-project/internal/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) Create(n *models.Note) error {
	return s.db.Create(n).Error
}

func (s *Repository) Delete(id int) error {
	return s.db.Delete(&models.Note{}, "id = ?", id).Error
}

func (s *Repository) Update(id int, n *models.UpdateNote) error {
	return s.db.Model(&models.UpdateNote{}).Where("id = ?", id).Updates(n).Error
}

func (s *Repository) GetByID(ID int) (*models.Note, error) {

	var note models.Note

	if err := s.db.First(&note, "id = ?", ID).Error; err != nil {
		return nil, err
	}

	return &note, nil
}

func (s *Repository) GetAll() ([]models.Note, error) {
	var notes []models.Note

	if err := s.db.Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}
