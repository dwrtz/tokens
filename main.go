package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkoukk/tiktoken-go"
)

func printHelp() {
	fmt.Print(`Usage: tokens [options] [file]

This command reads text from either:
  - A specified file, if provided.
  - Standard input, if no file is provided.

It then prints the number of tokens according to the chosen encoding.

If no input is provided, this help message is displayed.

Options:
  -e, --encoding    Specify the encoding to use (default: cl100k_base)
`)
}

func main() {
	var encodingName string
	flag.StringVar(&encodingName, "e", "cl100k_base", "Encoding to use (e.g. cl100k_base, r50k_base, p50k_base, etc.)")
	flag.StringVar(&encodingName, "encoding", "cl100k_base", "Encoding to use (e.g. cl100k_base, r50k_base, p50k_base, etc.)")
	flag.Usage = printHelp
	flag.Parse()

	args := flag.Args()

	var input []byte
	var err error

	switch {
	case len(args) == 1:
		// Read from the specified file
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Failed to open file %s: %v", filename, err)
		}
		defer file.Close()

		input, err = io.ReadAll(file)
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", filename, err)
		}

	case len(args) > 1:
		// Too many arguments
		printHelp()
		os.Exit(0)

	default:
		// No file provided, try reading from stdin
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			// Stdin is from a terminal, no input is piped or redirected
			printHelp()
			os.Exit(0)
		}

		reader := bufio.NewReader(os.Stdin)
		var lines []byte
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil && err != io.EOF {
				log.Fatalf("error reading stdin: %v", err)
			}
			lines = append(lines, line...)
			if err == io.EOF {
				break
			}
		}
		input = lines
	}

	// If no input provided, print help
	if len(input) == 0 {
		printHelp()
		os.Exit(0)
	}

	enc, err := tiktoken.GetEncoding(encodingName)
	if err != nil {
		log.Fatalf("Failed to get encoding %q: %v", encodingName, err)
	}

	tokens := enc.Encode(string(input), nil, nil)
	fmt.Println(len(tokens))
}
