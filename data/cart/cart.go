package cart

type Cart struct {
	CartType  string `json:"carttype"`
	ValueName string `json:"valuename"`
	Value     int8   `json:"value"`
}

func NewCart(cartType string, valueName string, value int8) *Cart {
	return &Cart{CartType: cartType, ValueName: valueName, Value: value}
}

/*
func (cart Cart) GetType() (cartType string) {
	cartType = cart.cartType
	return
}

func (cart Cart) GetValueName() (cartType string) {
	cartType = cart.valueName
	return
}

func (cart Cart) GetValue() (value int8) {
	value = cart.value
	return
}
*/
