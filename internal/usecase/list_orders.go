package usecase

import (
	"github.com/renatonasc/20-CleanArch/internal/entity"
)

type ListOrderOutputDTO struct {
	Orders []*OrderOutputDTO `json:"orders"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() (ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return ListOrderOutputDTO{}, err
	}

	dto := ListOrderOutputDTO{
		Orders: make([]*OrderOutputDTO, len(orders)),
	}

	for i, order := range orders {
		dto.Orders[i] = &OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return dto, nil
}
