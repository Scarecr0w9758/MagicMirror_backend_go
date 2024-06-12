package models

type Tokens struct {
	Id             uint
	UserId         uint
	AccessToken    string
	RefreshToken   string
	ExpirationDate int64
}

type User struct {
	Id            uint   `json:"id"`
	Login         string `json:"login"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	RememberToken string `json:"rememberToken"`
}

type Task struct {
	Id          uint   `json:"id"`
	UserId      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed" sql:"DEFAULT:false;type:boolean;index" gorm:"not null"`
	StartDate   int64  `json:"start_date"`
	EndDate     int64  `json:"end_date"`
}