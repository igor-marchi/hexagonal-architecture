package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type IProductInterface interface {
	IsValid() (bool, error)
	EnableD() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type IProductServiceInterface interface {
	Get(id string) (IProductInterface, error)
	Create(name string, price float64) (IProductInterface, error)
	Enable(product IProductInterface) (IProductInterface, error)
	Disable(product IProductInterface) (IProductInterface, error)
}

type IProductReader interface {
	Get(id string) (IProductInterface, error)
}

type IProductWriter interface {
	Save(product IProductInterface) (IProductInterface, error)
}

type IProductPersistenceInterface interface {
	IProductReader
	IProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float, optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}

	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the status must be greater or equal zero")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price <= 0 {
		return errors.New("the price must be greater than zero to enabled the product")
	}

	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() error {
	if p.Price != 0 {
		return errors.New("the price must be zero in order to have the product disabled")
	}

	p.Status = DISABLED
	return nil
}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
