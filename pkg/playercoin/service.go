package playercoin

import "errors"

type Service interface {
	CreateCoin(coin *PlayerCoin) error
	GetCoinByPlayerID(playerID string) (*PlayerCoin, error)
	UpdateCoin(coin *PlayerCoin) error
	DeleteCoin(id uint64) error
	GetAllCoins() ([]PlayerCoin, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCoin(coin *PlayerCoin) error {
	existing, err := s.repo.GetByPlayerID(coin.PlayerID)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("coin record already exists for this player")
	}
	return s.repo.Create(coin)
}

func (s *service) GetCoinByPlayerID(playerID string) (*PlayerCoin, error) {
	return s.repo.GetByPlayerID(playerID)
}

func (s *service) UpdateCoin(coin *PlayerCoin) error {
	return s.repo.Update(coin)
}

func (s *service) DeleteCoin(id uint64) error {
	return s.repo.Delete(id)
}

func (s *service) GetAllCoins() ([]PlayerCoin, error) {
	return s.repo.ListAll()
}
