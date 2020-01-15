package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n int
)

func usage() {
	fmt.Fprintf(os.Stderr, "mnemo is a education tool for memorize words and idioms\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	show := flag.NewFlagSet("show", flag.ExitOnError)
	show.IntVar(&n, "n", 10, "")
	//i := show.String("i", "input.txt", "")

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "show":
		show.Parse(os.Args[2:])
		file, err := os.Open(`./input.txt`)
		if err != nil {
			fmt.Println(err)
			// need to fix it
			os.Exit(1)
		}
		defer file.Close()

		sc := bufio.NewScanner(file)
		for i := 1; sc.Scan(); i++ {
			if err := sc.Err(); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(sc.Text())
		}
	}

}
