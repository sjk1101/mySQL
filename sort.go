package mySQL

type Sorter interface {
	GetSorts() Sorts
}

type Sort struct {
	Column string
	IsDesc bool
}

type Sorts []*Sort

var _ Sorter = (Sorts)(nil)

func (s Sorts) GetSorts() Sorts {
	return s
}

func (s *Sorts) Sort(column string, isDesc bool) {
	*s = append(*s, &Sort{
		Column: column,
		IsDesc: isDesc,
	})
}
