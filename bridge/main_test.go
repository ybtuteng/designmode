package bridge

import (
    "log"
    "testing"
)

func TestBridge(t *testing.T) {
    var l Coffee
    l = NewLatte(NewLarge())
    log.Println(l.order(2))
    var m Coffee
    m = NewMocha(NewMiddle())
    log.Println(m.order(4))
}
