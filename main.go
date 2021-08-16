package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

//get all users from
func GetUsersCotroller(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id done
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{}

	for _, val := range users {
		if val.Id == id {
			user.Id = val.Id
			user.Name = val.Name
			user.Email = val.Email
			user.Password = val.Password
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"users":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for key, val := range users {
		if val.Id == id {
			users = append(users[:key], users[key+1:]...)
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
		"users":    users,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	nameIn := c.FormValue("name")
	emailIn := c.FormValue("email")
	passwordIn := c.FormValue("password")
	user := User{}

	for key, val := range users {
		if val.Id == id {
			user.Id = id
			if val.Name != nameIn && nameIn != "" {
				user.Name = nameIn
				users[key].Name = nameIn
			} else {
				user.Name = val.Name
			}
			if val.Email != emailIn && emailIn != "" {
				user.Email = emailIn
				users[key].Email = emailIn
			} else {
				user.Email = val.Email
			}
			if val.Password != passwordIn && passwordIn != "" {
				user.Password = passwordIn
				users[key].Password = passwordIn
			} else {
				user.Password = val.Password
			}
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"users":    user,
	})
}

// create user by id
func CreateUserController(c echo.Context) error {
	//binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1

	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()

	//routing
	e.GET("/users", GetUsersCotroller)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}
