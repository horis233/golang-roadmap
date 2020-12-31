package Iterator

type Iterator interface{
	Index() int
	Next()
	HasNext() bool
	Value() interface{}
}

type ArrayIterator struct{
	array []interface{}
	index *int
}

func (a *ArrayIterator) Index() *int{
	return a.index
}

func (a *ArrayIterator) Next(){
	if a.HasNext(){
		*a.index ++
	}
}

func (a *ArrayIterator) HasNext() bool{
	return *a.index + 1 <= len(a.array)
}

func (a *ArrayIterator) Value() interface{}{
	return a.array[*a.index]
}
