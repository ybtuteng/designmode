package decorate

type Topping struct {
    drink Drink
    price int
    description string
}

func (t Topping) Cost() int {
    return t.drink.Cost() + t.price
}

func (t Topping) Describe() string {
    return t.description
}
