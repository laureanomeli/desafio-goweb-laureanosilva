package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	allTickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	ticketsByDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return float64(len(ticketsByDestination)) / float64(len(allTickets)), nil
}
