package repository

import (
	"fiber-crud/app/address/model"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func NewRepositoryMock() *repositoryMock {
	return &repositoryMock{}
}

func (m *repositoryMock) Inquiry(auth *model.Address) ([]model.Address, error) {
	args := m.Called()
	return args.Get(0).([]model.Address), args.Error(1)
}
