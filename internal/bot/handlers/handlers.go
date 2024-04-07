package handlers

import (
	"github.com/FM1337/ASB/internal/bot/memory"
	"github.com/FM1337/ASB/internal/models"
	"github.com/apex/log"
)

type Handlers struct {
	logger log.Interface
	config *models.ConfigDiscord
	data   *memory.Data
}

func InitHandlers(lg log.Interface, config *models.ConfigDiscord, data *memory.Data) *Handlers {
	return &Handlers{
		config: config,
		data:   data,
	}
}
