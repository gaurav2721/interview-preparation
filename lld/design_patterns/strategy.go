package design_patterns

// This situation arises when different child are using the same functionality but they are not inheriting it from the interface
// for eg if there is a Vehicle class but there is subclasses Motorcycle , Car , Bicycle for eg they are having the code for eg Display which is there only in the Motorcycle and Car
// Now this causes duplicacy of the code for each and every child class
// so you create a strategy interface and put iStrategy object in parent and then set it in the child's constructor

// type IDisplayStrategy interface {
// 	Display()
// }

// type DieselDisplayStrategy struct {
// }

// func (dieselDisplay DieselDisplayStrategy) Display() {
// 	fmt.Println("display diesel vehicle")
// }

// type ElectricDisplayStrategy struct {
// }

// func (electricDisplay ElectricDisplayStrategy) Display() {
// 	fmt.Println("display electric vehicle")
// }

// type Vehicle struct {
// 	IDisplay IDisplayStrategy
// }

// func (v Vehicle) Display() {
// 	v.IDisplay.Display()
// }

// type Bike struct {
// 	Vehicle
// 	No_of_wheels int
// }

// type Car struct {
// 	Vehicle
// 	No_of_wheels int
// }

// code in main.go
// electricDisplayStrategy := design_patterns.ElectricDisplayStrategy{}
// b := design_patterns.Bike{Vehicle: design_patterns.Vehicle{IDisplay: electricDisplayStrategy}, No_of_wheels: 2}
// b.Display()
