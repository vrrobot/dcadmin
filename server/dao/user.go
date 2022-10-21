package dao

import (
	"dcadmin/models"
	"log"

	// "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *models.User)
	ListUser() []models.User
	GetUser(username string) models.User
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "host=localhost user=dcadmin password=DCadmin@123 dbname=dcsys port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库初始化失败！", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.User{})
}

func (mgr *manager) AddUser(user *models.User) {
	mgr.db.Create(user)
}

func (mgr *manager) ListUser() []models.User {
	var users []models.User
	mgr.db.Find(&users)
	return users
}

func (mgr *manager) GetUser(username string) models.User {
	var user models.User
	mgr.db.Where("username = ?", username).First(&user)
	return user
}
