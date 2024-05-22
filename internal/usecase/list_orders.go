package usecase

import (
	"github.com/fullcycle_cleanarch/internal/entity"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
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

func (c *ListOrderUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return []ListOrderOutputDTO{}, err
	}

	dto := make([]ListOrderOutputDTO, len(orders))
	for i := 0; i < len(orders); i++ {
		dto[i] = ListOrderOutputDTO{
			ID:         orders[i].ID,
			Price:      orders[i].Price,
			Tax:        orders[i].Tax,
			FinalPrice: orders[i].FinalPrice,
		}
	}
	return dto, nil
}
