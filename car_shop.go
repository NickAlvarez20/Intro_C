package main

import "fmt"

// learn oop by creating a car shop

type Car struct {
	make  string
	model string
	year  int
	price float64
	color string
}

func (c Car) getMake() string {
	return c.make
}

func (c Car) getModel() string {
	return c.model
}

func (c Car) getYear() int {
	return c.year
}

func (c Car) getPrice() float64 {
	return c.price
}

func (c Car) getColor() string {
	return c.color
}

func (c *Car) setMake(make string) {
	c.make = make
}

func (c *Car) setModel(model string) {
	c.model = model
}

func (c *Car) setYear(year int) {
	c.year = year
}

func (c *Car) setPrice(price float64) {
	c.price = price
}

func (c *Car) setColor(color string) {
	c.color = color
}

type CarSlice []Car

func (cs *CarSlice) addCar(car Car) {
	*cs = append(*cs, car)
}

func (cs *CarSlice) removeCar(index int) bool {
	if index < 0 || index >= len(*cs) {
		return false
	}
	*cs = append((*cs)[:index], (*cs)[index+1:]...)
	return true
}

func (cs CarSlice) getCar(index int) (Car, bool) {
	if index < 0 || index >= len(cs) {
		return Car{}, false
	}
	return cs[index], true
}

// CarShop wraps inventory and exposes a small public API.
type CarShop struct {
	inventory CarSlice // private field — use AddCar / RemoveCar instead of touching directly
}

func NewCarShop() *CarShop {
	return &CarShop{inventory: CarSlice{}}
}

// --- public (exported) methods ---

func (s *CarShop) AddCar(car Car) bool {
	if !s.validateCar(car) {
		return false
	}
	s.inventory.addCar(car)
	return true
}

func (s *CarShop) RemoveCar(index int) bool {
	return s.inventory.removeCar(index)
}

func (s *CarShop) GetCar(index int) (Car, bool) {
	return s.inventory.getCar(index)
}

func (s *CarShop) InventorySummary() string {
	return fmt.Sprintf("%d cars, total value $%.2f", len(s.inventory), s.totalInventoryValue())
}

// --- private (unexported) helpers ---

func (s *CarShop) validateCar(car Car) bool {
	if car.getMake() == "" || car.getModel() == "" {
		return false
	}
	if car.getYear() < 1886 || car.getPrice() < 0 {
		return false
	}
	return true
}

func (s *CarShop) totalInventoryValue() float64 {
	var total float64
	for _, car := range s.inventory {
		total += car.getPrice()
	}
	return total
}

func main() {
	car := Car{
		make:  "Toyota",
		model: "Corolla",
		year:  2020,
		price: 10000.00,
		color: "Red",
	}

	fmt.Printf("Make: %s\n", car.getMake())
	fmt.Printf("Model: %s\n", car.getModel())
	fmt.Printf("Year: %d\n", car.getYear())
	fmt.Printf("Price: %f\n", car.getPrice())
	fmt.Printf("Color: %s\n", car.getColor())

	car.setMake("Ford")
	car.setModel("Mustang")
	car.setYear(2021)
	car.setPrice(15000.00)
	car.setColor("Blue")

	fmt.Printf("Make: %s\n", car.getMake())
	fmt.Printf("Model: %s\n", car.getModel())
	fmt.Printf("Year: %d\n", car.getYear())
	fmt.Printf("Price: %f\n", car.getPrice())
	fmt.Printf("Color: %s\n", car.getColor())
	fmt.Printf("Car: %+v\n", car)

	shop := NewCarShop()
	if shop.AddCar(car) {
		fmt.Println(shop.InventorySummary())
	}

	if got, ok := shop.GetCar(0); ok {
		fmt.Printf("Shop car at index 0: %+v\n", got)
	}

	if shop.RemoveCar(0) {
		fmt.Println("After remove:", shop.InventorySummary())
	}

	invalid := Car{make: "", model: "Ghost", year: 2020, price: -1}
	if !shop.AddCar(invalid) {
		fmt.Println("Rejected invalid car (private validateCar blocked it)")
	}
}
