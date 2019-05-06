package collection

import "fmt"

type Collection struct {
	len int
	firstEl, lastEl *Element
}

/*Add new node Element type*/
func (c *Collection) Add(name string) { //returns index/-1 ?
	el := Element {name: name, prev: c.lastEl}

	if (c.len == 0) {
		c.firstEl = &el
	} else if (c.lastEl != nil) { //check not required
		c.lastEl.next = &el
	}

	c.lastEl =  &el
	c.len++
} 

/*Get element/node by index*/
func (c *Collection) Get(index int) *Element {
	var el *Element
	if (index < c.len) {
		el = c.firstEl
		for i:=0; i < index && el != nil; i++ {
			el = el.Next()
		}
	}
	return el //return <nil> if not found element by index
}

/*Remove node by index
* "Need translate
* Полагаем что элементы не могут прерываться в списке, следовательно последный элемент имеет next = nil, первый элемент имеет prev = nil"
*/
func (c *Collection) Remove(index int) { //retuns 0/-1 or true/false ?
	el := c.Get(index)
	if (el != nil) {
		if (el.Prev() != nil) {
			el.Prev().next = el.Next()
		} else {
			c.firstEl = el.Next()
		}

		if (el.Next() != nil) {
			el.Next().prev = el.Prev()
		} else {
			c.lastEl = el.Prev()
		}

		c.len--
	}
}

/*Print - writes to standard output*/
func (c *Collection) Print() {
	el := c.firstEl
	count := 0
	for ;el != nil; {
		fmt.Printf("Index %d: %+v \n", count, el)
		el = el.Next()
		count++
	}
}

/*returns First element/node*/
func (c *Collection) First() *Element {
	return c.firstEl
}

/*returns Last element/node*/
func (c *Collection) Last() *Element {
	return c.lastEl
}

func (c *Collection) Length() int {
	return c.len
}