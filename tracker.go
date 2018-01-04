package anomalytracker

// Publishable validates that all required fields are set for inserting.
func (a *Anomaly) Publishable() bool {
	return a != nil && a.Id != nil && a.System != nil &&
		a.Type != nil && a.Name != nil
}
