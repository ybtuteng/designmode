package decorate

type Latte struct {
    Coffee
}

func NewLatte() Latte {
    l := Latte{}
    l.price = 10
    l.description = "latte"
    return l
}
