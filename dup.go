package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/labbsr0x/goh/gohtypes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var _path = "path"
var _lines = map[string]bool{}
var _duplicates = map[string]bool{}
var _ndups = 0

var rootCMD = &cobra.Command{
	Use:   "dup",
	Short: "An extremely simple utility to check if the input have duplicate lines for lines with a maximum of 65536 characters",
	RunE:  Run,
}

func init() {
	cobra.OnInitialize(func() {
		viper.SetEnvPrefix("DUP") // all
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
		viper.AutomaticEnv() // read in environment variables that matchs
	})

	flags := rootCMD.Flags()
	flags.StringP(_path, "p", "", "The file path to read from")
	gohtypes.PanicIfError("Unable to bind viper flags", 500, viper.GetViper().BindPFlags(rootCMD.Flags()))
}

func main() {
	go PrintResultsOnOSSignals()
	rootCMD.Execute()
}

// Run defines what should happen when the user runs 'dup'
func Run(cmd *cobra.Command, args []string) (err error) {
	path := viper.GetString(_path)
	if len(path) == 0 {
		STDIN()
	} else {
		FILE(path)
	}
	return nil
}

// STDIN gets the input from STDIN
func STDIN() {
	fmt.Println("Type bellow your lines:")

	FindDuplicateLines(os.Stdin)
}

// FILE gets the input from a file
func FILE(path string) {
	file, err := os.Open(path)
	defer file.Close()
	gohtypes.PanicIfError("Unable to read file", 500, err)

	FindDuplicateLines(file)

	PrintResults()
}

// FindDuplicateLines finds the duplicate lines from a *os.File object
func FindDuplicateLines(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matched, _ := regexp.Match(`[^\s\t\n\0]{1}`, []byte(line))
		if matched && _lines[line] {
			_duplicates[line] = true
			_ndups++
		}
		_lines[line] = true
	}
}

// PrintResultsOnOSSignals prints the results after program ends by an Os Signal
func PrintResultsOnOSSignals() {
	stopCh := make(chan os.Signal)

	signal.Notify(stopCh, syscall.SIGTERM)
	signal.Notify(stopCh, syscall.SIGINT)

	<-stopCh

	fmt.Printf("\n\n\n\n\n\n")
	os.Exit(PrintResults())
}

// PrintResults prints the results
func PrintResults() int {
	fmt.Printf("Found %v duplicates!\n\n", len(_duplicates))

	for dup := range _duplicates {
		fmt.Println(dup)
	}

	return len(_duplicates)
}
