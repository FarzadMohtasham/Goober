package main

import (
	"context"
	"log"
	"time"

	"github.com/FarzadMohtasham/Goober/services/trip-service/internal/domain"
	"github.com/FarzadMohtasham/Goober/services/trip-service/internal/infrastructure/repository"
	"github.com/FarzadMohtasham/Goober/services/trip-service/internal/service"
)

func main() {
	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID: "42",
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	// keep program running for now
	for {
		time.Sleep(time.Second)
	}
}
