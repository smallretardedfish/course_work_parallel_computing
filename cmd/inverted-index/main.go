package main

import (
	"bufio"
	"fmt"
	utils "github.com/smallretardedfish/inverted-index/pkg"
	"github.com/smallretardedfish/inverted-index/pkg/inverted_index"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	FirstText = `In computer science, an inverted index (also referred to as a postings list, postings file,
or inverted file) is a database index storing a mapping from content, such as words or numbers,
to its locations in a table, or in a document or a set of documents
(named in contrast to a forward index, which maps from documents to content).
The purpose of an inverted index is to allow fast full-text searches,
at a cost of increased processing when a document is added to the database.
The inverted file may be the database file itself, rather than its index. 
It is the most popular data structure used in document retrieval systems`
	SecondText = "THIS IS computer science inverted index"
)

var stringSources = []inverted_index.StringSource{{
	Name: "First",
	Text: FirstText,
}, {
	Name: "Second",
	Text: SecondText,
}}

func run() error {
	var args []string
	if len(os.Args) == 0 {
		args = append(args, "1", ".")
	} else if len(os.Args) == 1 {
		args = append(args, ".")
	}

	args = os.Args[1:]
	num := args[0]
	n, err := strconv.Atoi(num)
	if err != nil {
		return err
	}

	dir := args[1]
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var fileSources []string
	for _, entry := range dirEntries {
		fileSources = append(fileSources, filepath.Join("data", entry.Name()))
	}

	invIndex := inverted_index.NewMapInvertedIndex(inverted_index.FileSourceType)

	var (
		e error
	)

	t := utils.EstimateExecutionTime(func() {
		if err := invIndex.Build(n, fileSources); err != nil {
			e = err
		}
	})
	if e != nil {
		return e
	}
	log.Println(t)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(">> ")
	for scanner.Scan() {
		word := scanner.Text()
		res := invIndex.Search(word)
		fmt.Printf("word: %s - sources: %s\n>> ", word, strings.Join(res, ","))
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
