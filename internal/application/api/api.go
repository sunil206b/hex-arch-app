package api

import "github.com/sunil206b/hex-arc-app/internal/ports"

type Application struct {
	db    ports.DBPort
	arith Arithmetic
}

func NewApplication(db ports.DBPort, arith Arithmetic) *Application {
	return &Application{
		db:    db,
		arith: arith,
	}
}

func (ad *Application) GetAddition(a, b int32) (int32, error) {
	ans, err := ad.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = ad.db.AddToHistory(ans, "addition")
	if err != nil {
		return 0, err
	}
	return ans, nil
}
func (ad *Application) GetSubtraction(a, b int32) (int32, error) {
	ans, err := ad.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = ad.db.AddToHistory(ans, "subtraction")
	if err != nil {
		return 0, err
	}
	return ans, nil
}
func (ad *Application) GetMultiplication(a, b int32) (int32, error) {
	ans, err := ad.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = ad.db.AddToHistory(ans, "multiplication")
	if err != nil {
		return 0, err
	}

	return ans, nil
}
func (ad *Application) GetDivision(a, b int32) (int32, error) {
	ans, err := ad.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = ad.db.AddToHistory(ans, "division")
	if err != nil {
		return 0, err
	}
	return ans, nil
}
