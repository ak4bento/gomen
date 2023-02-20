package orm

import (
	"database/sql"
	"gomen/app/exceptions"
	"gomen/database"
	"gorm.io/gorm"
)

var (
	db database.PostgreSQL
)

type Model struct {
	ID        uint  `gorm:"primarykey" json:"id" form:"id"`
	CreatedAt int64 `json:"created_at" form:"created_at"`
	UpdatedAt int64 `json:"updated_at" form:"updated_at"`
	DeletedAt int64 `json:"deleted_at" form:"deleted_at"`

	//Connection
	conn       *gorm.DB `json:"-" form:"-" gorm:"-"`
	sql        *sql.DB  `json:"-" form:"-" gorm:"-"`
	customConn bool     `json:"-" form:"-" gorm:"-"`
}

func (m *Model) open() {
	m.conn, m.sql = db.Connect(nil)
}

func (m *Model) close() {
	exceptions.ConnectionRefused(m.sql.Close())
}
