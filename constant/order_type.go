package constant

type OrderType int

const (
	ItemBase OrderType = iota + 1
	ItemVariant
	ItemPromo
)

var OrderTypes []string = []string{"BASE", "VARIANT", "PROMO"}

func (t OrderType) String() string {
	return OrderTypes[t-1]
}

func (t OrderType) EnumIndex() int {
	return int(t)
}
