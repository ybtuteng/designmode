package decorate

type Coffee struct {
    price int
    description string
}

func (c Coffee) Cost() int {
    return c.price
}

func (c Coffee) Describe() string {
    return c.description
}
