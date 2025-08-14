package itemshop

import "errors"

type Service interface {
	CreateItem(item *Item) error
	GetItemByID(id uint64) (*Item, error)
	UpdateItem(item *Item) error
	DeleteItem(id uint64) error
	GetAllItems() ([]Item, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateItem(item *Item) error {
	existing, err := s.repo.GetByID(item.ID)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("item with this ID already exists")
	}
	return s.repo.Create(item)
}

func (s *service) GetItemByID(id uint64) (*Item, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateItem(item *Item) error {
	return s.repo.Update(item)
}

func (s *service) DeleteItem(id uint64) error {
	return s.repo.Delete(id)
}

func (s *service) GetAllItems() ([]Item, error) {
	return s.repo.ListAll()
}
