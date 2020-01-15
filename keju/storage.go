package keju

import (
	"encoding/gob"
	"io"
)

type QAPairRepository map[Hash]QAPair

func NewQAPairRepository() *QAPairRepository {
	repo := QAPairRepository(make(map[Hash]QAPair, 0))
	return &repo
}

func (repository *QAPairRepository) SaveQAPair(pair QAPair) (exist bool) {
	hash := pair.HashCode()
	_, exist = (*repository)[hash]
	(*repository)[hash] = pair
	return
}

func (repository *QAPairRepository) ListQAPairs() QAPairCollection {
	var collection QAPairCollection
	for _, pair := range *repository {
		collection.Insert(pair)
	}
	return collection
}

func (repository *QAPairRepository) Persist(writer io.Writer) error {
	return gob.NewEncoder(writer).Encode(*repository)
}

func LoadRepository(reader io.Reader) (*QAPairRepository, error) {
	var repo QAPairRepository
	err := gob.NewDecoder(reader).Decode(&repo)
	return &repo, err
}
