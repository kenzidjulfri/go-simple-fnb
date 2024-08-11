package constant

type Category int

const (
	CategoryFood Category = iota + 1
	CategoryBeverage
)

var Categories []string = []string{"FOOD", "BEVERAGE"}

func (t Category) String() string {
	return Categories[t-1]
}

func (t Category) EnumIndex() int {
	return int(t)
}
