package entity

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json:"description" binding:"max=20"`
	Url         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}

type Person struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=30"`
	Email     string `json:"email" validate:"email,required"`
}
