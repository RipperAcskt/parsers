package queue

type Queue []string

func New(list []string) *Queue {
	var q Queue
	for _, l := range list {
		q = append(q, l)
	}
	return &q
}

func (q *Queue) Get() string {
	var elem string
	if len(*q) > 0 {
		elem = (*q)[0]
		*q = (*q)[1:]
	}
	*q = append(*q, elem)
	return elem
}
