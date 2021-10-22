package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SearchHit struct {
	docIndex, count int
}

type SearchResults []SearchHit

func (sr SearchResults) String() (s string) {
	for _, v := range sr {
		s += fmt.Sprintf("%d ", v.docIndex+1)
	}
	return
}

type SearchIndex map[string][]SearchHit

type SearchEngine struct {
	docs  []string
	index SearchIndex
	cache map[string]SearchResults
}

func (s *SearchEngine) init(size int) {
	s.docs = make([]string, 0, size)
	s.index = make(SearchIndex)
	s.cache = make(map[string]SearchResults)
}

func (s *SearchEngine) addDoc(doc string) {
	s.docs = append(s.docs, doc)
	docIndex := len(s.docs) - 1
	termCounter := make(map[string]int)
	for _, term := range strings.Fields(doc) {
		termCounter[term]++
	}
	for term, count := range termCounter {
		s.index[term] = append(s.index[term], SearchHit{docIndex, count})
	}
}

func (s *SearchEngine) query(q string) SearchResults {
	if sr, ok := s.cache[q]; ok {
		return sr
	}
	limit := 5
	s.cache = make(map[string]SearchResults)
	counter := make(map[int]int)
	seen := make(map[string]bool)
	for _, term := range strings.Fields(q) {
		if _, ok := seen[term]; !ok {
			seen[term] = true
			for _, v := range s.index[term] {
				counter[v.docIndex] += v.count
			}
		}
	}
	sr := make(SearchResults, 0, len(counter))
	for k, v := range counter {
		sr = append(sr, SearchHit{k, v})
	}
	sort.Slice(sr, func(i, j int) bool {
		if sr[i].count != sr[j].count {
			return sr[i].count > sr[j].count
		}
		return sr[i].docIndex < sr[j].docIndex
	})
	if len(sr) < limit {
		limit = len(sr)
	}
	s.cache[q] = sr[0:limit]
	return sr[0:limit]
}

func main() {
	var s SearchEngine
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	s.init(n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s.addDoc(scanner.Text())
	}
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < m; i++ {
		scanner.Scan()
		searchResults := s.query(scanner.Text())
		fmt.Fprintln(writer, searchResults)
	}
	writer.Flush()
}
