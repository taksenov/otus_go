/*
Copyright © 2022 taksenov@gmail.com
*/

// Package cmd -- comand line interface app.
package cmd

import (
	"log"

	goflag "flag"

	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

// Root -- rootCmd pseudo constructor.
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "",
		Short: "",
		Long: string("\033[92m") + `
▒███████████████████▒█████████████████████▒░
░▒▒ ▓░▒░▒░▒░▓░░░ ▒▒░▒ ░▓░▒ ▒ ░░ ▒░ ▒▒ ░ ░▓ ░
░░▒ ▒ ░ ▒░░ ▒░ ░ ▒░ ░  ▒ ░ ░  ░ ░  ░░   ░▒ ░
░ ░ ░ ░ ░ ░ ░    ░  ░  ░   ░    ░   ░    ░  
  ░ ░       ░░    ░      ░      ░   ░    ░  
 _____ _____ _____ _____ _____ _____ __ __ 
|_   _|     |     |     |   __|   __|  |  |
  | | |-   -| | | |  |  |   __|   __|_   _|
  |_| |_____|_|_|_|_____|__|  |_____| |_|  
                                           
 _____ _____ _____ _____ _____ _____ _____ 
|  _  |  |  |   __|   __|   | |     |  |  |
|     |    -|__   |   __| | | |  |  |  |  |
|__|__|__|__|_____|_____|_|___|_____|\___/ 

  ░ ░       ░░    ░      ░      ░   ░    ░  
░ ░ ░ ░ ░ ░ ░    ░  ░  ░   ░    ░   ░    ░  
░░▒ ▒ ░ ▒░░ ▒░ ░ ▒░ ░  ▒ ░ ░  ░ ░  ░░   ░▒ ░
░▒▒ ▓░▒░▒░▒░▓░░░ ▒▒░▒ ░▓░▒ ▒ ░░ ▒░ ▒▒ ░ ░▓ ░
▒███████████████████▒████████████OTUS█HW07▒░` + string("\033[0m") +
			`

Пример использования:
` + string("\033[92m") + `
  go run main.go --from      '/путь/к/исходному/файлу.tst'
                 --to        '/путь/к/копии/'
                 --offset     отступ в источнике (0 по умочанию)
                 --limit      кол-во копируемых байт (0 по умолчанию)
` + string("\033[0m"),

		Run: func(cmd *cobra.Command, _ []string) {
			f, err := cmd.Flags().GetString("from")
			if err != nil || f == "" {
				log.Fatal(err)
			}

			t, err := cmd.Flags().GetString("to")
			if err != nil || t == "" {
				log.Fatal(err)
			}

			o, err := cmd.Flags().GetInt64("offset")
			if err != nil {
				o = 0
			}

			l, err := cmd.Flags().GetInt64("limit")
			if err != nil {
				l = 0
			}

			Copy(f, t, o, l)
		},
	}

	var (
		from, to      string
		limit, offset int64
	)

	// NB: Для обеспечения best practice переключено на pflag, где
	// флаги-слова используют `--` вместо общепризнанных `-` для сокращений (shorthands)
	goflag.StringVar(&from, "from", "", "file to read from")
	goflag.StringVar(&to, "to", "", "file to write to")
	goflag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	goflag.Int64Var(&offset, "offset", 0, "offset in input file")

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	return rootCmd
}
