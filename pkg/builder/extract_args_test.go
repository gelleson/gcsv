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

package builder

import (
	"fmt"
	"github.com/gelleson/gcsv/pkg/builder/types"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

var test = strings.NewReader(`
---
a: Easy!
b:
  value: 1
  name: 2
 
`)

type TestModel struct {
	Value int `mapstructure:"value"`
	Name  int `mapstructure:"name"`
	Tags  []string
}

type mainStruct struct {
	A string
	B interface{}
}

func TestDecoder(t *testing.T) {

	m := &mainStruct{}

	err := yaml.NewDecoder(test).Decode(m)
	fmt.Println(err)

	v := &TestModel{}

	err = mapstructure.Decode(m.B, v)
	fmt.Println(err)

	fmt.Println(m)
	fmt.Println(*v)
}

func TestExtractArgs(t *testing.T) {
	case1 := strings.NewReader(`
initial: 25
`)

	obj1 := map[string]interface{}{}

	err := yaml.NewDecoder(case1).Decode(&obj1)

	assert.Nil(t, err)

	args, err := ExtractArgs(types.SEQ, obj1)

	assert.Nil(t, err)

	seq, exist := args.(types.Sequence)

	if !exist {
		assert.Fail(t, "cannot to cast Sequence")
		return
	}

	assert.Equal(t, 25, seq.Initial)

	case2 := strings.NewReader(`
from: 25
to: 14
`)

	obj2 := map[string]interface{}{}

	err = yaml.NewDecoder(case2).Decode(&obj2)

	assert.Nil(t, err)

	args, err = ExtractArgs(types.INT, obj2)

	assert.Nil(t, err)

	num, exist := args.(types.Number)

	if !exist {
		assert.Fail(t, "cannot to cast Number")
		return
	}

	assert.Equal(t, float64(25), num.From)
	assert.Equal(t, float64(14), num.To)

	case3 := strings.NewReader(`
from: 25
to: 14
`)
	obj3 := map[string]interface{}{}

	err = yaml.NewDecoder(case3).Decode(&obj3)

	assert.Nil(t, err)

	args, err = ExtractArgs(types.FLOAT, obj3)

	assert.Nil(t, err)

	num2, exist := args.(types.Number)

	if !exist {
		assert.Fail(t, "cannot to cast Number")
		return
	}

	assert.Equal(t, float64(25), num2.From)
	assert.Equal(t, float64(14), num2.To)

	case4 := strings.NewReader(`
mode: first_name
`)
	obj4 := map[string]interface{}{}

	err = yaml.NewDecoder(case4).Decode(&obj4)

	assert.Nil(t, err)

	args, err = ExtractArgs(types.PERSONAL, obj4)

	assert.Nil(t, err)

	hum, exist := args.(types.Human)

	if !exist {
		assert.Fail(t, "cannot to cast Human")
		return
	}

	assert.Equal(t, types.FirstNameMode, hum.Mode)

	case5 := strings.NewReader(`
from: 2001-01-01
to: 2021-01-11
format: 2020-11-12
`)

	obj5 := map[string]interface{}{}

	err = yaml.NewDecoder(case5).Decode(&obj5)

	assert.Nil(t, err)

	args, err = ExtractArgs(types.DATE, obj5)

	assert.Nil(t, err)

	date, exist := args.(types.Date)

	if !exist {
		assert.Fail(t, "cannot to cast Date")
		return
	}

	assert.Equal(t, "2001-01-01", date.From)
	assert.Equal(t, "2021-01-11", date.To)
	assert.Equal(t, "2020-11-12", date.Format)
}
