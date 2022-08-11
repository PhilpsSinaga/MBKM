package repository

import (
	"GO_MINIPROJECT/model"
	"fmt"

	"gorm.io/gorm"
)

type IGymRepo interface {
	//Login
	GetUserUsernamePassword(username, password string) (model.User, error)

	//User
	GetAllUser() ([]model.User, error)
	GetUserById(id int) (model.User, error)
	UpDateUserById(id int, user model.User) error
	InsertUser(user model.User) error
	DeleteUserById(id int) error

	//Admin
	GetAllAdmin() ([]model.Admins, error)
	GetAdminById(id int) (model.Admins, error)
	UpDateAdminById(id int, admin model.Admins) error
	InsertAdmin(admin model.Admins) error
	DeleteAdminById(id int) error

	//Trainer
	GetAllTrainer() ([]model.Trainer, error)
	GetTrainerById(id int) (model.Trainer, error)
	UpdateTrainerById(id int, trainer model.Trainer) error
	InsertTrainer(trainer model.Trainer) error
	DeleteTrainerById(id int) error

	//Class
	GetAllClass() ([]model.Class, error)
	GetClassById(id int) (model.Class, error)
	UpdateClassById(id int, class model.Class) error
	InsertClass(clas model.Class) error
	DeleteClassById(id int) error

	//Membership
	GetAllMembership() ([]model.Membership, error)
	GetMembershipById(id int) (model.Membership, error)
	UpDateMembershipById(id int, membership model.Membership) error
	InsertMembership(membership model.Membership) error
	DeleteMembershipById(id int) error

	//PaymentMethod
	GetAllPaymentMethod() ([]model.Payment_method, error)
	GetPaymentMethodById(id int) (model.Payment_method, error)
	UpdatePaymentMethodbyId(id int, payment model.Payment_method) error
	InsertPaymentMethod(payment model.Payment_method) error
	DeletePaymentMethodById(id int) error
}

type GymRepo struct {
	db *gorm.DB
}

func NewGymRepo(db *gorm.DB) IGymRepo {
	return &GymRepo{db}
}

func (r GymRepo) GetUserUsernamePassword(username, password string) (model.User, error) {
	var user model.User
	err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		fmt.Println("Error while Get Email and Password", err)
	}
	return user, err
}

func (r GymRepo) GetAllUser() ([]model.User, error) {
	user := []model.User{}
	err := r.db.Find(&user).Error
	if err != nil {
		fmt.Println("Error while Get All User", err)
	}
	return user, err
}

func (r GymRepo) GetUserById(id int) (model.User, error) {
	var userId model.User
	err := r.db.First(&userId, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println("Error while Get User")
	}
	return userId, err
}

func (r GymRepo) UpDateUserById(id int, user model.User) error {
	err := r.db.Model(&user).Where("id = ?", id).Updates(user).Error
	if err != nil {
		fmt.Println("Error While Update User")
	}
	return err
}
func (r GymRepo) InsertUser(user model.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		fmt.Println("Error while Insert User", err)
	}
	return err
}

func (r GymRepo) DeleteUserById(id int) error {
	err := r.db.Delete(&model.User{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println("Error While Delete User", err)
	}
	return err
}

func (r GymRepo) GetAllAdmin() ([]model.Admins, error) {
	admin := []model.Admins{}
	err := r.db.Find(&admin).Error
	if err != nil {
		fmt.Println("Error while Get All Admin", err)
	}
	return admin, err

}

func (r GymRepo) GetAdminById(id int) (model.Admins, error) {
	var adminId model.Admins
	err := r.db.First(&adminId, "id = ?", id).Error
	if err != nil {
		fmt.Println("Error while Get  Admin", id)
	}
	return adminId, err

}

func (r GymRepo) UpDateAdminById(id int, admin model.Admins) error {
	err := r.db.Model(&admin).Where("id = ?", id).Updates(admin).Error
	if err != nil {
		fmt.Println("Error While Update Admin")
	}
	return err
}

func (r GymRepo) InsertAdmin(admin model.Admins) error {
	err := r.db.Create(&admin).Error
	if err != nil {
		fmt.Println("Error while InsertAdmin", err)
	}
	return err
}

func (r GymRepo) DeleteAdminById(id int) error {
	err := r.db.Delete(&model.Admins{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println("Error While Delete Admin", err)
	}
	return err
}

func (r GymRepo) GetAllTrainer() ([]model.Trainer, error) {
	trainer := []model.Trainer{}
	err := r.db.Find(&trainer).Error
	if err != nil {
		fmt.Println("Error while Get All Trainer", err)
	}
	return trainer, err
}

func (r GymRepo) GetTrainerById(id int) (model.Trainer, error) {
	var trainerId model.Trainer
	err := r.db.First(&trainerId, "id = ?", id).Error
	if err != nil {
		fmt.Println("Error while Get trainer ", id)
	}
	return trainerId, err
}

func (r GymRepo) UpdateTrainerById(id int, trainer model.Trainer) error {
	err := r.db.Model(&trainer).Where("id = ?", id).Updates(trainer).Error
	if err != nil {
		fmt.Println("Error While Update Trainer")
	}
	return err
}

func (r GymRepo) InsertTrainer(trainer model.Trainer) error {
	err := r.db.Create(&trainer).Error
	if err != nil {
		fmt.Println("Error while InsertTrainer", err)
	}
	return err
}

func (r GymRepo) DeleteTrainerById(id int) error {
	err := r.db.Delete(&model.Trainer{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println("Error While Delete Trainer", err)
	}
	return err
}

func (r GymRepo) GetAllClass() ([]model.Class, error) {
	clas := []model.Class{}
	err := r.db.Find(&clas).Error
	if err != nil {
		fmt.Println("Error while Get Class", err)
	}
	return clas, err
}

func (r GymRepo) GetClassById(id int) (model.Class, error) {
	var classId model.Class
	err := r.db.First(&classId, "id = ?", id).Error
	if err != nil {
		fmt.Println("Error while Get Class ", id)
	}
	return classId, err
}

func (r GymRepo) UpdateClassById(id int, class model.Class) error {
	err := r.db.Model(&class).Where("id = ?", id).Updates(class).Error
	if err != nil {
		fmt.Println("Error While Update Class")
	}
	return err
}

func (r GymRepo) InsertClass(class model.Class) error {
	err := r.db.Create(&class).Error
	if err != nil {
		fmt.Println("Error while Insert Class", err)
	}
	return err
}

func (r GymRepo) DeleteClassById(id int) error {
	err := r.db.Delete(&model.Class{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println("Error While Delete Class", err)
	}
	return err
}

func (r GymRepo) GetAllMembership() ([]model.Membership, error) {
	member := []model.Membership{}
	err := r.db.Find(&member).Error
	if err != nil {
		fmt.Println("Error while Get Membership", err)
	}
	return member, err
}

func (r GymRepo) GetMembershipById(id int) (model.Membership, error) {
	var memberId model.Membership
	err := r.db.First(&memberId, "id = ?", id).Error
	if err != nil {
		fmt.Println("Error while Get membership ", id)
	}
	return memberId, err
}

func (r GymRepo) UpDateMembershipById(id int, member model.Membership) error {
	err := r.db.Model(&member).Where("id = ?", id).Updates(member).Error
	if err != nil {
		fmt.Println("Error While Update Class")
	}
	return err
}

func (r GymRepo) InsertMembership(member model.Membership) error {
	err := r.db.Create(&member).Error
	if err != nil {
		fmt.Println("Error while Insert Membership", err)
	}
	return err
}

func (r GymRepo) DeleteMembershipById(id int) error {
	err := r.db.Delete(&model.Membership{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println("Error While Delete Membership", err)
	}
	return err
}

func (r GymRepo) GetAllPaymentMethod() ([]model.Payment_method, error) {
	payment := []model.Payment_method{}
	err := r.db.Find(&payment).Error
	if err != nil {
		fmt.Println("Error while Get All Payment", err)
	}
	return payment, err
}

func (r GymRepo) GetPaymentMethodById(id int) (model.Payment_method, error) {
	var paymentId model.Payment_method
	err := r.db.First(&paymentId, "id = ?", id).Error
	if err != nil {
		fmt.Println("Error while Get Payment ", id)
	}
	return paymentId, err
}

func (r GymRepo) UpdatePaymentMethodbyId(id int, payment model.Payment_method) error {
	err := r.db.Model(&payment).Where("id = ?", id).Updates(payment).Error
	if err != nil {
		fmt.Println("Error While Update payment")
	}
	return err
}

func (r GymRepo) InsertPaymentMethod(payment model.Payment_method) error {
	err := r.db.Create(&payment).Error
	if err != nil {
		fmt.Println("Error while Insert payment", err)
	}
	return err
}

func (r GymRepo) DeletePaymentMethodById(id int) error {
	err := r.db.Delete(&model.Payment_method{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println("Error While Delete payment", err)
	}
	return err
}
