package service

import (
	"errors"
	"log"
	"strconv"
	"time"

	"f5-project/internal/models"
	"f5-project/internal/repository"
)

type NoteService struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(n *models.Note) error {
	if n.Name == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(n)
}

func (s *NoteService) Update(id int, n *models.Note) error {
	if n.Name == "" {
		return errors.New("can't update with an empty name")
	}

	updateNote := models.UpdateNote{
		Name:        n.Name,
		Description: n.Description,
		UpdatedAt:   time.Now(),
	}

	return s.repo.Update(id, &updateNote)
}

func (s *NoteService) Delete(id string) error {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(idInt)
}
func (s *NoteService) GetByID(id string) (*models.Note, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	note, err := s.repo.GetByID(idInt)
	if err != nil {
		log.Printf("service get note by id err:%v", err)
		return nil, err
	}

	return note, nil
}

func (s *NoteService) GetAll() ([]models.Note, error) {
	return s.repo.GetAll()
}
