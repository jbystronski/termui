package termui

type Navigator struct {
	MinOffset    int
	MaxOffset    int
	index        int
	PrevIndex    int
	TotalEntries int
}

func (c *Navigator) SetIndex(i int) {
	if i > c.TotalEntries-1 || i < 0 {
		return
	}

	c.index = i
}

func (c *Navigator) Index() int {
	return c.index
}

func (c *Navigator) PrevEntry() bool {
	if c.TotalEntries == 0 || c.index == 0 {
		return false
	}

	c.PrevIndex = c.index
	c.index--
	return true
}

func (c *Navigator) NextEntry() bool {
	if c.TotalEntries == 0 || c.index == c.TotalEntries-1 {
		return false
	}

	c.PrevIndex = c.index
	c.index++
	return true
}

func (c *Navigator) MovePgUp() bool {
	if c.index == 0 {
		return false
	}

	c.PrevIndex = c.index

	start := c.CalculateStartIndex(c.index)

	if c.index == start {
		c.index = c.index - 1
	} else {
		c.index = start
	}

	return true
}

func (c *Navigator) MovePgDown(visibleLines int) bool {
	if c.index == c.TotalEntries-1 {
		return false
	}

	c.PrevIndex = c.index

	end := c.CalculateEndIndex(c.index, c.TotalEntries-1)

	if c.index == end {
		c.index = c.index + 1
	} else {
		c.index = end
	}

	return true
}

func (c *Navigator) FirstEntry() (res bool) {
	if c.TotalEntries != 0 || c.index != 0 {

		c.PrevIndex = c.index
		c.index = 0
		res = true
	}

	return
}

func (c *Navigator) LastEntry() (res bool) {
	if c.TotalEntries != 0 || c.index != c.TotalEntries-1 {
		c.PrevIndex = c.index

		c.index = c.TotalEntries - 1
		res = true

	}

	return
}

func (c *Navigator) Jump(from, to, maxIndex int) bool {
	if from > maxIndex || to > maxIndex {
		return false
	}

	c.PrevIndex = from
	c.index = to

	return true
}

func (c *Navigator) ShouldUpdateChunk() bool {
	if c.PrevIndex == 0 && c.index == 0 {
		return true
	}

	return c.CalculateStartIndex(c.PrevIndex) != c.CalculateStartIndex(c.index)
}

func (c *Navigator) GetIndexOffset(index int) int {
	lines := c.MaxOffset - c.MinOffset + 1
	return c.MinOffset + index%lines
}

func (c *Navigator) CalculateStartIndex(index int) int {
	return index - (c.GetIndexOffset(index) - c.MinOffset)
}

func (c *Navigator) CalculateEndIndex(index, maxIndex int) int {
	end := index + c.MaxOffset - c.GetIndexOffset(index)

	if end > maxIndex {
		return maxIndex
	}

	return end
}

func (c *Navigator) Reset() {
	c.index, c.PrevIndex = 0, 0
}

func (c *Navigator) IndiceRange(index, lastIndex int) (int, int) {
	return c.CalculateStartIndex(index), c.CalculateEndIndex(index, lastIndex)
}
