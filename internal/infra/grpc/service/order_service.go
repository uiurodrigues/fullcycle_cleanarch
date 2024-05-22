package service

import (
	"context"

	"github.com/fullcycle_cleanarch/internal/infra/grpc/pb"
	"github.com/fullcycle_cleanarch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(ctx context.Context, in *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
	orderList, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	output := make([]*pb.Order, len(orderList))
	for i := 0; i < len(orderList); i++ {
		ord := orderList[i]
		output[i] = &pb.Order{
			Id:         ord.ID,
			Price:      float32(ord.Price),
			Tax:        float32(ord.Tax),
			FinalPrice: float32(ord.FinalPrice),
		}
	}
	return &pb.ListOrderResponse{
		Orders: output,
	}, nil
}
