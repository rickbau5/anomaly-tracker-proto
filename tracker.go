package anomalytracker

// Publishable validates that all required fields are set for inserting.
func (a *Anomaly) Publishable() bool {
	return a != nil && a.Id != "" && a.System != "" &&
		a.Type != "" && a.Name != ""
}
