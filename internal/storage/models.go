package storage

type User struct {
	Id           int
	Username     string
	PasswordHash string
}

type Url struct {
	Id     int
	Url    string
	Alias  string
	UserId int
}
