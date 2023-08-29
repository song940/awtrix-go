package awtrix

type Fill struct {
	Effect
	Color Color `json:"color"`
}

func NewFill(r, g, b int) Fill {
	return Fill{
		Effect: Effect{Type: "fill"},
		Color:  Color{r, g, b},
	}
}
