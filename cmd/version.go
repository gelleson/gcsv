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

package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var (
	tag       = ""
	commitStr = ""
	dateStr   = ""
)

// SetVersion save arguments to vars
func SetVersion(version, commit, date string) {
	tag = version
	commitStr = commit
	dateStr = date
}

var version = &cobra.Command{
	Use:     "version",
	Short:   "Reflect build version, hash",
	Aliases: []string{"v", "ver"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		fmt.Println("")

		myFigure := figure.NewColorFigure("GCSV", "", "blue", true)
		myFigure.Print()

		fmt.Println("")
		fmt.Println("")
		fmt.Println("Version: ", tag)
		fmt.Println("Commit: ", commitStr)
		fmt.Println("Build Date: ", dateStr)
		fmt.Println("")
		fmt.Println("")

	},
}
