package table

type cell struct {
	value string
}

func NewCell(value string) *cell {
	return &cell{value: value}
}

func (c cell) String() string {
	return c.value
}
