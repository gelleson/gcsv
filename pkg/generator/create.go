package generator

import (
	"encoding/csv"
	"fmt"
	"github.com/icrowley/fake"
	"log"
	"math/rand"
	"os"
	"time"
)

type Generator struct {
	documents []Document
}

func NewGenerator(documents []Document) *Generator {
	return &Generator{documents: documents}
}

func (g Generator) generate() {
	for _, doc := range g.documents {
		document, err := os.OpenFile(fmt.Sprintf("%s.csv", doc.Name), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		cv := csv.NewWriter(document)
		cv.Comma = ','
		err = cv.WriteAll(g.newRecords(doc))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (g Generator) newRecords(document Document) [][]string {
	result := make([][]string, document.Rows)
	for i := 0; i < document.Rows; i++ {
		record := make([]string, 0)
		for _, value := range document.Columns {
			field := value.Field
			if containOption(SEQ, value.Field.Option) && value.Field.Type == INT {
				record = append(record, fmt.Sprintf("%d", i))
				continue
			} else if containOption(SEQ, field.Option) {
				log.Fatal("you can use seq only with int")
			}

			switch field.Type {
			case DATE:
				t := time.Date(fake.Year(1999, 2019), time.Month(fake.MonthNum()), fake.Day(), 0, 0, 0, 0, time.UTC).String()
				record = append(record, t)
			case STRING:
				if containOption(NAME, field.Option) {
					record = append(record, fake.FirstName())
				} else if containOption(LAST_NAME, field.Option) {
					record = append(record, fake.LastName())
				} else {
					record = append(record, fake.Word())
				}
			case FLOAT:
				record = append(record, fmt.Sprintf("%f", rand.Int63()))
			case INT:
				record = append(record, fmt.Sprintf("%d", rand.Int()))
			}

		}
		result[i] = record
	}
	return result
}

func (g Generator) Generate() {
	g.generate()
}

func containOption(term OPTION, arr []OPTION) bool {
	for _, value := range arr {
		if term == value {
			return true
		}

	}
	return false
}
