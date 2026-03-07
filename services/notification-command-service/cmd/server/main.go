package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/handler"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/repository"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/internal/service"
	"gitverse.ru/go-portfolio/notify-platform/notification-command-service/pkg/kafka"
	notificationpb "gitverse.ru/go-portfolio/notify-platform/notification-command-service/proto"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	connString := "postgres://notify:notify@localhost:5432/notify?sslmode=disable"
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("Невозможно подключититься к пулу соединений: %v", err)
	}
	defer pool.Close()

	repo := repository.NewRepository(pool)
	producer := kafka.NewProducer([]string{"localhost:9092"}, "notifications")

	svc := service.NewService(*repo, *producer)
	h := handler.NewHandler(svc)

	grpcServer := grpc.NewServer()
	notificationpb.RegisterNotificationServiceServer(grpcServer, h)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Сервис нотификаций запущен на порту :50051")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
