package inventory

import "errors"

type Service interface {
	AddInventory(inv *Inventory) error
	GetInventoryByID(id uint64) (*Inventory, error)
	UpdateInventory(inv *Inventory) error
	RemoveInventory(id uint64) error
	GetInventoriesByPlayer(playerID string) ([]Inventory, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) AddInventory(inv *Inventory) error {
	return s.repo.Create(inv)
}

func (s *service) GetInventoryByID(id uint64) (*Inventory, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateInventory(inv *Inventory) error {
	return s.repo.Update(inv)
}

func (s *service) RemoveInventory(id uint64) error {
	return s.repo.Delete(id)
}

func (s *service) GetInventoriesByPlayer(playerID string) ([]Inventory, error) {
	if playerID == "" {
		return nil, errors.New("playerID is required")
	}
	return s.repo.ListByPlayerID(playerID)
}
