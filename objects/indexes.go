package objects

import "github.com/ForceCLI/force-md/objects/index"

func (o *CustomObject) GetIndexes(filters ...index.IndexFilter) []index.BigObjectIndex {
	var indexes []index.BigObjectIndex
INDEXES:
	for _, i := range o.Indexes {
		for _, filter := range filters {
			if !filter(i) {
				continue INDEXES
			}
		}
		indexes = append(indexes, i)
	}
	return indexes
}
