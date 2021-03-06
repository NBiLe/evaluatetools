package db

import (
	"time"

	"github.com/go-xorm/xorm"
)

const (
	KeyAccount     = "TAccount"
	KeyAAStatics   = "TStaticsActiveAccount"
	KeyMainAccount = "TMainAccount"
	QtGetRecord    = iota + 1
	QtCheckRecord
	QtQuaryAllRecord
	QtAddRecord
	QtAddRecords
	QtUpdateRecord
	QtDeleteRecord
	QtSeachRecord
	QtClearAllRecord
	QtGetCount
	QtGetCountRecords
	QtCommand

	MySqlDriver    DatabaseType = "mysql"
	SqliteDriver   DatabaseType = "sqlite3"
	PostgresDriver DatabaseType = "postgres"
)

// DatabaseType 数据库类型
type DatabaseType string

// DatabaseInfo 数据库基本信息定义
type DatabaseInfo struct {
	DbType    DatabaseType
	AliasName string
	Host      string
	Port      string
	UserName  string
	Password  string
	IsDebug   bool
}

// OperationInterface 数据库接口定义
type OperationInterface interface {
	Init(e *xorm.Engine)
	GetKey() string
	Quary(qtype int, v ...interface{}) error
}

// TMainAccount 账户定义
type TMainAccount struct {
	ID         uint64 `xorm:"'id' pk autoincr"`
	Index      int
	AccountID  string `xorm:"'account_id' notnull unique"`
	SecertAddr string `xorm:"notnull unique"`
	Balance    float64
	Success    string `xorm:"varchar(1) notnull"`
}

// TAccount 账户定义
type TAccount struct {
	ID             uint64    `xorm:"'id' pk autoincr"`
	AccountID      string    `xorm:"'account_id' notnull unique"`
	SecertAddr     string    `xorm:"notnull unique"`
	CreateTime     time.Time `xorm:"DateTime"`
	CreateTimeUnix int64
	LastUpdateTime time.Time `xorm:"DateTime"`
	UpdateTimeUnix int64
	Active         string
}

// TStaticsActiveAccount 统计激活账户
type TStaticsActiveAccount struct {
	ID            uint64 `xorm:"'id' pk autoincr"`
	Index         int
	CreateTime    time.Time `xorm:"DateTime"`
	Counter       int
	StartTimeTick int64
	EndTimeTick   int64
	SignerB64     string `xorm:"'signer' varchar(65535) notnull"`
	Action        string `xorm:"varchar(100) notnull"`
	Success       string `xorm:"varchar(1) notnull"`
}
