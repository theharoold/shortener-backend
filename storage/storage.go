package storage

type Storage interface {
	ShortenURL(url string) (string, error)
	RetrieveURL(key string) (string, error)
}
