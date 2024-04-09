package main

import (
	"fmt"

	"github.com/gaurav2721/interview-preparation/lld/design_patterns"
)

func main() {
	fmt.Println("gaurav is starting his preps")
	//var iDisplayStrategy design_patterns.IDisplayStrategy
	electricDisplayStrategy := design_patterns.ElectricDisplayStrategy{}
	b := design_patterns.Bike{Vehicle: design_patterns.Vehicle{IDisplay: electricDisplayStrategy}, No_of_wheels: 2}
	b.Display()

}
