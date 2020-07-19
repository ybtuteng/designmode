package bridge

type Middle struct {
    price int
}

func NewMiddle() Middle {
    return Middle{5}
}

func (l Middle) cost() int {
    return l.price
}
