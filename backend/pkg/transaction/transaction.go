package transaction

import "gorm.io/gorm"

func WithTransaction(
	db *gorm.DB,
	fn func(tx *gorm.DB) error,
) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
