package cache

type Cache interface {
	SaveUrlMapping(url, alias string) error
	RetrieveUrl(alias string) (string, error)
	DeleteUrl(alias string) error
}
