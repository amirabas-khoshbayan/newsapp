package newsparam

import (
	"newsapp/entity"
	"time"
)

type AddToWaitingListRequest struct {
	UserID   uint            `json:"user_id"`
	Category entity.Category `json:"category"`
}

type AddToWaitingListResponse struct {
	Timeout time.Duration `json:"timeout_in_nano_seconds"`
}
