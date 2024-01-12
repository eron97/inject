package models

type CreateUser struct {
	First_Name string `json:"first_name" binding:"required,min=4,max=100" example:""`
	Last_Name  string `json:"last_name" binding:"required,min=4,max=100" example:""`
	Email      string `json:"email" binding:"required,email" example:"test@test.com"`
	Password   string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"password#@#@!2121"`
}

type GetUser struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BirthDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
}

type GetUserID struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

/*

// Get Users
type GetUserByEmail struct {
	First_Name string `json:"first_name" binding:"required,min=4,max=100" example:""`
	Last_Name  string `json:"last_name" binding:"required,min=4,max=100" example:""`
}

*/

/*

Camada de controllers faz o binding dos dados (pode chamar um package validate para isso) dentro da própria função ou simplesmente passar como middleware em routes.Routes

Camada de service tem os métodos a serem implementados:


*/
