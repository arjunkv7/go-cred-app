package models

type User struct {
	Firstname string `json:"firstName" binding:"required"`
	Lastname  string `json:"lastName" binding:"required"`
	Age       string `json:"age" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Userid    string `json:"userid"`
}

type Products struct {
	Productid          string `json:"productid" binding:"required" bson:"productid"`
	Productname        string `json:"productname" binding:"required" bson:"productname"`
	Productprice       string `json:"productprice" binding:"required" bson:"productprice"`
	Productdescription string `json:"productdescription" bson:"productdescription"`
	Productimage       string `json:"productimage" bson:"productimage"`
	Productcategory    string `json:"productcategory" bson:"productcategory"`
	Productstock       string `json:"productstock" bson:"productstock"`
	Productrating      string `json:"productrating" bson:"productrating"`
}
