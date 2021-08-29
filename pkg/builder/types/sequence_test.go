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
	"strconv"
	"testing"
)

type SeqSuite struct {
	suite.Suite
	sequenceBuilder SequenceBuilder
}

func (s *SeqSuite) SetupTest() {
	s.sequenceBuilder = *NewSequenceBuilder()
}

func (s *SeqSuite) TestInitial() {
	err := s.sequenceBuilder.Initiate(Sequence{
		Initial: 2,
	})

	s.Assert().Nil(err)
}

func (s *SeqSuite) TestBuild() {
	err := s.sequenceBuilder.Initiate(Sequence{
		Initial: 2,
	})

	s.Assert().Nil(err)

	n1 := s.sequenceBuilder.Build()

	number, err := strconv.Atoi(n1)

	s.Assert().Nil(err)

	s.Assert().Equal(3, number)

	n1 = s.sequenceBuilder.Build()

	number, err = strconv.Atoi(n1)

	s.Assert().Nil(err)

	s.Assert().Equal(4, number)
}

func TestSeq(t *testing.T) {
	suite.Run(t, new(SeqSuite))
}
