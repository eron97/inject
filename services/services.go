package service

import (
	database "github.com/eron97/inject.git/repository"
)

type DomainService interface {
	CreateUserServices()
}

type useDomainService struct {
	userRepository database.Database
}

func NewDomainService(
	db database.Database,
) DomainService {
	return &useDomainService{db}
}

func (uds *useDomainService) CreateUserServices() {
}
