package main

import (
	"fmt"

	"github.com/christoffer1009/perceptron/perceptron"
	"github.com/christoffer1009/perceptron/review"
)

func main() {
	reviews := review.GetReviewsFromCsv()
	p := perceptron.NewPerceptron(reviews, 1.0, 0.001)
	p.Init()
	for i := 0; i < 1000; i++ {
		p.Train()
		fmt.Printf("Iter %d\n", i)
	}

	count := 0
	for i := 0; i < len(p.Reviews); i++ {
		pred := p.Predict(p.Reviews[i])
		if pred != p.Reviews[i].Class {
			count++
		}
	}
	fmt.Printf("Perc: %f%%\n", float64(count)/float64(len(p.Reviews))*100)
}
