package models

type User struct {
	ID             uint `gorm:"primaryKey"`
	Username       string
	Email          string
	HashedPassword string
	IsSuperuser    bool
}
type Roles struct {
	ID          uint
	ServiceName string
	Permission  string
}
type UserRoles struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint
	User   User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Role   Roles `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE;"`
}
type Info struct {
	ID  uint
	Smt string
}
