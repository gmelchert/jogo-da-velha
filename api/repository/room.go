package repository

import (
	"fmt"
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
	if query.Status != "" {
		tx = tx.Where("status = ?", query.Status)
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

	if err := tx.Offset(offset).Limit(int(limit)).Preload("Owner").Preload("Opponent").Find(&rooms).Error; err != nil {
		return err
	}

	r.Rows = rooms
	r.TotalRows = totalRows
	r.TotalPages = totalPages
	r.Page = page
	r.Limit = limit

	return nil
}

func CreateRoom(req *validator.CreateRoomPayload) (models.Room, error) {
	room := models.Room{
		RoomID:     req.RoomID,
		OwnerID:    req.OwnerID,
		OpponentID: 0,
		Status:     req.Status,
	}

	if err := Db.Create(&room).Error; err != nil {
		return models.Room{}, err
	}

	return room, nil
}

func UpdateRoomStatus(req *validator.UpdateRoomRequest) error {
	result := Db.Model(&models.Room{}).Where("id = ?", req.ID).Updates(map[string]string{"status": req.Status})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows updated")
	}
	return nil
}

func DeleteRoom(roomID string) error {
	result := Db.Where("room_id = ?", roomID).Delete(&models.Room{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows deleted")
	}
	return nil
}

func JoinRoom(roomID string, opponentID uint) error {
	result := Db.Model(&models.Room{}).Where("room_id = ? AND status = 'OPEN'", roomID).Updates(map[string]interface{}{"opponent_id": opponentID, "status": "RUNNING"})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows updated")
	}
	return nil
}

func CloseRoom(roomID string, ownerID uint) error {
	result := Db.Model(&models.Room{}).Where("room_id = ? AND owner_id = ?", roomID, ownerID).Updates(map[string]string{"status": "CLOSED"})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows updated")
	}
	return nil
}
