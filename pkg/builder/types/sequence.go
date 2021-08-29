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
	"fmt"
	"strconv"
	"sync"
)

type SequenceBuilder struct {
	sequence int
	mx       sync.Mutex
}

func NewSequenceBuilder() *SequenceBuilder {
	return &SequenceBuilder{}
}

func (b *SequenceBuilder) Initiate(config map[string]string) error {
	b.mx.Lock()
	defer b.mx.Unlock()

	b.sequence = 1

	_, exist := config["initial_sequence"]

	if exist {
		sequence := config["initial_sequence"]

		sequenceInt, err := strconv.Atoi(sequence)

		if err != nil {
			return err
		}

		b.sequence = sequenceInt
	}

	return nil
}

func (b *SequenceBuilder) Validate() error {
	return nil
}

func (b *SequenceBuilder) Build(args ...string) string {
	b.mx.Lock()
	defer b.mx.Unlock()

	b.sequence++

	return fmt.Sprintf("%d", b.sequence)
}