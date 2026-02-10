package database

import (
	"user-service/models"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

type IUserDB interface {
	Create(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetAllByLimit(limit, offset int) ([]models.User, error)
	GetUserByEmailWithPassword(email string) (*models.User, error)
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db}
}

func (u *UserDB) Create(user *models.User) (*models.User, error) {
	// err := u.DB.AutoMigrate(models.User{}) // It creates a table in the database
	// if err != nil {
	// 	return nil, err
	// }
	//luser, err := u.GetUserByEmail(user.Email)

	tx := u.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (c *UserDB) GetAll() ([]models.User, error) {
	users := make([]models.User, 0)
	tx := c.DB.Find(&users)
	return users, tx.Error
}

func (c *UserDB) GetAllByLimit(limit, offset int) ([]models.User, error) {
	users := make([]models.User, 0)
	tx := c.DB.Limit(limit).Offset(offset).Find(&users)
	return users, tx.Error
}

func (u *UserDB) GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	tx := u.DB.Where("email = ?", email).First(user)
	return user, tx.Error
}

func (u *UserDB) GetUserByEmailWithPassword(email string) (struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}, error) {
	var user struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//tx := u.DB.Where("email = ?", email).Table("users").First(&user)
	tx := u.DB.Where("email = ?", email).Model(models.User{}).First(&user)
	return user, tx.Error
}
