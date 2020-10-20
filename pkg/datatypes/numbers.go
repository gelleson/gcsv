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

package datatypes

import (
	"github.com/icrowley/fake"
	"strconv"
)

type numbersType string

const (
	INTEGER_NUMBER = "int"
	FLOAT_NUMBER   = "float"
)

type NumberBuilder struct {
	numbersMode numbersType
	from        float64
	to          float64
}

func NewNumberBuilder(numbersMode numbersType) *NumberBuilder {
	return &NumberBuilder{numbersMode: numbersMode}
}

func (n *NumberBuilder) Initiate(config map[string]string) error {

	from, exist := config["from"]

	if exist {
		fromInt, err := strconv.Atoi(from)

		if err != nil {
			return err
		}

		n.from = float64(fromInt)
	} else {
		n.from = 100_000_000
	}

	to, exist := config["from"]

	if exist {
		toInt, err := strconv.Atoi(to)

		if err != nil {
			return err
		}

		n.to = float64(toInt)

	} else {
		n.to = 0
	}

	return nil
}

func (n *NumberBuilder) Build(args ...string) string {
	return fake.Digits()
}

func (n *NumberBuilder) Validate() error {
	return nil
}
