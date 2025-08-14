package purchasehistory

type Service interface {
	RecordPurchase(history *PurchaseHistory) error
	GetHistoryByID(id uint64) (*PurchaseHistory, error)
	GetPlayerHistory(playerID string) ([]PurchaseHistory, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) RecordPurchase(history *PurchaseHistory) error {
	return s.repo.Create(history)
}

func (s *service) GetHistoryByID(id uint64) (*PurchaseHistory, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetPlayerHistory(playerID string) ([]PurchaseHistory, error) {
	return s.repo.ListByPlayerID(playerID)
}
