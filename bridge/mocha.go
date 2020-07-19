package bridge

type Mocha struct {
    size Size
}

func NewMocha(size Size) Mocha {
    return Mocha{size: size}
}

func (m Mocha) order(count int) int {
    return (m.size.cost() + 15) * count
}
