package arithmetic

type Arith struct {
}

func New() *Arith {
	return &Arith{}
}

func (ad *Arith) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}
func (ad *Arith) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}
func (ad *Arith) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}
func (ad *Arith) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
