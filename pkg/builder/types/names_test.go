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

type NameSuite struct {
	suite.Suite
	nameBuilder HumanNameBuilder
}

func (s *NameSuite) SetupTest() {
	s.nameBuilder = *NewHumanNameBuilder()
}

func (s NameSuite) TestInitial() {
	err := s.nameBuilder.Initiate(Sequence{
		Initial: 2,
	})

	s.Assert().Error(err)

	err = s.nameBuilder.Initiate(Human{
		Mode: FullNameMode,
	})

	s.Assert().Nil(err)
}

func (s NameSuite) TestBuild() {
	err := s.nameBuilder.Initiate(Human{
		Mode: FirstNameMode,
	})

	s.Assert().Nil(err)

	n1 := s.nameBuilder.Build()

	s.Assert().NotZero(n1)
}

func TestName(t *testing.T) {
	suite.Run(t, new(NameSuite))
}
