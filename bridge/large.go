package bridge

type Large struct {
    price int
}

func NewLarge() Large {
    return Large{10}
}

func (l Large) cost() int {
    return l.price
}
