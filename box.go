package golang_united_school_homework

import (
	"errors"
)

var (
	valueNotFound   = errors.New("Value not found")
	valueOutOfRange = errors.New("Value out of range")
)

// коробка содержит список фигур и может выполнять над ними операции
type box struct {
	shapes         []Shape
	shapesCapacity int // 	Максимальное количество фигур, которые могут быть внутри коробки.
}

// NewBox создает новый экземпляр блока
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape добавляет фигуру в коробку
// возвращает ошибку, если она выходит за пределы диапазона shapeCapacity.
func (b *box) AddShape(shape Shape) error {

	if len(b.shapes)+1 > b.shapesCapacity {
		return valueOutOfRange
	}
	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex позволяет получить фигуру по индексу
// если фигура по индексу не существует или индекс вышел за пределы диапазона, то возвращает ошибку
func (b *box) GetByIndex(i int) (Shape, error) {

	if i+1 > b.shapesCapacity {
		return nil, valueOutOfRange
	}
	for idx, value := range b.shapes {
		if idx == i {
			if value == nil {
				return nil, valueNotFound
			}
			return value, nil
		}

	}
	return nil, nil
}

// ExtractByIndex позволяет получить фигуру по индексу и удаляет эту фигуру из списка.
// если фигура по индексу не существует или индекс вышел за пределы диапазона, то возвращает ошибку
func (b *box) ExtractByIndex(i int) (Shape, error) {

	if i+1 > b.shapesCapacity {
		return nil, valueOutOfRange
	}
	for idx, value := range b.shapes {
		if idx == i {
			if value == nil {
				return nil, valueNotFound

			}
			b.shapes = append(b.shapes[0:i], b.shapes[i+1:]...)
			return value, nil

		}

	}
	return nil, nil
}

// ReplaceByIndex позволяет заменить фигуру по индексу и возвращает удаленную фигуру.
// если фигура по индексу не существует или индекс вышел за пределы диапазона, то возвращает ошибку
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {

	if i+1 > b.shapesCapacity {
		return nil, valueOutOfRange
	}
	for idx, value := range b.shapes {
		if i == idx {
			if value == nil {
				return nil, valueNotFound
			}

			b.shapes[i] = shape
			return value, nil

		}

	}
	return nil, nil
}

// SumPerimeter предоставляет суммарный периметр всех фигур в списке.
func (b *box) SumPerimeter() float64 {

	var sum float64
	for _, value := range b.shapes {
		sum += value.CalcPerimeter()
	}
	return sum

}

// SumArea обеспечивает суммарную площадь всех фигур в списке.
func (b *box) SumArea() float64 {

	var sum float64
	for _, value := range b.shapes {
		sum += value.CalcArea()
	}
	return sum
}

// RemoveAllCircles удаляет все круги в списке
// если кружков в списке нет, то возвращает ошибку
func (b *box) RemoveAllCircles() error {

	firstLen := len(b.shapes)
	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); ok {
			b.shapes = append(b.shapes[0:i], b.shapes[i+1:]...)
			i--
		}

	}
	if firstLen == len(b.shapes) {
		return errors.New("Circles not found")
	}

	return nil
}
