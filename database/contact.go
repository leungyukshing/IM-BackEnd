package database

import (
	"github.com/backend/database/entities"
)

func GetContactIDsByUserID(userID string) ([]int64, error) {
	var contacts []entities.Contact
	var result []int64
	db := Server.Raw("SELECT user_id2 FROM contact WHERE user_id1 = ? UNION ALL SELECT user_id1 FROM contact WHERE user_id2 = ?", userID, userID).Find(&contacts)
	if db.Error != nil {
		return result, db.Error
	}

	for _, c := range contacts {
		result = append(result, c.UserID2)
	}
	return result, nil
}