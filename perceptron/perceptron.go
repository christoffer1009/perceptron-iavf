package perceptron

import (
	"math/rand"
	"time"

	"github.com/christoffer1009/perceptron/review"
)

type Perceptron struct {
	Reviews      []*review.Review
	Words        map[string]float64
	Weigths      map[string]float64
	Bias         float64
	LearningRate float64
}

func NewPerceptron(entries []*review.Review, bias float64, learningRate float64) *Perceptron {
	return &Perceptron{
		Reviews:      entries,
		Weigths:      map[string]float64{},
		Words:        map[string]float64{},
		Bias:         bias,
		LearningRate: learningRate,
	}
}

func (p *Perceptron) Init() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	positive := review.GetPositiveWordsFromTxt()
	negative := review.GetNegativeWordsFromTxt()

	//inicializa com aleatorios
	for _, r := range p.Reviews {
		for _, token := range r.Tokens {
			if _, ok := p.Weigths[token]; !ok {
				p.Weigths[token] = random.Float64()
			}
			if _, ok := p.Words[token]; !ok {
				isPositive := review.Includes(positive, token)
				isNegative := review.Includes(negative, token)
				if isPositive {
					p.Words[token] = 1
				}
				if isNegative {
					p.Words[token] = -1
				}
			}
		}
	}
}

func calcWeights(tokens []string, weights map[string]float64, words map[string]float64) float64 {
	var sum float64 = 0
	for _, t := range tokens {
		sum = sum + weights[t]*words[t]
	}
	return sum
}

func calcNewWeights(tokens []string, weights map[string]float64, words map[string]float64, learningRate float64, e float64) map[string]float64 {

	for _, t := range tokens {
		weights[t] = weights[t] + (learningRate * e * words[t])
	}

	return weights

}

func (p *Perceptron) Train() {

	var step float64 = 0
	for _, r := range p.Reviews {
		sum := calcWeights(r.Tokens, p.Weigths, p.Words)

		if sum >= 0 {
			step = 1
		} else {
			step = 0
		}

		e := r.Class - step

		if e != 0 {
			p.Weigths = calcNewWeights(r.Tokens, p.Weigths, p.Words, p.LearningRate, e)
		}

	}
}

func (p *Perceptron) Predict(r *review.Review) float64 {
	var step float64 = 0
	sum := calcWeights(r.Tokens, p.Weigths, p.Words)

	if sum >= 0 {
		step = 1
	} else {
		step = 0
	}

	// e := r.Class - step

	return step
}
