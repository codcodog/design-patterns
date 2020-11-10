package chain

import "fmt"

type IBaseCase interface {
	setNextCase(IBaseCase)
	handleRequest(int)
}

type Case1 struct {
	next IBaseCase
}

func (c *Case1) setNextCase(cc IBaseCase) {
	c.next = cc
}

func (c *Case1) handleRequest(price int) {
	if price < 100 {
		fmt.Printf("case1 handle %d RMB\n", price)
		return
	}

	if c.next != nil {
		c.next.handleRequest(price)
	}
}

type Case2 struct {
	next IBaseCase
}

func (c *Case2) setNextCase(cc IBaseCase) {
	c.next = cc
}

func (c *Case2) handleRequest(price int) {
	if price >= 100 && price < 1000 {
		fmt.Printf("case2 handle %d RMB\n", price)
		return
	}

	if c.next != nil {
		c.next.handleRequest(price)
	}
}

type Case3 struct {
	next IBaseCase
}

func (c *Case3) setNextCase(cc IBaseCase) {
	c.next = cc
}

func (c *Case3) handleRequest(price int) {
	if price >= 1000 {
		fmt.Printf("case3 handle %d RMB\n", price)
		return
	}

	if c.next != nil {
		c.next.handleRequest(price)
	}
}
