package proxy

import (
	"fmt"
	"testing"
)

type Base struct {
	name string
	age  int
}

func (this *Base) Println() {
	fmt.Println(this)
}

type Child struct {
	*Base
	fire bool
}

func TestIt(t *testing.T) {

	base := Base{name: "yuankui", age: 12}

	base.Println()

	child := Child{Base: &base, fire: true}

	child.Println()
}
