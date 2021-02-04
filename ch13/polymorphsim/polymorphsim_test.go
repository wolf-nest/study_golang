package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WirteHelloworld() Code
}

type GoProgrammer struct {
}

func (sel *GoProgrammer) WirteHelloworld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaPorgrammer struct {
}

func (sel *JavaPorgrammer) WirteHelloworld() Code {
	return "System.out.println(\"Hello World!\")"
}

func wirteFirstProgrammmer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WirteHelloworld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaPorgrammer)
	wirteFirstProgrammmer(goProg)
	wirteFirstProgrammmer(javaProg)
}
