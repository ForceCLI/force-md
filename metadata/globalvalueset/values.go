package globalvalueset

func (o *GlobalValueSet) GetValues(filters ...ValueFilter) []CustomValue {
	var values []CustomValue
VALUES:
	for _, v := range o.CustomValue {
		for _, filter := range filters {
			if !filter(v) {
				continue VALUES
			}
		}
		values = append(values, v)
	}
	return values
}
