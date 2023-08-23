package patterns

import "fmt"

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct {
	a int
	b int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Считаем площадь квадрата...")
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("Считаем площадь круга...")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Считаем площадь прямоугольника...")
}

func main() {
	square := &Square{2}
	circle := &Circle{5}
	rectangle := &Rectangle{2, 4}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)
}

// Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
// а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.
