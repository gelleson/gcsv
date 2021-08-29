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
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
)

type Generator struct {
	documents []Document
	logger    *logrus.Entry
}

func NewGenerator(documents []Document, logger *logrus.Entry) *Generator {
	return &Generator{documents: documents, logger: logger}
}

func (g Generator) generate() error {

	g.logger.Debugln("Generation of the documents started")

	g.logger.Debugf("Total documents is %d", len(g.documents))

	for _, doc := range g.documents {
		document, err := os.OpenFile(fmt.Sprintf("%s.csv", doc.Name), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			g.logger.Error(err)
			return err
		}

		g.logger.Debugf("Created file %s.csv", doc.Name)

		cv := csv.NewWriter(document)
		cv.Comma = ','

		records, err := g.newRecords(doc)

		if err != nil {
			g.logger.Error(err)
			return err
		}

		g.logger.Debugf("Generated total records is %d", len(records))

		err = cv.WriteAll(records)

		if err != nil {
			g.logger.Error(err)
			return err
		}

		g.logger.Debugf("Writed total records is %d", len(records))
	}

	return nil
}

func (g Generator) newRecords(document Document) ([][]string, error) {

	result := make([][]string, 0)

	if document.WithHeader {

		g.logger.Debugln("Generate csv with header")

		headers := make([]string, 0)
		for _, column := range document.Columns {
			headers = append(headers, column.Name)
		}

		g.logger.Debugf("Total headers are %d", len(headers))

		result = append(result, headers)
	}

	builders := make([]builder.Builder, 0)

	for _, value := range document.Columns {
		field := value.Field

		builderInstance, err := builder.Factory(field.Type)

		if err != nil {
			return nil, err
		}

		g.logger.Debugf("Generate %s column with %v type", value.Name, field.Type)

		args, err := builder.ExtractArgs(field.Type, value.Kwargs)

		if err != nil {
			return nil, err
		}

		g.logger.Debugf("Extract Args with type %v", reflect.TypeOf(args).Name())

		if err := builderInstance.Initiate(args); err != nil {
			return nil, err
		}

		g.logger.Debugln("Initiated Builder")

		if err := builderInstance.Validate(); err != nil {
			return nil, err
		}

		g.logger.Debugln("Validated Builder")

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
