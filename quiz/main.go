package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	question string
	result   string
}

var score = 0

func main() {
	filename := flag.String("csv", "problems.csv", "The full path to csv file.")
	timeLimit := flag.Int("limit", 40, "The quiz limit time in seconds.")

	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		fmt.Printf("File is not readable: %q", err.Error())
		os.Exit(0)
	}

	quiz := getQuizFromFile(file)

	//Init channels.
	t := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	c := make(chan string)

quizLoop:
	for _, q := range quiz {
		go askQuestion(q, c)

		select {
		case <-t.C:
			fmt.Println("\nTimes up!")
			break quizLoop
		case r := <-c:
			if r == q.result {
				score++
			}
		}
	}

	fmt.Println("You scored", score, "out of", len(quiz))
}

func getQuizFromFile(file *os.File) []Quiz {
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("error", err.Error())
		os.Exit(1)
	}

	quiz := make([]Quiz, len(lines))

	for i, l := range lines {
		quiz[i] = Quiz{
			question: l[0],
			result:   strings.TrimSpace(l[1]),
		}
	}

	return quiz
}

func askQuestion(q Quiz, c chan string) {
	var answer string
	fmt.Printf("what %q, sir?", q.question)
	fmt.Scanf("%s\n", &answer)

	c <- answer
}
