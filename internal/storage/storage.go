package storage

type Storage interface {
	SaveUrl(url string, alias string) error
	GetUrl(alias string) (string, error)
	DeleteUrl(alias string) error
	GetAllUrls(user_id int) (*[]Url, error)
	SaveUser(username, password_hash string) error
	GetUserByUsername(username string) (*User, error)
	GetAllUsers() (*[]User, error)
	DeleteUser(username string) error
}
