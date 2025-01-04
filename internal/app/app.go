package app

import "keeper/internal/service/user"

type App struct {
	UserService *user.Service
}

func NewApp(us *user.Service) *App {
	return &App{
		UserService: us,
	}
}
