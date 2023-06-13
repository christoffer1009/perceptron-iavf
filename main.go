package main

import (
	"fmt"

	"github.com/christoffer1009/perceptron/perceptron"
	"github.com/christoffer1009/perceptron/review"
)

func main() {
	reviews := review.GetReviewsFromCsv("reviews1")
	p := perceptron.NewPerceptron(reviews, 1.0, 0.001)
	p.Init()
	for i := 0; i < 100; i++ {
		p.Train()
		fmt.Printf("Iter %d\n", i)
	}

	reviews_test := review.GetReviewsFromCsv("test")
	p_test := perceptron.NewPerceptron(reviews_test, 1.0, 0.001)
	p_test.Init()
	p_test.Weigths = p.Weigths

	count := 0
	for i := 0; i < len(p_test.Reviews); i++ {
		pred := p_test.Predict(p_test.Reviews[i])
		if pred != p_test.Reviews[i].Class {
			count++
		}
		fmt.Printf("Review: %d Class: %f Pred: %f\n", i, reviews_test[i].Class, pred)
	}
	fmt.Printf("Perc: %f%%\n", (1 - float64(count)/float64(len(p_test.Reviews)))*100)

}
