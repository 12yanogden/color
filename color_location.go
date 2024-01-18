package colors

type Location struct {
	Start int
	End   int
}

func (l1 *Location) Equal(l2 Location) bool {
	if (*l1).Start != l2.Start ||
		(*l1).End != l2.End {
		return false
	}

	return true
}

func (l1 *Location) Length(l Location) int {
	return l.End - l.Start
}
