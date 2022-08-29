package service

type User struct {
	Nama string
}

type UserService struct {
}

type UserIface interface {
	Register(u *User) string
}

func NewUserService() UserIface {
	return &UserService{}
}

func (u *UserService) Register(user *User) string {
	return user.Nama + "berhasil didaftarkan"
}
