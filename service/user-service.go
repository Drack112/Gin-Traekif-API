package service

import (
    "log"

    "github.com/Drack112/Crud-Golang-API/dto"
    "github.com/Drack112/Crud-Golang-API/entity"
    "github.com/Drack112/Crud-Golang-API/repository"
    "github.com/mashingan/smapping"
)

type UserService interface {
    Update(user dto.UserUpdateDTO) entity.User
    Profile(userID string) entity.User
}

type userService struct {
    userRepository repository.UserRepository
}

func NewUserService(userRep repository.UserRepository) UserService {
    return &userService{
        userRepository: userRep,
    }
}

func (serviceUser *userService) Update(user dto.UserUpdateDTO) entity.User {
    userToUpdate := entity.User{}
    err := smapping.FillStruct(&userToUpdate, smapping.MapFields(user))
    if err != nil {
        log.Fatalf("Failed map %v", err)
    }

    updateUser := serviceUser.userRepository.UpdateUser(userToUpdate)
    return updateUser
}

func (serviceUser *userService) Profile(userID string) entity.User {
    return serviceUser.userRepository.ProfileUser(userID)
}
