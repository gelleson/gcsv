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
	"github.com/araddon/dateparse"
	"math/rand"
	"time"
)

const defaultFormat = "2006-01-02"

// Date is strategy to generate date related column
type Date struct {
	// From is date time which should start from this date
	From string
	// To is last range of the possible date
	To string
	// Format is means result of the date generation.
	// Also, From and To fields are validated by Format
	Format string
}

// GetFormat is returns defaultFormat or Format field
func (d Date) GetFormat() string {

	if d.Format == "" {
		return defaultFormat
	}

	return d.Format
}

// Validate is check struct
func (d Date) Validate() error {

	_, err := time.Parse(d.GetFormat(), d.To)

	if err != nil {
		return err
	}

	_, err = time.Parse(d.GetFormat(), d.From)

	if err != nil {
		return err
	}

	return nil
}

// DateBuilder is generate dates
type DateBuilder struct {
	format string
	from   string
	to     string

	fromTime time.Time
	toTime   time.Time
}

// NewDateBuilder is constructor of the DateBuilder
func NewDateBuilder() *DateBuilder {
	return &DateBuilder{}
}

func randate(since, until time.Time) time.Time {
	min := since.Unix()
	max := until.Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

// Initiate is need to init struct
func (d *DateBuilder) Initiate(config Config) error {

	if err := config.Validate(); err != nil {
		return err
	}

	format, exist := config.(Date)

	if !exist {
		return errors.New("invalidate type")
	}

	d.from = format.From
	d.to = format.To

	d.format = format.GetFormat()

	return nil
}

// Build return generated value
func (d *DateBuilder) Build(args ...string) string {

	from := time.Date(1970, time.February, 1, 1, 1, 1, 1, time.Local)
	until := time.Now()

	if !d.fromTime.IsZero() {
		from = d.fromTime
	}

	if !d.toTime.IsZero() {
		until = d.toTime
	}

	return randate(from, until).Format(d.format)
}

// Validate check internal state of the builder
func (d *DateBuilder) Validate() error {

	if d.from != "" {
		parsedTime, err := dateparse.ParseAny(d.from)

		if err != nil {
			return err
		}

		d.fromTime = parsedTime
	}

	if d.to != "" {
		parsedTime, err := dateparse.ParseAny(d.to)

		if err != nil {
			return err
		}

		d.toTime = parsedTime
	}

	return nil
}
