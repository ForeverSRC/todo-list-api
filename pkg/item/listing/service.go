package listing

import "fmt"

type Service interface {
	ListItems(*ItemListQuery) (ItemList, error)
}

type Repository interface {
	FetchItems(*ItemListQuery) (ItemList, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListItems(q *ItemListQuery) (ItemList, error) {
	if q.Uid == "" {
		return nil, fmt.Errorf("user not login")
	}

	if err := q.CheckAndFix(); err != nil {
		return nil, err
	}

	l, err := s.repo.FetchItems(q)
	if err != nil {
		return nil, err
	}

	return l, nil

}
