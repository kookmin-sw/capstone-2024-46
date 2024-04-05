package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var NameConverter = schema.NamingStrategy{
	SingularTable: false, // use plural table name
	NoLowerCase:   false, // use lower case
}

type DBConfig interface {
	NewDialector() gorm.Dialector
	getDSN() string
}

// MySQLConfig is config struct for general mysql usage
// See: https://github.com/go-sql-driver/mysql?tab=readme-ov-file#dsn-data-source-name
type MySQLConfig struct {
	// DSN format like "[username[:password]@][protocol[(address)]]/dbname"
	DSN string

	// Charset sets the charset used for client-server interaction
	Charset string

	// Location sets the location for time.Time values (when using parseTime=true)
	// "Local" sets the system's location.
	Location string

	// ParseTime changes the output type of DATE and DATETIME values to time.Time instead of []byte / string
	ParseTime bool

	// Timeout for establishing connections, aka dial timeout.
	// The value must be a decimal number with a unit suffix ("ms", "s", "m", "h"), such as "30s", "0.5m" or "1m30s".
	Timeout string
}

var _ DBConfig = (*MySQLConfig)(nil)

func (c MySQLConfig) getDSN() string {
	return fmt.Sprintf("%s?charset=%s&parseTime=%t&loc=%s&timeout=%s", c.DSN, c.Charset, c.ParseTime, c.Location, c.Timeout)
}

func (c MySQLConfig) NewDialector() gorm.Dialector {
	return mysql.Open(c.getDSN())
}

func NewMySQLConfig(mysqlDSN string) *MySQLConfig {
	return &MySQLConfig{
		DSN:       mysqlDSN,
		Charset:   "utf8mb4",
		Location:  "UTC",
		ParseTime: true,
		Timeout:   "10s",
	}
}

func OpenConnection(dbConfig DBConfig) (*gorm.DB, error) {
	return gorm.Open(
		dbConfig.NewDialector(),
		&gorm.Config{
			NamingStrategy:  NameConverter,
			CreateBatchSize: 1000,
		},
	)
}
