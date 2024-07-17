package main

import (
	"api/api"
	"api/api/handler"
	"api/genproto/content"
	"api/genproto/itineraries"
	"api/genproto/story"
	"api/pkg/logger"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	hand := NewHandler()
	router := api.Router(hand)
	log.Printf("server is running...")
	log.Fatal(router.Run(":8080"))
}

func NewHandler() *handler.Handler {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	contenservice := content.NewContentClient(conn)
	itinerariesservice := itineraries.NewItinerariesClient(conn)
	storyservice := story.NewStoryClient(conn)
	return &handler.Handler{
		Content:   contenservice,
		Story:     storyservice,
		Itinerary: itinerariesservice,
		Log:       logger.NewLogger(),
	}
}
