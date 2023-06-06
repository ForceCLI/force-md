package reportFolder

type FolderShareFilter func(FolderShare) bool

func (o *ReportFolder) GetShares(filters ...FolderShareFilter) []FolderShare {
	var shares []FolderShare
SHARES:
	for _, a := range o.FolderShares {
		for _, filter := range filters {
			if !filter(a) {
				continue SHARES
			}
		}
		shares = append(shares, a)
	}
	return shares
}

func (o *ReportFolder) DeleteShares(filters ...FolderShareFilter) {
	var shares []FolderShare
SHARES:
	for _, a := range o.FolderShares {
		for _, filter := range filters {
			if !filter(a) {
				shares = append(shares, a)
				continue SHARES
			}
		}
	}
	o.FolderShares = shares
}
