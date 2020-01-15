package main

import (
	"haimanKejuCrawler/keju"
	"os"
)

func main() {
	file, err := keju.OpenRepositoryFileForRead()
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	repo, err := keju.LoadRepository(file)
	if err != nil {
		panic(err.Error())
	}

	f, err := os.OpenFile("question.list", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		panic(err.Error())
	}
	collection := repo.ListQAPairs()
	collection.Range(func(pair *keju.QAPair) {
		f.WriteString(pair.Question)
		f.WriteString("\n")
		f.WriteString(pair.Answer)
		f.WriteString("\n")
	})
}
