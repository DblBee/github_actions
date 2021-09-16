package models

import (
	"time"

	"github.com/dblbee/github_actions/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Account struct {
	Base
	Name    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Enabled bool   `gorm:"type:boolean" json:"enabled"`
}

func (*Account) TableName() string {
	return "account.accounts"
}

func NewAccountRepo() AccountRepo {
	dbConn := database.NewPostgresDatabase()
	dbConn.Setup()

	repo := AccountRepo{
		DB: dbConn.GetDB(),
	}

	repo.DB.AutoMigrate(
		&Account{},
	)

	return repo
}

type AccountRepo struct {
	DB *gorm.DB
}

func (r *AccountRepo) Get(id uuid.UUID) (*Account, error) {
	account := new(Account)

	err := r.DB.Where("id = ?", id).Find(account).Error

	return account, err
}
