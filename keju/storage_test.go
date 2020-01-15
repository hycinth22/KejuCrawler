package keju

import (
	"io/ioutil"
	"testing"
)

func TestQAPairRepository_Persist(t *testing.T) {
	repo := NewQAPairRepository()
	repo.SaveQAPair(QAPair{
		Question: "Question0",
		Answer:"Answer0",
	})
	err := repo.Persist(ioutil.Discard)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
