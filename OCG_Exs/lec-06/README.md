# Lecture 6- Final assignment
In this lecture, I learned the following:
* how to write middleware for logging and authentication
* how to design ER Diagram.
> For more clealy, I build my own project to manage Products, Orders.
Please let me know if you are not clear on somewhere.
## Database design 
![Image of ERD]
(/home/duy/Downloads/erd6.png)
![GitHub Logo](/home/duy/Downloads/erd6.png)
For more clearly redirect to : ![dbdiagram.io](url)

## Features
* Basic Login
* CRUD APIs
* Logging

## Usage
> To start project, let login first

* Login: 
  * Open Postman and choose Authorization
  * Select type BasicAuth
  * Enter: Duy for User, 123 for Password 

* C>R>U>D 
Customer
```t
    /api/customers > > > Methods(POST)
	/api/customersMethods(GET)
	/api/customers/{id} > > > Methods(PUT)
	/api/customers/{id}> > > Methods(DELETE)
```
Product
```t
	/api/productsMethods(POST)
	/api/productsMethods(GET)
	/api/products/{id}> > > Methods(PUT)
	/api/products/{id}> > > Methods(DELETE)
```
Order
```t
	/api/ordersMethods(POST)
	/api/ordersMethods(GET)
	/api/orders/{id}> > > Methods(PUT)
	/api/orders/{id}> > > Methods(DELETE)
```
Orderdetail
```t
	/api/orderdetailsMethods(POST)
	/api/orderdetailsMethods(GET)
	/api/orderdetails/{id}> > > Methods(PUT)
	"/api/orderdetails/{id}"> > > > Methods("DELETE")
```


## Some errors
### Have trouble creating tables: 
Description:
>In the customer table I have a customer with id =1, then I deleted this customer and i can not  to return this customer, but its still there. Because when gorm.model can record my  activities as column create_at or delete_at. However, in the orders table I try to add this deleted  id =1 and it says success.