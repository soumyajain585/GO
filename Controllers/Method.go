package Controllers

type Rectangle struct {
	Width  float64
	Height float64
}

// Method definition
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
