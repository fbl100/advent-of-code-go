package grid

type AbyssError struct {
}

func (m AbyssError) Error() string {
	return "Fell off the world"
}
