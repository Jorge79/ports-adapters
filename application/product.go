package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	ChangePrice(price float64) error
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
}

func (p *Product) Enable() error {
}

func (p *Product) ChangePrice(price float64) error {
}

func (p *Product) Disable() error {
}

func (p *Product) GetID() string {
}

func (p *Product) GetName() string {
}

func (p *Product) GetStatus() string {
}

func (p *Product) GetPrice() float64 {
}
