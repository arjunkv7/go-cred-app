package store

import "go-cred-app/models"

var Users []models.User

var Products []models.Products = []models.Products{
	{
		Productid: "1",
		Productname: "Product 1",
		Productprice: "100",
		Productdescription: "Product 1 description",
	},
	{
		Productid: "2",
		Productname: "Product 2",
		Productprice: "200",
		Productdescription: "Product 2 description",
	},
	{
		Productid: "3",
		Productname: "Product 3",
		Productprice: "300",
		Productdescription: "Product 3 description",
	},
	{
		Productid: "4",
		Productname: "Product 4",
		Productprice: "400",
		Productdescription: "Product 4 description",
	},
}	
