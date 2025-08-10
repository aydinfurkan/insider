package src

import (
	"insider/src/config"
	"insider/src/controller"
	"insider/src/db"
	"insider/src/service"
)

type Handler struct {
	MessageController *controller.MessageController
	ProbeController   *controller.ProbeController
	MessageDb         *db.MessageDb
	WebhookService    *service.WebhookService
	MessageService    *service.MessageService
	RedisService      *service.RedisService
}

func NewHandler(cfg *config.ConfigType) *Handler {
	db := db.NewMessageDb(cfg)

	rS := service.NewRedisService(cfg)
	wS := service.NewWebhookService(cfg, rS)
	mS := service.NewMessageService(db, wS)
	mC := controller.NewMessageController(db, mS)
	pC := controller.NewProbeController()

	return &Handler{
		MessageController: mC,
		ProbeController:   pC,
		MessageDb:         db,
		WebhookService:    wS,
		MessageService:    mS,
		RedisService:      rS,
	}
}
