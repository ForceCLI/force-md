package globalvalueset

func (o *GlobalValueSet) UpdateGlobalValueSet(updates GlobalValueSet) error {
	if updates.Sorted != nil {
		o.Sorted = updates.Sorted
	}
	return nil
}
