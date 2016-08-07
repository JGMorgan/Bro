package neural_network

import (
	"math"
	"math/rand"
	"time"
)

type NeuralNet struct {
	NumInputs       int
	HiddenLayers    int
	NeuronsPerLayer int //placeholder this will layer be an array of ints
	W1              [][]float64
	W2              []float64 //currently only 2 weight layers this will be changed if time allows
}

func CreateNeuralNet(numInputs int, numHiddenLayers int, numNeuronsPerLayer int) *NeuralNet {
	W1 := [][]float64{}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < numInputs; i++ {
		row := make([]float64, numNeuronsPerLayer)
		for j := 0; j < numNeuronsPerLayer; j++ {
			row[j] = r1.NormFloat64()
		}
		W1 = append(W1, row)
	}
	W2 := make([]float64, numNeuronsPerLayer)
	for i := 0; i < numNeuronsPerLayer; i++ {
		W2[i] = r1.NormFloat64()
	}
	return &NeuralNet{numInputs, numHiddenLayers, numNeuronsPerLayer, W1, W2}
}

func (n *NeuralNet) Predict(X []float64) float64 {
	return Sigmoid(n.zFinal(X))
}

func (n *NeuralNet) zFinal(X []float64) float64 {
	z := n.multW1(X)
	for index, element := range z {
		z[index] = Sigmoid(element)
	}
	zFinal := n.multW2(z)
	return zFinal
}

func (n *NeuralNet) Cost(X []float64, y float64) float64 {
	return 0.5 * math.Pow(y-n.Predict(X), 2)
}

func (n *NeuralNet) delta2(X []float64, y float64) float64 {
	return -(y - n.Predict(X)) * SigmoidPrime(n.zFinal(X))
}

func (n *NeuralNet) delta1(X []float64, y float64) []float64 {
	out := make([]float64, len(n.W2))
	for index, element := range n.W2 {
		out[index] = element * n.delta2(X, y) * n.W2[index]
	}
	return out
}

func (n *NeuralNet) DJDW2(X []float64, y float64) []float64 {
	a2 := n.multW1(X)
	out := make([]float64, len(a2))
	for index, element := range a2 {
		out[index] = element * n.delta2(X, y)
	}
	return out
}

func (n *NeuralNet) DJDW1(X []float64, y float64) [][]float64 {
	out := [][]float64{}
	d1 := n.delta1(X, y)
	for _, element1 := range X {
		row := make([]float64, len(d1))
		for j, element2 := range d1 {
			row[j] = element1 * element2
		}
		out = append(out, row)
	}
	return out
}

func (n *NeuralNet) multW1(X []float64) []float64 {
	Z := make([]float64, n.NeuronsPerLayer)
	for i := 0; i < n.NeuronsPerLayer; i++ {
		sum := 0.0
		for index, element := range X {
			sum = sum + (element * n.W1[index][i])
		}
		Z[i] = sum
	}
	return Z
}

func (n *NeuralNet) multW2(Z []float64) float64 {
	sum := 0.0
	for index, element := range Z {
		sum = sum + (element * n.W2[index])
	}
	return sum
}

func Sigmoid(x float64) float64 {
	return (1 / (1 + math.Pow(math.E, -x)))
}

func SigmoidPrime(x float64) float64 {
	return (math.Pow(math.E, -x) / math.Pow((1+math.Pow(math.E, -x)), 2))
}

func (n *NeuralNet) Train(X [][]float64, Y []float64, a float64) {
	for i := 0; i < len(X); i++ {

		//			DW2 := n.DJDW2(X[i], Y[i])
		//			DW1 := n.DJDW1(X[i], Y[i])
		//	accurate1 := true
		//	for accurate1 {
		for h := 0; h < 20; h++ {

			DW2 := n.DJDW2(X[i], Y[i])
			DW1 := n.DJDW1(X[i], Y[i])
			for index, element := range DW2 {
				n.W2[index] -= (element * a)
				//	accurate1 = !(accurate1 && (element < .05) && (element > -.05))
			}
			//	}
			//	accurate2 := true
			//	for accurate2 {
			for j := 0; j < n.NumInputs; j++ {
				for k := 0; k < n.NeuronsPerLayer; k++ {
					n.W1[j][k] -= (DW1[j][k] * a)
					//		accurate2 = !(accurate2 && (DW1[j][k] < .05) && (DW1[j][k] > -.05))
				}
			}
		}
		//	}
	}

}
