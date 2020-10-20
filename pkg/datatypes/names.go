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
	"errors"
	"github.com/icrowley/fake"
)

const (
	FIRST_NAME = "first_name"
	LAST_NAME  = "last_name"
	FULLNAME   = "fullname"
)

type HumanNameBuilder struct {
	mode      string
	generator func() string
}

func NewHumanNameBuilder() *HumanNameBuilder {
	return &HumanNameBuilder{}
}

func (h *HumanNameBuilder) Validate() error {

	switch h.mode {
	case FIRST_NAME:
		h.generator = fake.FirstName

		return nil
	case LAST_NAME:
		h.generator = fake.LastName

		return nil
	case FULLNAME:
		h.generator = fake.FullName

		return nil
	default:
		return errors.New("not valid argument")
	}
}

func (h *HumanNameBuilder) Initiate(config map[string]string) error {

	mode, exist := config["mode"]

	if !exist {
		mode = FULLNAME
	}

	h.mode = mode

	return nil
}

func (h *HumanNameBuilder) Build(s ...string) string {

	return h.generator()
}
