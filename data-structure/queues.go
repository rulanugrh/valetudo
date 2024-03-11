package data

type Queues struct {
	Elems []string
}

func (q *Queues) InsertQueues(elem string) {
	if elem == "" {
		println("data null")
	}

	q.Elems = append(q.Elems, elem)
}

func (q *Queues) DeleteQueues() string {
	if len(q.Elems) == 0 {
		println("Null Data")
	}

	element := q.Elems[0]
	if len(q.Elems) == 1 {
		q.Elems = nil
		return element
	}

	q.Elems = q.Elems[1:]
	return element
}
