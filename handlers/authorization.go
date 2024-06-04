package handlers

type signInInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GenerateToken(username,password)  {
	
}