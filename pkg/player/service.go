package player

import "errors"

type Service interface {
	CreatePlayer(player *Player) error
	GetPlayerByID(id string) (*Player, error)
	UpdatePlayer(player *Player) error
	DeletePlayer(id string) error
	GetAllPlayers() ([]Player, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreatePlayer(player *Player) error {
	exisiting, err := s.repo.GetByID(player.ID)
	if err != nil {
		return err
	}

	if exisiting != nil {
		return errors.New("player with this ID already exists")
	}

	return s.repo.Create(player)
}

func (s *service) GetPlayerByID(id string) (*Player, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdatePlayer(player *Player) error {
	return s.repo.Update(player)
}

func (s *service) DeletePlayer(id string) error {
	return s.repo.Delete(id)
}

func (s *service) GetAllPlayers() ([]Player, error) {
	return s.repo.ListAll()
}
