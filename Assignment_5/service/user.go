package service

type User struct {
	Nama string
}

type UserService struct {
	users []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
}

func NewUserService() UserIface {
	return &UserService{}
}

func (u *UserService) Register(user *User) string {
	u.users = append(u.users, user)
	return user.Nama + " berhasil didaftarkan"
}

func (u *UserService) GetUser() []*User {
	return u.users
}
