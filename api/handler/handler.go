package handler

import (
	"api/genproto/content"
	"log/slog"
)

type Handler struct {
	Content content.ContentClient
	Log     *slog.Logger
}
