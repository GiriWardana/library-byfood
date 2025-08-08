package test

import (
	"github.com/stretchr/testify/mock"

	"backend/models"
)

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) GetAllBooks() ([]models.Book, error) {
	args := m.Called()
	return args.Get(0).([]models.Book), args.Error(1)
}

func (m *MockBookService) GetBookByID(id int) (models.Book, error) {
	args := m.Called(id)
	return args.Get(0).(models.Book), args.Error(1)
}

func (m *MockBookService) CreateBook(book models.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookService) UpdateBook(id int, book models.Book) error {
	args := m.Called(id, book)
	return args.Error(0)
}

func (m *MockBookService) DeleteBook(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
