package data

import "ChatServe/utils"

// User 用户的结构体
type User struct {
	Userid           int
	Username         string
	Password         string
	Email            string
	Image            string
	Integral         int
	Registrationdate utils.JSONTime
}
