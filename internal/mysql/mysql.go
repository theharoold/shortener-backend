package mysql

import (
	"fmt"

	"github.com/mailru/dbr"
	"github.com/theharoold/shortener-backend/config"
)

type MySQL struct {
	DbConn *dbr.Connection
	DbConf config.DbConfig
}

func (m MySQL) generateDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", m.DbConf.Username, m.DbConf.Password, m.DbConf.Host, m.DbConf.Port, m.DbConf.DbName)
}

func (m *MySQL) OpenConnection(driver string) error {
	var err error
	m.DbConn, err = dbr.Open(driver, m.generateDSN(), nil)
	return err
}

func (m *MySQL) CloseConnection() error {
	return m.DbConn.Close()
}

