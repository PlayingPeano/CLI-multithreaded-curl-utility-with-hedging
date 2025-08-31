package main

import (
	"fmt"
	"strconv"
)

const (
	standardTimeout int = 15
)

type Config struct {
	Timeout int
	URLs    []string
	Help    bool
}

func parseArgs(args []string) (*Config, error) {
	config := &Config{
		Timeout: standardTimeout,
		Help:    false,
	}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			config.Help = true
		case "-t", "--timeout":
			if i+1 >= len(args) {
				return nil, fmt.Errorf("Bad usage: missing a value after -t/--timeout flag")
			}
			timeout, err := strconv.Atoi(args[i+1])
			if err != nil {
				return nil, fmt.Errorf("Bad usage: timout value must be an integer")
			}
			config.Timeout = timeout
			i++
		default:
			config.URLs = append(config.URLs, args[i])
		}
	}
	return config, nil
}

func printHelp() {
	fmt.Println("Usage: hedgedcurl [flags] URL1, [URL2...]")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -t, --timeout SECONDS  Sets timeout for all HTTP queries in seconds.")
	fmt.Println("                         By default: 15 секунд")
	fmt.Println("  -h, --help             Prints this help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  hedgedcurl -t 30 https://example.com https://example.org")
	fmt.Println("  hedgedcurl --timeout 5 https://example.com")
	fmt.Println("  hedgedcurl --help")
}

func main() {
	printHelp()

}
