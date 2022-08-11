package main

import (
	"GO_MINIPROJECT/configuration"
	"GO_MINIPROJECT/controller"
	"GO_MINIPROJECT/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	db, err := configuration.InitDB()
	if err != nil {
		panic(err)
	}
	err = configuration.InitialMigration(db)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	iGymRepo := repository.NewGymRepo(db)
	ilogin := controller.NewLoginController(iGymRepo)
	iGymController := controller.NewGymController(iGymRepo)
	// Routes

	//Login
	e.POST("/user/login", ilogin.Login)

	//User
	e.GET("/user", iGymController.GetAllUser)
	e.GET("/user/:id", iGymController.GetUserById)
	e.PUT("/user/:id", iGymController.UpdateUserById)
	e.POST("/user", iGymController.InsertUser)
	e.DELETE("/user/:id", iGymController.DeleteUserById)

	//Admin
	e.GET("/admin", iGymController.GetAllAdmin)
	e.GET("/admin/:id", iGymController.GetAdminById)
	e.PUT("/admin/:id", iGymController.UpDateAdminById)
	e.POST("/admin", iGymController.InsertAdmin)
	e.DELETE("/admin/:id", iGymController.DeleteAdminById)

	//Trainer
	e.GET("/trainer", iGymController.GetAllTrainer)
	e.GET("/trainer/:id", iGymController.GetTrainerById)
	e.PUT("/trainer/:id", iGymController.UpdateTrainerById)
	e.POST("/trainer", iGymController.InsertTrainer)
	e.DELETE("/trainer/:id", iGymController.DeleteTrainerById)

	//Class
	e.GET("/class", iGymController.GetAllClass)
	e.GET("/class/:id", iGymController.GetClassById)
	e.PUT("/class/:id", iGymController.UpdateClassById)
	e.POST("/class", iGymController.InsertClass)
	e.DELETE("/class/:id", iGymController.DeleteClassById)

	//Membership
	e.GET("/member", iGymController.GetAllMembership)
	e.GET("/member/:id", iGymController.GetMembershipById)
	e.PUT("/member/:id", iGymController.UpDateMembershipById)
	e.POST("/member", iGymController.InsertMembership)
	e.DELETE("/member/:id", iGymController.DeleteMembershipById)

	//Payment Method
	e.GET("/payment", iGymController.GetAllPaymentMethod)
	e.GET("/payment/:id", iGymController.GetPaymentMethodById)
	e.PUT("/payment/:id", iGymController.UpdatePaymentMethodbyId)
	e.POST("/payment", iGymController.InsertPaymentMethod)
	e.DELETE("/payment/:id", iGymController.DeletePaymentMethodById)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
