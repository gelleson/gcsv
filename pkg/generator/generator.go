/*
 * Copyright (c) 2020. gelleson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package generator

import (
	"encoding/csv"
	"fmt"
	"github.com/gelleson/gcsv/pkg/builder"
	"log"
	"os"
)

type Generator struct {
	documents []Document
}

func NewGenerator(documents []Document) *Generator {
	return &Generator{documents: documents}
}

func (g Generator) generate() error {
	for _, doc := range g.documents {
		document, err := os.OpenFile(fmt.Sprintf("%s.csv", doc.Name), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		cv := csv.NewWriter(document)
		cv.Comma = ','

		records, err := g.newRecords(doc)

		if err != nil {
			return err
		}

		err = cv.WriteAll(records)

		if err != nil {
			return err
		}
	}

	return nil
}

func (g Generator) newRecords(document Document) ([][]string, error) {

	result := make([][]string, 0)

	if document.WithHeader {
		headers := make([]string, 0)
		for _, column := range document.Columns {
			headers = append(headers, column.Name)
		}

		result = append(result, headers)
	}

	builders := make([]Builder, 0)

	for _, value := range document.Columns {
		field := value.Field

		builderInstance, err := builder.Factory(field.Type)

		if err != nil {
			return nil, err
		}

		if err := builderInstance.Initiate(value.Kwargs); err != nil {
			return nil, err
		}

		if err := builderInstance.Validate(); err != nil {
			return nil, err
		}

		builders = append(builders, builderInstance)
	}

	for i := 1; i < document.Rows; i++ {
		record := make([]string, 0)

		for _, builderInstance := range builders {
			generatedValue := builderInstance.Build()

			record = append(record, generatedValue)
		}

		result = append(result, record)
	}

	return result, nil
}

func (g Generator) Generate() error {
	return g.generate()
}
