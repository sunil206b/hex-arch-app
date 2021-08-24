package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ad *Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}
func (ad *Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}
func (ad *Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}
func (ad *Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
