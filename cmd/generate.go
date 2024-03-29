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
	"github.com/gelleson/gcsv/pkg/generator"
	"github.com/gelleson/gcsv/pkg/parser"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

var (
	verboseFlag bool
)

var generate = &cobra.Command{
	Use:     "generate",
	Short:   "Generate csv",
	Aliases: []string{"g", "gen"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(" ", len(args))
			os.Exit(1)
		}
		if !strings.HasSuffix(args[0], ".yaml") && !strings.HasSuffix(args[0], ".yml") {
			fmt.Println(fmt.Sprintf("Not valid file extension. it's should be .yml or .yaml"))
			os.Exit(1)

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger := logrus.New()

		if verboseFlag {
			logger.SetLevel(logrus.DebugLevel)
		} else {
			logger.SetLevel(logrus.InfoLevel)
		}

		file, err := os.OpenFile(args[0], os.O_RDONLY, 0600)

		if err != nil {
			logger.Fatal(err)
		}
		bytesByFile, err := ioutil.ReadAll(file)
		if err != nil {
			logger.Fatal(err)
		}

		var config parser.Config
		err = yaml.Unmarshal(bytesByFile, &config)

		if err != nil {
			logger.Fatal(err)
		}

		parser := parser.NewParser(config, logger.WithField("context", "parser"))
		documents := parser.PreparedDocument()

		gen := generator.NewGenerator(documents, logger.WithField("context", "generator"))

		if err = gen.Generate(); err != nil {
			logger.Fatal(err)
		}

		fmt.Println("")
		fmt.Println("CSV is generated")
		fmt.Println("")

	},
}

func init() {
	generate.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "")
}
