package parse

// Derived from https://gist.github.com/bemasher/1777766

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) Peek() (value interface{}) {
	if s.size > 0 {
		return s.top.value
	}
	return nil
}

