package users

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"column:name"`
	LastName string `json:"last_name" gorm:"column:last_name"`

	Email string `json:"email" gorm:"column:email;unique_index"`

	Username string `json:"username" gorm:"column:username;unique_index"`
	Password string `json:"password" gorm:"column:password;not null"`
}

// FullName return user full name (name + last_name)
func (u *User) FullName() string {
	return u.Name + " " + u.LastName
}
