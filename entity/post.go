package entity

type Person struct {
	Firstname 	string `json:"firstname" binding:"required"`
	Lastname	string `json:"lastname" binding:"required"`
	Age			int8	`json:"age" binding:"gte=1,lte=30"`
	Email		string	`json:"email" binding:"required,email"`
}

type Post struct {
	Title       string `json:"title" binding:"min=2,max=200" validate:"is-okay"`
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	Author 		Person `json:"author" binding:"required"`
}
