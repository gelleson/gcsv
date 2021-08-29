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

package types

import (
	"errors"
	"fmt"
	"github.com/icrowley/fake"
	"math/rand"
)

type Number struct {
	From float64 `yaml:"from" json:"from"`
	To   float64 `yaml:"to" json:"to"`
}

func (n Number) Validate() error {

	if (n.From != 0 && n.To != 0) && n.From >= n.To {
		return errors.New("from should be lower than to")
	}

	return nil
}

type NumberMode string

const (
	IntegerNumberMode NumberMode = "int"
	FloatNumberMode   NumberMode = "float"
)

type NumberBuilder struct {
	numbersMode NumberMode
	from        float64
	to          float64
}

func NewNumberBuilder(numbersMode NumberMode) *NumberBuilder {
	return &NumberBuilder{numbersMode: numbersMode}
}

func (n *NumberBuilder) Initiate(config Config) error {

	if err := config.Validate(); err != nil {
		return err
	}

	numbers, exist := config.(Number)

	if !exist {
		return errors.New("invalidate type")
	}

	n.to = numbers.To

	if n.to == 0 {
		n.to = 100_000_000
	}

	n.from = numbers.From

	return nil
}

func (n NumberBuilder) Build(args ...string) string {

	switch n.numbersMode {
	case IntegerNumberMode:
		return fmt.Sprintf("%v", rand.Intn(int(n.to-n.from+1))+int(n.from))
	case FloatNumberMode:
		return fmt.Sprintf("%v", n.from+rand.Float64()*(n.to-n.from))
	default:
		return fake.Digits()
	}
}

func (n NumberBuilder) Validate() error {
	return nil
}
