package main

import "github.com/common-nighthawk/go-figure"

func main() {
	myFigure := figure.NewFigure("Hallo", "", true)
	myFigure.Dance(10000, 200)
	myFigure.Print()
}
