package keju

type QAPairCollection struct {
	pairs []QAPair
}

func (c *QAPairCollection) Insert(pair QAPair) {
	c.pairs = append(c.pairs, pair)
}

func (c *QAPairCollection) Range(f func(pair *QAPair)) {
	for _, pair := range c.pairs {
		f(&pair)
	}
}

func (c *QAPairCollection) Size() int {
	return len(c.pairs)
}
