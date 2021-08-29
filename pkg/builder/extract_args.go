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

package builder

import (
	"errors"
	"github.com/gelleson/gcsv/pkg/builder/types"
	"github.com/mitchellh/mapstructure"
)

// ExtractArgs function to map to types builder which implement types.Config
// e.g. if typeObj = types.INT
// it should map decoded to builder that implement types.Config interface
func ExtractArgs(typeObj types.TYPE, decoded map[string]interface{}) (types.Config, error) {

	switch typeObj {
	case types.INT, types.FLOAT:
		obj := types.Number{}

		if err := mapstructure.Decode(decoded, &obj); err != nil {
			return nil, err
		}

		return obj, nil

	case types.SEQ:
		obj := types.Sequence{}

		if err := mapstructure.Decode(decoded, &obj); err != nil {
			return nil, err
		}

		return obj, nil

	case types.PERSONAL:
		obj := types.Human{}

		if err := mapstructure.Decode(decoded, &obj); err != nil {
			return nil, err
		}

		return obj, nil

	case types.DATE:
		obj := types.Date{}

		if err := mapstructure.Decode(decoded, &obj); err != nil {
			return nil, err
		}

		return obj, nil

	default:
		return nil, errors.New("doesn't support type")
	}
}
