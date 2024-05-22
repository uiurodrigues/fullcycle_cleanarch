//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/fullcycle_cleanarch/internal/entity"
	"github.com/fullcycle_cleanarch/internal/event"
	"github.com/fullcycle_cleanarch/internal/infra/database"
	"github.com/fullcycle_cleanarch/internal/infra/web"
	"github.com/fullcycle_cleanarch/internal/usecase"
	"github.com/fullcycle_cleanarch/pkg/events"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrderUseCase(db *sql.DB) *usecase.ListOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrderUseCase,
	)
	return &usecase.ListOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
