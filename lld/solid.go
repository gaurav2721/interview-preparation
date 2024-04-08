package lld

// Single Responsibility Principle :- a class should have only one reason to change(Marker , calculateInvoice and SaveInvoice) -> Marker and InvoiceDao object
// Open for extension but closed for modification , (InvoiceDao) interface is inherited by (InvoiceFile) , (InvoiceDb)

// Liskov substitution principle :- if class b is subtype of class A , then we should be able to replace object of class A with class B without breaking the program
// Bike interface with two objects bike and cycle , but turnOnEngine will not be present in cycle , however increaseSpeed will be present in both
// To fix for Liskov Substitution Principle : we split classes for eg Vehicle is parent and EngineVehicle and Bicycle are subclass and EngineVehicle has subclasses Car and MotorBike

// Interface Segregation Principle :- A client should not have to implement unnecessary functions they do not need to implement for example
// Restaurant Employee (take order , make dish ) -> split into two interface waiter interface and chef interface

// Dependency inversion principle :- a class should depend on interfaces , rather that concrete classes , for eg if there is a macbook class , it should use iMouse and iKeyboard,
// for eg these can be mouse can be a wired and bluetooth one

type Marker struct {
	id     string
	colour string
	price  int
}

func (m *Marker) GetInvoice(quantity int) int {
	return m.price * quantity
}

func (m *Marker) SaveInvoice(invoice_amount int) error {

}
