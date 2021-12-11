package user

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

const (
	userTable = "users"
)

func (User) TableName() string {
	return userTable
}
