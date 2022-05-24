package main

import (
	"crud/config"
	"crud/controller"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	ErrorCode int         `json:"erorr_code" form:"erorr_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}

func main() {
	config.ConDB()
	route := echo.New()
	route.GET("user/:id", func(c echo.Context) error {
		response := new(Response)
		id, err := controller.GetUserById(c.Param("id"))
		if err != nil {
			response.ErrorCode = 500
			response.Message = "Cannot read data"
		} else {
			response.ErrorCode = 0
			response.Message = "Succes"
			response.Data = id
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/all", func(c echo.Context) error {
		response := new(Response)
		users, err := controller.GetUserByAll(c.QueryParam("alldata"))
		if err != nil {
			response.ErrorCode = 500
			response.Message = "Cannot read data"
		} else {
			response.ErrorCode = 0
			response.Message = "Succes"
			response.Data = users
		}
		return c.JSON(http.StatusOK, response)
	})

	route.POST("user/create", func(c echo.Context) error {
		user := new(controller.Users)
		c.Bind(user)
		response := new(Response)
		if user.CreateData() != nil {
			response.ErrorCode = 400
			response.Message = "Cannot create data"
		} else {
			response.ErrorCode = 0
			response.Message = "Succes"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update/:id", func(c echo.Context) error {
		user := new(controller.Users)
		c.Bind(user)
		response := new(Response)
		if user.UpdateData(c.Param("id")) != nil {
			response.ErrorCode = 400
			response.Message = "Cannot update data"
		} else {
			response.ErrorCode = 0
			response.Message = "Succes update data"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("user/delete/:id", func(c echo.Context) error {
		user, _ := controller.GetUserById(c.Param("id"))
		response := new(Response)
		if user.DeleteData() != nil {
			response.ErrorCode = 500
			response.Message = "Cennot delete data"
		} else {
			response.ErrorCode = 0
			response.Message = "Succes delete data"
		}
		return c.JSON(http.StatusOK, response)
	})

	fmt.Println("server start on port 9000")
	fmt.Println("Succes")
	route.Start(":9000")
}
