package bridge

type Latte struct {
    size Size
}

func NewLatte(size Size) Latte {
    return Latte{size: size}
}

func (m Latte) order(count int) int {
    return (m.size.cost() + 20) * count
}