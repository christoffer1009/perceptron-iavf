package review

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Review struct {
	ID     int
	Text   string
	Class  float64
	Tokens []string
}

func NewReview(id int, text string, class float64) *Review {
	return &Review{
		ID:     id,
		Text:   text,
		Class:  class,
		Tokens: getTokens(text),
	}
}

func (r *Review) ToString() string {
	return fmt.Sprintf("ID: %d\nClass: %f\nText: %s\n", r.ID, r.Class, r.Text)
}

func getTokens(str string) []string {
	tokens := strings.Fields(str)
	stopWords := getStopWordsFromTxt()
	tokens = removeNonAlpha(tokens)
	tokens = removeStopWords(tokens, stopWords)
	return tokens
}

func GetReviewsFromCsv(filename string) []*Review {
	// Abrir o arquivo CSV
	file, err := os.Open(fmt.Sprintf("./data/%s.csv", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Criar um leitor CSV
	reader := csv.NewReader(file)

	// Ler as colunas do cabeçalho
	cols, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Mapear as colunas para os índices
	mapCols := make(map[string]int)
	for i, col := range cols {
		mapCols[col] = i
	}

	count := 0
	reviews := []*Review{}

	// Ler as linhas do CSV
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// Mapear os valores da linha para os campos da struct
		str, _ := strconv.ParseFloat(row[mapCols["polarity"]], 64)
		r := NewReview(count, row[mapCols["review_text"]], str)
		reviews = append(reviews, r)
		count++
	}
	return reviews
}

func removeStopWords(words []string, stopWords []string) []string {
	var filtered []string

	for _, word := range words {
		if !Includes(stopWords, word) {
			filtered = append(filtered, word)
		}
	}

	return filtered
}

func removeNonAlpha(str []string) []string {
	regex := regexp.MustCompile("[^a-zA-Z ]")

	var filteredList []string

	for _, text := range str {
		filteredText := regex.ReplaceAllString(text, "")
		filteredText = strings.TrimSpace(filteredText)
		filteredList = append(filteredList, filteredText)
	}

	return filteredList
}

func Includes(list []string, val string) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}

func GetPositiveWordsFromTxt() []string {
	positiveWords := []string{}
	file, err := os.Open("./data/positive_words_en.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		positiveWords = append(positiveWords, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return positiveWords
}

func GetNegativeWordsFromTxt() []string {
	negativeWords := []string{}
	file, err := os.Open("./data/negative_words_en.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		negativeWords = append(negativeWords, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return negativeWords
}

func getStopWordsFromTxt() []string {
	stopWords := []string{}
	file, err := os.Open("./data/stopwords_en.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		stopWords = append(stopWords, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stopWords
}
