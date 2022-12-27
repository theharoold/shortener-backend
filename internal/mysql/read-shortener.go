package mysql

func (m *MySQL) ReadShortener(key string) (string, error) {
	sess := m.DbConn.NewSession(nil)
	var url string
	builder := sess.Select("plaintext_url").From(m.DbConf.Tables.Shorteners).Where("shortened_path = ?", key)
	err := builder.LoadValue(&url)
	return url, err
}
