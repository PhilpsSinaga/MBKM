package controller

import (
	"GO_MINIPROJECT/model"
	"GO_MINIPROJECT/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IGymController interface {
	GetAllUser(c echo.Context) error
	GetUserById(c echo.Context) error
	UpdateUserById(c echo.Context) error
	InsertUser(c echo.Context) error
	DeleteUserById(c echo.Context) error

	//Admin
	GetAllAdmin(c echo.Context) error
	GetAdminById(c echo.Context) error
	UpDateAdminById(c echo.Context) error
	InsertAdmin(c echo.Context) error
	DeleteAdminById(c echo.Context) error

	//Trainer
	GetAllTrainer(c echo.Context) error
	GetTrainerById(c echo.Context) error
	UpdateTrainerById(c echo.Context) error
	InsertTrainer(c echo.Context) error
	DeleteTrainerById(c echo.Context) error

	//Class
	GetAllClass(c echo.Context) error
	GetClassById(c echo.Context) error
	UpdateClassById(c echo.Context) error
	InsertClass(c echo.Context) error
	DeleteClassById(c echo.Context) error

	//Membership
	GetAllMembership(c echo.Context) error
	GetMembershipById(c echo.Context) error
	UpDateMembershipById(c echo.Context) error
	InsertMembership(c echo.Context) error
	DeleteMembershipById(c echo.Context) error

	//Payment_Method
	GetAllPaymentMethod(c echo.Context) error
	GetPaymentMethodById(c echo.Context) error
	UpdatePaymentMethodbyId(c echo.Context) error
	InsertPaymentMethod(c echo.Context) error
	DeletePaymentMethodById(c echo.Context) error
}

type GymController struct {
	iGymRepo repository.IGymRepo
}

func NewGymController(iGymRepo repository.IGymRepo) IGymController {
	return &GymController{iGymRepo: iGymRepo}
}

var appJSON = "application/json"

func (ctrl GymController) GetAllUser(c echo.Context) error {
	user, err := ctrl.iGymRepo.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":     "Berhasil Melihat Semua Data User",
		"Daftar User": user,
	})
}

func (ctrl GymController) GetUserById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	user, err := ctrl.iGymRepo.GetUserById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, user)
}

func (ctrl GymController) UpdateUserById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	NamaUser := model.User{}
	c.Bind(&NamaUser)
	err := ctrl.iGymRepo.UpDateUserById(Id, NamaUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaUser)
}

func (ctrl GymController) InsertUser(c echo.Context) error {
	NamaUser := model.User{}
	c.Bind(&NamaUser)
	err := ctrl.iGymRepo.InsertUser(NamaUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes insert User",
	})

}

func (ctrl GymController) DeleteUserById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	err := ctrl.iGymRepo.DeleteUserById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes Delete data User dengan id '" + id + "'",
	})
}

func (ctrl GymController) GetAllAdmin(c echo.Context) error {
	admin, err := ctrl.iGymRepo.GetAllAdmin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":      "Berhasil Melihat Semua Data Admin",
		"Daftar Admin": admin,
	})
}

func (ctrl GymController) GetAdminById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	adminId, err := ctrl.iGymRepo.GetAdminById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, adminId)
}

func (ctrl GymController) UpDateAdminById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	NamaAdmin := model.Admins{}
	c.Bind(&NamaAdmin)
	err := ctrl.iGymRepo.UpDateAdminById(Id, NamaAdmin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaAdmin)
}

func (ctrl GymController) InsertAdmin(c echo.Context) error {
	NamaAdmin := model.Admins{}
	c.Bind(&NamaAdmin)
	err := ctrl.iGymRepo.InsertAdmin(NamaAdmin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes insert Admin",
	})
}

func (ctrl GymController) DeleteAdminById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	err := ctrl.iGymRepo.DeleteAdminById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes Delete data Admin dengan id '" + id + "'",
	})
}

func (ctrl GymController) GetAllTrainer(c echo.Context) error {
	trainer, err := ctrl.iGymRepo.GetAllTrainer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":        "Berhasil Melihat Semua Data Trainer",
		"Daftar Trainer": trainer,
	})
}

func (ctrl GymController) GetTrainerById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	trainerId, err := ctrl.iGymRepo.GetTrainerById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, trainerId)
}

func (ctrl GymController) UpdateTrainerById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	NamaTrainer := model.Trainer{}
	c.Bind(&NamaTrainer)
	err := ctrl.iGymRepo.UpdateTrainerById(Id, NamaTrainer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaTrainer)
}

func (ctrl GymController) InsertTrainer(c echo.Context) error {
	NamaTrainer := model.Trainer{}
	c.Bind(&NamaTrainer)
	err := ctrl.iGymRepo.InsertTrainer(NamaTrainer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes insert Trainer",
	})
}

func (ctrl GymController) DeleteTrainerById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	err := ctrl.iGymRepo.DeleteTrainerById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes Delete data Trainer dengan id '" + id + "'",
	})

}

func (ctrl GymController) GetAllClass(c echo.Context) error {
	class, err := ctrl.iGymRepo.GetAllClass()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":        "Berhasil Melihat Semua Data Class",
		"Daftar Trainer": class,
	})
}

func (ctrl GymController) GetClassById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	classId, err := ctrl.iGymRepo.GetClassById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, classId)
}

func (ctrl GymController) UpdateClassById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	NamaClass := model.Class{}
	c.Bind(&NamaClass)
	err := ctrl.iGymRepo.UpdateClassById(Id, NamaClass)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaClass)
}

func (ctrl GymController) InsertClass(c echo.Context) error {
	NamaClass := model.Class{}
	c.Bind(&NamaClass)
	err := ctrl.iGymRepo.InsertClass(NamaClass)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes insert Class",
	})
}

func (ctrl GymController) DeleteClassById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	err := ctrl.iGymRepo.DeleteClassById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes Delete data Class dengan id '" + id + "'",
	})

}

func (ctrl GymController) GetAllMembership(c echo.Context) error {
	member, err := ctrl.iGymRepo.GetAllMembership()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":        "Berhasil Melihat Semua Data Membership",
		"Daftar Trainer": member,
	})
}

func (ctrl GymController) GetMembershipById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	memberId, err := ctrl.iGymRepo.GetMembershipById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, memberId)
}

func (ctrl GymController) UpDateMembershipById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	NamaMember := model.Membership{}
	c.Bind(&NamaMember)
	err := ctrl.iGymRepo.UpDateMembershipById(Id, NamaMember)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaMember)
}

func (ctrl GymController) InsertMembership(c echo.Context) error {
	NamaMembership := model.Membership{}
	c.Bind(&NamaMembership)
	err := ctrl.iGymRepo.InsertMembership(NamaMembership)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes insert Membership",
	})
}

func (ctrl GymController) DeleteMembershipById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	err := ctrl.iGymRepo.DeleteMembershipById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes Delete data Membership dengan id '" + id + "'",
	})

}

func (ctrl GymController) GetAllPaymentMethod(c echo.Context) error {
	payment, err := ctrl.iGymRepo.GetAllPaymentMethod()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":        "Berhasil Melihat Semua Data Payment Method",
		"Daftar Trainer": payment,
	})
}

func (ctrl GymController) GetPaymentMethodById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	paymentId, err := ctrl.iGymRepo.GetPaymentMethodById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, paymentId)
}

func (ctrl GymController) UpdatePaymentMethodbyId(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	NamaPayment := model.Payment_method{}
	c.Bind(&NamaPayment)
	err := ctrl.iGymRepo.UpdatePaymentMethodbyId(Id, NamaPayment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaPayment)
}

func (ctrl GymController) InsertPaymentMethod(c echo.Context) error {
	NamaPayment := model.Payment_method{}
	c.Bind(&NamaPayment)
	err := ctrl.iGymRepo.InsertPaymentMethod(NamaPayment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes insert Payment Method",
	})
}

func (ctrl GymController) DeletePaymentMethodById(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	err := ctrl.iGymRepo.DeletePaymentMethodById(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succes Delete data Payment Method dengan id '" + id + "'",
	})

}
