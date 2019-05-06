package collection

/*Element knows only about the addresses of its neighbors*/
type Element struct {
	name string
	prev, next *Element
}

/*Next returns the next element address*/
func (e *Element) Next() *Element {
	if(e.next != nil) {return e.next}
	return nil //probably last element
}

/*Prev returns the previous element address*/
func (e *Element) Prev() *Element {
	if (e.prev != nil) {return e.prev}
	return nil //probably first element
}

/*Value returns the current element name*/
func (e Element) Value() string {
	return e.name
}