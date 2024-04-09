package lld

// -------------------->    example of abstraction

// type Shape struct {
// 	area      int
// 	perimeter int
// }

// func PrintShape(s *Shape) {
// 	fmt.Println(Shape.area, Shape.perimeter)
// }

// --------------------->    example of encapsulation , encapsulation = bundling of a member variables and member functions into a single entity ,
// --------------------->    encapsulation = data hiding (using access modifiers , in case of golang it is capital and small letters) + abstraction (you should be able to interact with an object only by using functions)

// type Shape struct {
// 	area      int
// 	perimeter int
// }

// func (s *Shape) PrintArea() {
// 	fmt.Println(s.area)
// }

// func (s *Shape) PrintPerimeter() {
// 	fmt.Println(s.perimeter)
// }

// func GetShape(area int, perimeter int) *Shape {
// 	return &Shape{area: area, perimeter: perimeter}
// }

// ----------------------> example of inheritance is achieved through structure embedding

// type Shape struct {
// 	area      int
// 	perimeter int
// }

// type Square struct {
// 	Shape
// 	sideLength int
// }

// func GetSquare(sideLength int) *Square {
// 	return &Square{sideLength: sideLength, Shape: Shape{area: 1, perimeter: 1}}
// }

// func (s *Square) PrintArea() {
// 	fmt.Println(s.area)
// }

// ------------------------> example of polymorphism is achieved through interfaces
// ------------------------> to remember when you are importing an interface on the subclass the function should be like func (s Square) instead of func (s *Square)

// type Shape interface {
// 	PrintArea()
// 	PrintPerimeter()
// }

// type Square struct {
// 	sideLength int
// }

// func GetSquare(sideLength int) Shape {
// 	return Square{sideLength: sideLength}
// }

// func (s Square) PrintArea() {
// 	fmt.Println(s.sideLength * s.sideLength)
// }

// func (s Square) PrintPerimeter() {
// 	fmt.Println(4 * s.sideLength)
// }
