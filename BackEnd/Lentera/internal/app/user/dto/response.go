package dto

import 
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/entity"

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type AuthResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

func FromUser(u entity.User) UserResponse {
	return UserResponse{ID: u.ID, Name: u.Name, Email: u.Email}
}

