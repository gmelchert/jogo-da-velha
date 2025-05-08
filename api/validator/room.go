package validator

import (
	"fmt"
)

type FindRoomQuery struct {
	RoomID     string   `form:"roomId"`
	OwnerID    *uint    `form:"ownerId"`
	OpponentID *uint    `form:"opponentID"`
	Page       *float64 `form:"page"`
	Limit      *float64 `form:"limit"`
}

type CreateRoomRequest struct {
	RoomID     string `json:"roomId"`
	OwnerID    uint   `json:"ownerId"`
	OpponentID uint   `json:"opponentId"`
}

func (r *CreateRoomRequest) ValidateCreate() error {
	if r == nil {
		return fmt.Errorf("body da requisição está malformado")
	}
	if r.RoomID == "" {
		return errParamIsRequired("roomId", "string")
	}
	if r.OwnerID == 0 {
		return errParamIsRequired("ownerId", "uint")
	}
	if r.OpponentID == 0 {
		return errParamIsRequired("opponentId", "uint")
	}

	return nil
}
