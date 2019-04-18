package main

import (
	"fmt"
	"github.com/gelleson/generator/csv/src/generator"
	"github.com/gelleson/generator/csv/src/parser"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	var rootCmd = &cobra.Command{
		Use: "root [sub]",
	}

	var subCmd = &cobra.Command{
		Use:   "generate [FILE NAME]",
		Short: "Generate csv",
		PreRun: func(cmd *cobra.Command, args []string) {
			if  len(args) == 0 {
				fmt.Println(" ", len(args))
				os.Exit(0)
			}
			if !strings.HasSuffix(args[0], ".yaml") {
				log.Fatal("1")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.OpenFile(args[0], os.O_RDONLY, 0600)

			if err != nil {
				log.Fatal(err)
			}
			bytesByFile, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}
			var config parser.Config
			err = yaml.Unmarshal(bytesByFile, &config)
			if err != nil {
				log.Fatal(err)
			}
			parser := parser.NewParser(config)
			documents := parser.PrepareDocument()
			gen := generator.NewGenerator(documents)
			gen.Generate()
			log.Println("CSV is generated")
		},
	}

	rootCmd.AddCommand(subCmd)

	rootCmd.Execute()

}
