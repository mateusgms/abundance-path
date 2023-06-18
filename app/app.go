package app

import (
    "AbundancePath/core/service"
    "fmt"
)

type App struct {
    userService service.UserService
}

func New(userService service.UserService) *App {
    return &App{userService: userService}
}

func (a *App) Run() error {
    fmt.Println("App is running...")
    return nil
}
