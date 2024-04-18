package postgreDatabase

import "gorm.io/gorm"

type DBAdapter struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) *DBAdapter {
	return &DBAdapter{
		conn: conn,
	}
}
