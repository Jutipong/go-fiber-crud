package repository

import (
	"fiber-crud/features/address/model"

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

func (m *repositoryMock) Create(auth *model.Address) error {
	args := m.Called()
	return args.Error(1)
}

func (m *repositoryMock) Update(auth *model.Address) error {
	args := m.Called()
	return args.Error(1)
}

func (m *repositoryMock) Delete(auth *model.Address) error {
	args := m.Called()
	return args.Error(1)
}
