/*
 * Copyright (c) 2021. gelleson
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
	"github.com/stretchr/testify/suite"
	"testing"
)

type DateSuite struct {
	suite.Suite
	dateBuilder DateBuilder
}

func (s *DateSuite) SetupTest() {
	s.dateBuilder = *NewDateBuilder()
}

func (s DateSuite) TestInitial() {
	err := s.dateBuilder.Initiate(Sequence{
		Initial: 2,
	})

	s.Assert().Error(err)

	err = s.dateBuilder.Initiate(Date{
		From: "2001-01-01",
		To:   "2010-10-01",
	})

	s.Assert().Nil(err)
}

func (s DateSuite) TestBuild() {
	err := s.dateBuilder.Initiate(Date{
		From: "2001-01-01",
		To:   "2010-10-01",
	})

	s.Assert().Nil(err)

	n1 := s.dateBuilder.Build()

	s.Assert().NotZero(n1)
}

func TestDate(t *testing.T) {
	suite.Run(t, new(DateSuite))
}
