package game

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Quiz type: parse CSV file
type Quiz struct {
	Question string
	Answer   string
	Hits     int
	Misseds  int
}

// Quizes type: quiz list
type Quizes []Quiz

// NewQuiz initialize a empty struct
func NewQuiz() *Quiz {
	return &Quiz{}
}

// StartTimer initilize a timer
func (quiz *Quiz) StartTimer(timer *int64) {
	time.Sleep(time.Duration(*timer) * time.Second)
	fmt.Println("Time out!")
	os.Exit(1)
}

// GetFlags handles all possible flags in the Commmand-Line
func (quiz *Quiz) GetFlags() (filepath *string, timer *int64) {

	// Handling a --filepath flag
	filepath = flag.String("filepath", "", "Specifies a file with content of the questions and answers for Quiz.")

	// Handling --timer flag
	timer = flag.Int64("timer", 0, "Specifies a value to timer of the Quiz.")

	flag.Parse()

	if *filepath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return
}

// ParseCSV receive a filepath, open the file and return csv
func (quiz *Quiz) ParseCSV(filepath *string) [][]string {
	file, err := os.Open(*filepath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	parse, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	return parse
}

// StartQuiz function to parse CSV and initialize quiz game
func (quiz *Quiz) StartQuiz() {

	filepath, timer := quiz.GetFlags()
	parse := quiz.ParseCSV(filepath)

	if timer != nil {
		go quiz.StartTimer(timer)
	}

	var quizes Quizes
	for _, each := range parse {
		quiz.Question = each[0]
		quiz.Answer = each[1]

		fmt.Printf("%s?\n", quiz.Question)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// Trimming user input
		input = strings.TrimSpace(input)
		fmt.Println(input)

		if input == quiz.Answer {
			quiz.Hits++
		}

		quizes = append(quizes, *quiz)
	}

	total := len(quizes)
	quiz.Misseds = total - quiz.Hits

	fmt.Printf("Hits: %d\n", quiz.Hits)
	fmt.Printf("Misseds: %d\n", quiz.Misseds)
	fmt.Printf("Total questions: %d\n", total)
}
