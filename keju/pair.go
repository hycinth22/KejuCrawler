package keju

type QAPair struct {
	Question string
	Answer   string
}

// related to question & answer
func (pair *QAPair) HashCode() Hash {
	return NewMD5HashFromString(pair.Question + pair.Answer)
}
