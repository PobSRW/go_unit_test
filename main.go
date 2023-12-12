package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

// การใช้ mock ของ testify
func main() {

	c := CustomerRepositoryMock{}

	// mock data
	c.On("GetCustomer", 1).Return("Pob", 18, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))

	// operate
	name, age, err := c.GetCustomer(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name, age)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

// ทุกครั้งที่ทำเป็น mock ต้องเรียก recv func เป็น pointer
func (c *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := c.Called(id)

	return args.String(0), args.Int(1), args.Error(2)
}
