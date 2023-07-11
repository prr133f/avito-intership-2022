package v1

func (r *Route) Report() {
	(*r.Group).Get("/report/:date", r.View.MonthlyReport)
}
