/**
ID успешной посылки: 55187308

АЛГОРИТМ

Поисковый индекс содержит структуру:
слово1 => [
	{индекс документа, сколько раз встречается},
	{индекс документа, сколько раз встречается},
]
слово2 => [
	{индекс документа, сколько раз встречается},
	{индекс документа, сколько раз встречается},
]

Далее парсим входящий поисковый запрос, присваивая каждому документу приоритет в выдаче:
по каждому слову в запросе обращаемся в поисковый индекс, и увеличиваем docIndex на count


ВЫЧИСЛИТЕЛЬНАЯ СЛОЖНОСТЬ

Добавить документ в поисковый индекс:
O(M) – где M длина текста документа

Сформировать весь поисковый индекс:
O(L) - где L суммарная длина текста в N входящих документах

Получить результаты поиска по запросу:
O(L) – чтобы распарсить запрос по словам, где L – длина поискового запроса
O(N * M * L)– где N - кол-во слов в поисковом запросе, M – кол-во слов в индексе, L - кол-во повторений слова в индексе


ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ
O(M * L) - для поискового индекса – где M – суммарное кол-во слов во всех документах, L - кол-во повторений слова
O(L) – для текущего поискового запроса – где L - длина поискового запроса
Кэш хранит только результаты 1 последнего запроса, поэтому O(1)
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
docIndex – индекс документа
count – сколько раз встречается
*/
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

/**
слово => [
	{индекс документа, сколько раз встречается},
	{индекс документа, сколько раз встречается},
]
*/
type SearchIndex map[string][]SearchHit

/**
docs – слайс всех документов
index – поисковый индекс
cache – кешируем последний поисковый запрос
*/
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

/**
Добавить документ в поисковый индекс, распарсив по словам
слово1 => [
	{индекс добавляемого документа, сколько раз встречается},
]
слово2 => [
	{индекс добавляемого документа, сколько раз встречается},
]
*/
func (s *SearchEngine) addDoc(doc string) {
	s.docs = append(s.docs, doc)
	docIndex := len(s.docs) - 1
	wordCounter := make(map[string]int)
	for _, word := range strings.Fields(doc) {
		wordCounter[word]++
	}
	for word, count := range wordCounter {
		s.index[word] = append(s.index[word], SearchHit{docIndex, count})
	}
}

/**
Вернуть результаты поиска по запросу
*/
func (s *SearchEngine) query(q string) SearchResults {
	// есть ли поисковой запрос в кэше?
	if cached, ok := s.cache[q]; ok {
		return cached
	}
	// максимальное кол-во результатов в выдаче
	limit := 5
	// очищаем кэш
	s.cache = make(map[string]SearchResults)
	// индекс документа => вес в выдаче
	counter := make(map[int]int)
	// не учитываем дубли в запросе
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
	// сортируем по весу desc, индекс документа asc
	sort.Slice(sr, func(i, j int) bool {
		if sr[i].count != sr[j].count {
			return sr[i].count > sr[j].count
		}
		return sr[i].docIndex < sr[j].docIndex
	})
	if len(sr) < limit {
		limit = len(sr)
	}
	// кэшируем последний запрос
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
