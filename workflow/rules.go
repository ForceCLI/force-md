package workflow

func (w *Workflow) GetRules() []Rule {
	return w.Rules
}

func (w *Workflow) GetAlerts() []Alert {
	return w.Alerts
}

func (w *Workflow) GetFieldUpdates() []FieldUpdate {
	return w.FieldUpdates
}
