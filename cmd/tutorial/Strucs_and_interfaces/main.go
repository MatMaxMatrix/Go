package main

import "fmt"

type gasEngine struct {
	mpg uint16
	gals uint16
}

type electricEngine struct {
	mpkwh uint16
	kwh uint16
}

type engine interface{
	milesLeft() uint16
}

func (e gasEngine) milesLeft() uint16 {
	return e.gals * e.mpg
}

func (e electricEngine) milesLeft() uint16 {
	return e.kwh * e.mpkwh
}

// Generic function that works with any engine type
func canMakeIt(e engine, miles uint16) bool {
    remaining := e.milesLeft()
    if miles <= remaining {
        fmt.Printf("We will make it! (%d miles remaining)\n", remaining)
        return true
    }
    fmt.Printf("We will not make it! (only %d miles remaining)\n", remaining)
    return false
}

func main() {
	var myEngine gasEngine = gasEngine{25, 20}
	var myElectricEngine electricEngine = electricEngine{240, 55}

	canMakeIt(myEngine, 50)
	canMakeIt(myElectricEngine, 50)
	canMakeIt(myEngine, 255)
	canMakeIt(myElectricEngine, 255)
}