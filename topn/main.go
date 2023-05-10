package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/bitfield/script"
)

type Args struct {
	Path    string
	Numbers int
}

func PrintDefaultArgs() {
	fmt.Printf("Usage: topN\n\n")
	fmt.Printf("Output the largest N numbers, highest first.\n\n")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func CmdArgs() Args {
	pathStr := flag.String("p", "", "file path")
	nInt := flag.Int("n", 10, "largest N numbers")
	flag.Parse()

	return Args{
		Path:    *pathStr,
		Numbers: *nInt,
	}
}

func ValidateCmdArgs(args Args) {
	_, err := script.IfExists(args.Path).String()
	switch {
	case args.Path == "":
		fmt.Printf("ERROR: file path flag is empty\n\n")
		PrintDefaultArgs()
	case err != nil:
		fmt.Printf("ERROR: file path does not exist\n\n")
		PrintDefaultArgs()
	case args.Numbers <= 0:
		fmt.Printf("ERROR: largest N numbers is <= 0\n\n")
		PrintDefaultArgs()
	default:
		break
	}
}

func FileHandler(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

type NumberSet struct {
	highest [10]int
}

func (n *NumberSet) ReplaceWhenGreaterThan(num int) {
	for i, v := range n.highest {
		if num < v {
			continue
		}
		n.highest[i] = num
		break
	}
}

func (n *NumberSet) PrintHighest() {
	s := n.highest[:]
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	for _, v := range s {
		fmt.Println(v)
	}
}

func ParseFile(handler *os.File) NumberSet {
	ns := NumberSet{}
	scanner := bufio.NewScanner(handler)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		ns.ReplaceWhenGreaterThan(i)
	}
	return ns
}

func main() {
	args := CmdArgs()
	ValidateCmdArgs(args)

	handler := FileHandler(args.Path)
	defer handler.Close()

	ns := ParseFile(handler)

	ns.PrintHighest()
}
