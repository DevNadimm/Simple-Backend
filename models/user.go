package models

type User struct {
	ID          int    `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	IsShopOwner bool   `db:"is_shop_owner" json:"is_shop_owner"`
}

type PublicUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (u User) ToPublicUser() PublicUser {
	return PublicUser{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		IsShopOwner: u.IsShopOwner,
	}
}
