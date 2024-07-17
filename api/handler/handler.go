package handler

import (
	"api/genproto/content"
	"api/genproto/itineraries"
	"api/genproto/story"
	"log/slog"
)

type Handler struct {
	Itinerary itineraries.ItinerariesClient
	Story     story.StoryClient
	Content   content.ContentClient
	Log       *slog.Logger
}
