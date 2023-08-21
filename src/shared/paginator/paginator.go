package paginator

import "pulzo/app"

type Paginator struct {
	page            int
	limit           int
	numItemsPerPage int
	iniItem         int
	endItem         int
	totalItem       int
}

func NewPaginator(page int) *Paginator {
	paginator := Paginator{}
	paginator.SetPage(page)
	paginator.setNumItemsPerPage()
	return &paginator
}

func (c *Paginator) SetPage(page int) {
	c.page = page
}

func (c *Paginator) SetTotal(total int) {
	c.totalItem = total
}

func (c *Paginator) setNumItemsPerPage() {
	c.numItemsPerPage = 10
	if app.NumItemsPerPage() > 0 {
		c.numItemsPerPage = app.NumItemsPerPage()
	}
}

func (c *Paginator) Offset() int {
	if c.page == 1 {
		return 0
	}
	c.limit = (c.page * c.numItemsPerPage) - c.numItemsPerPage

	return c.limit
}

func (c *Paginator) generateLimitInitAndEnd() {
	c.iniItem = ((c.page * c.numItemsPerPage) - c.numItemsPerPage) + 1
	c.endItem = c.page * c.numItemsPerPage

	if c.iniItem > c.totalItem {
		c.totalItem = 0
		c.iniItem = 0
		c.endItem = 0
	}

	if c.endItem > c.totalItem {
		c.endItem = c.totalItem
	}
}

func (c *Paginator) NumItemsPerPage() int {
	return c.numItemsPerPage
}

func (c *Paginator) Initial() int {
	return c.iniItem
}

func (c *Paginator) End() int {
	return c.endItem
}

func (c *Paginator) Total() int {
	return c.totalItem
}

func (c *Paginator) Page() int {
	return c.page
}

func (c *Paginator) Refresh() {
	c.setNumItemsPerPage()
	c.generateLimitInitAndEnd()
}
