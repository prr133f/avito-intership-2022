package v1

func (r *Route) Balance() {
	(*r.Group).Get("/balance/:id", r.View.GetBalance)
	(*r.Group).Put("/balance", r.View.AddBalance)
	(*r.Group).Put("/reserve", r.View.ReserveAmount)
}
