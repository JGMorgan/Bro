package neural_net

import (
	"math"
)

type NeuralNet struct {
	NumInputs       int
	HiddenLayers    int
	NeuronsPerLayer int //placeholder this will layer be an array of ints
	W1              mat64.Dense
	W2              mat64.Dense //currently only 2 weight layers this will be changed if time allows
}

func CreateNeuralNet(numInputs int, numHiddenLayers int, numNeuronsPerLayer int) *NeuralNet {
	return &NeuralNet{numInputs, numHiddenLayers, numNeuronsPerLayer, mat64.NewDense(numInputs, numNeuronsPerLayer, nil), mat64.NewDense(numNeuronsPerLayer, 1, nil)}
}

func Predict(X []float64) float64 {
	X2 := [][]float64{}
	X2 = append(X2, X)
	inputMatrix := mat64.NewDense(1, len(X), X2)

}

func Sigmoid(x float64) float64 {
	return (1 / (1 + math.Pow(math.E, -x)))
}

func SigmoidPrime(x float64) float64 {
	return (math.Pow(math.E, -x) / math.Pow((1+math.Pow(math.E, -x)), 2))
}
