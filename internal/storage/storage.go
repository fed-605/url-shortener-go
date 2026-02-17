package storage

type Storage interface {
	SaveUrl(url string, alias string) error
	GetUrl(alias string) (string, error)
	DeleteUrl(alias string) error
	GetAllRecords() ([]Url, error)
}
