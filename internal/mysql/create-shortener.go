package mysql

func (m *MySQL) CreateShortener(url string, key string) error {
	sess := m.DbConn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	builder := tx.InsertInto(m.DbConf.Tables.Shorteners).Columns(ShortenedURLMySQLColumns...).Values(url, key)
	_, err = builder.Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
