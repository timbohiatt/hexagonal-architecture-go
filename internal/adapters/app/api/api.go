package api

import (
	"hexagonal-architecture-go/internal/ports"
)

type Adapter struct {
	db     ports.DbPort
	artith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, artith ports.ArithmeticPort) *Adapter {
	return &Adapter{
		artith: artith,
	}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.artith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	// Run Database Operation
	err = apia.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, err
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.artith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	// Run Database Operation
	err = apia.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}

	return answer, err
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.artith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	// Run Database Operation
	err = apia.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, err
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.artith.Division(a, b)
	if err != nil {
		return 0, err
	}

	// Run Database Operation
	err = apia.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, err
}
