package validator

import (
	"fmt"
)

type FindRoomQuery struct {
	RoomID     string   `form:"roomId"`
	OwnerID    *uint    `form:"ownerId"`
	OpponentID *uint    `form:"opponentId"`
	Page       *float64 `form:"page"`
	Limit      *float64 `form:"limit"`
	Status     string   `json:"status"`
}

type CreateRoomPayload struct {
	RoomID     string `json:"roomId"`
	OwnerID    uint   `json:"ownerId"`
	OpponentID uint   `json:"opponentId"`
	Status     string `json:"status"`
}

type CreateRoomRequest struct {
	RoomID string `json:"roomId"`
}

type UpdateRoomRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (r *CreateRoomRequest) ValidateCreate() error {
	if r == nil {
		return fmt.Errorf("body da requisição está malformado")
	}
	if r.RoomID == "" {
		return errParamIsRequired("roomId", "string")
	}

	return nil
}
