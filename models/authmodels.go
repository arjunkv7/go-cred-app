package models


type User struct{
	Firstname string `json:"firstName" binding:"required"`
	Lastname string `json:"lastName" binding:"required"`
	Age string `json:"age" binding:"required"`
	Password string `json:"password" binding:"required"`
	Userid string `json:"userid"`
}

type Products struct {
	Productid string `json:"productid" binding:"required"`
	Productname string `json:"productname" binding:"required"`
	Productprice string `json:"productprice" binding:"required"`
	Productdescription string `json:"productdescription" `
	Productimage string `json:"productimage"`
	Productcategory string `json:"productcategory"`
	Productstock string `json:"productstock"`
	Productrating string `json:"productrating"`
}