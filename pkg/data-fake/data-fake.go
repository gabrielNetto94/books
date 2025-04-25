package datafake

import "github.com/go-faker/faker/v4"

type DataFake interface {
	FirstName() string
	FullName() string
	Age() int
	Email() string
}

type dataFake struct{}

func New() DataFake {
	return &dataFake{}
}

func (d *dataFake) FirstName() string {
	return faker.Name()
}

func (d *dataFake) FullName() string {
	return faker.Name()
}
func (d *dataFake) Age() int {
	res, _ := faker.RandomInt(18, 99)
	return res[0]
}

func (d *dataFake) Email() string {
	return faker.Email()
}
