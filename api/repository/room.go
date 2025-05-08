package repository

import (
	"math"

	"github.com/gmelchert/jogo-da-velha/api/models"
	"github.com/gmelchert/jogo-da-velha/api/validator"
)

func FindRoom(r *models.PaginatedResponse, query *validator.FindRoomQuery) error {
	var (
		page  float64 = 1
		limit float64 = 20
	)

	rooms := []models.Room{}

	tx := Db.Model(&models.Room{})

	if query.RoomID != "" {
		tx = tx.Where("room_id = ?", query.RoomID)
	}
	if query.OpponentID != nil {
		tx = tx.Where("opponent_id = ?", *query.OpponentID)
	}
	if query.OwnerID != nil {
		tx = tx.Where("owner_id = ?", *query.OwnerID)
	}

	if query.Page != nil {
		page = *query.Page
	}
	if query.Limit != nil {
		limit = *query.Limit
	}

	var totalRows int64

	tx.Count(&totalRows)

	totalPages := math.Ceil(float64(totalRows) / float64(*query.Limit))
	offset := int((page - 1) * (limit))

	if err := tx.Offset(offset).Limit(int(limit)).Find(&rooms).Error; err != nil {
		return err
	}

	r.Rows = rooms
	r.TotalRows = totalRows
	r.TotalPages = totalPages
	r.Page = page
	r.Limit = limit

	return nil
}

func CreateRoom(req *validator.CreateRoomRequest) (models.Room, error) {
	room := models.Room{
		RoomID:     req.RoomID,
		OwnerID:    req.OwnerID,
		OpponentID: req.OpponentID,
	}

	if err := Db.Create(&room).Error; err != nil {
		return models.Room{}, err
	}

	return room, nil
}
