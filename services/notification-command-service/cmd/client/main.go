package main

import (
	"context"
	"log"
	"time"

	notificationpb "gitverse.ru/go-portfolio/notify-platform/notification-command-service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := notificationpb.NewNotificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SendNotification(ctx,
		&notificationpb.SendNotificationRequest{
			UserId:  "1",
			Type:    "email",
			Message: "Привет от GRPC-клиента",
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Идентификатор нотификации: ", resp.Id)
}
