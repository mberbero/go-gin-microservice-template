package post

import (
	"github.com/mberbero/go-microservice-template/pkg/dtos"
	"github.com/mberbero/go-microservice-template/pkg/entities"
)

// Service is a service interface
type Service interface {
	Create(post *entities.Post) error
	Get(id string) (*entities.Post, error)
	GetAll(page, perPage int64) (*dtos.PaginatedData, error)
	Update(id string, post *dtos.PostDTO) error
	Delete(id string) error
}

type service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(post *entities.Post) error {
	return s.repository.Create(post)
}

func (s *service) Get(id string) (*entities.Post, error) {
	return s.repository.Get(id)
}

func (s *service) GetAll(page, perPage int64) (*dtos.PaginatedData, error) {
	return s.repository.GetAll(page, perPage)
}

func (s *service) Update(id string, post *dtos.PostDTO) error {
	return s.repository.Update(id, post)
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(id)
}
