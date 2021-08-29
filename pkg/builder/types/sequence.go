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
	"sync"
)

type Sequence struct {
	Initial int
}

func (s Sequence) GetInitialSequence() int {

	if s.Initial == 0 {
		return 1
	}

	return s.Initial
}

// Validate check struct
func (s Sequence) Validate() error {
	return nil
}

type SequenceBuilder struct {
	sequence int
	mx       sync.Mutex
}

func NewSequenceBuilder() *SequenceBuilder {
	return &SequenceBuilder{}
}

func (b *SequenceBuilder) Initiate(config Config) error {
	b.mx.Lock()
	defer b.mx.Unlock()

	if err := config.Validate(); err != nil {
		return err
	}

	sequencer, exist := config.(Sequence)

	if !exist {
		return errors.New("invalidate type")
	}

	b.sequence = sequencer.GetInitialSequence()

	return nil
}

// Validate check struct
func (b *SequenceBuilder) Validate() error {
	return nil
}

func (b *SequenceBuilder) Build(args ...string) string {
	b.mx.Lock()
	defer b.mx.Unlock()

	b.sequence++

	return fmt.Sprintf("%d", b.sequence)
}
