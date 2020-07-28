package task2

type Book []*ContactItem

func (b Book) Len() int {
	return len(b)
}

func (b Book) Less(i, j int) bool {
	return b[i].Phone < b[j].Phone
}

func (b Book) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
