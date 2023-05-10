package entity

import (
	"time"

	"github.com/viettranx/service-context/core"
)

// Gender user's gender
type Gender string

// Gender enums
const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderUnknown Gender = "unknown"
)

// SystemRole user's role
type SystemRole string

// SystemRole enums
const (
	RoleSuperAdmin SystemRole = "sadmin"
	RoleAdmin      SystemRole = "admin"
	RoleUser       SystemRole = "user"
)

// Status user's status
type Status string

// Status enums
const (
	StatusActive        Status = "active"
	StatusPendingVerify Status = "waiting_verify"
	StatusBanned        Status = "banned"
)

// User defines data model for user.
type User struct {
	ID         uint       `json:"-" gorm:"primary_key;column:id;auto_increment:true"`
	FakeID     *core.UID  `json:"id" gorm:"-"`
	FirstName  string     `json:"first_name" gorm:"column:first_name"`
	LastName   string     `json:"last_name" gorm:"column:last_name"`
	Email      string     `json:"email" gorm:"column:email"`
	Phone      string     `json:"phone" gorm:"column:phone"`
	Avatar     string     `json:"avatar" gorm:"avatar"`
	Gender     Gender     `json:"gender" gorm:"column:gender"`
	SystemRole SystemRole `json:"system_role" gorm:"column:system_role"`
	Status     Status     `json:"status" gorm:"column:status"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at;autocreatetime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at;autoupdatetime"`
}

// NewUser returns a user.
func NewUser(firstName, lastName, email string) User {
	return User{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Phone:      "",
		Avatar:     "",
		Gender:     GenderUnknown,
		SystemRole: RoleUser,
		Status:     StatusActive,
	}
}

func (User) TableName() string { return "users" }

const (
	MaskTypeUser = 1
)

func (u *User) Mask() {
	uid := core.NewUID(uint32(u.ID), MaskTypeUser, 1)
	u.FakeID = &uid
}
