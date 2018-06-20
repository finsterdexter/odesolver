package odesolver

import "testing"
import "fmt"
import . "math"
import "github.com/sbinet/go-gnuplot"

func TestOdeSolve(test *testing.T) {
	var w0, a, wt, g float64 = 1, 0.1, 1, 0.1
	functions := func(t float64, y []float64) []float64 {
		output := []float64{y[1], Pow(-w0, 2)*Sin(y[0]) + a*Sin(wt*t) - g*y[1]}
		return output
	}
	count := int(Pow(2, 10))
	ti := generateRange(0, 2*Pi/w0*20, count)

	t, y := OdeSolve(functions, ti, []float64{0, 0})

	y = transposeArray(y)

	fmt.Println(t)
	fmt.Println(y[0])

	fname := ""
	persist := false
	debug := true
	p, err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.CheckedCmd("set terminal png")
	p.CheckedCmd("set output 'simple.1.png'")
	p.PlotXY(t, y[0], "a graph")
	p.CheckedCmd("quit")
	return
}

func generateRange(start, final float64, count int) []float64 {
	h := (final - start) / float64(count)
	output := make([]float64, count+1)
	for i := range output {
		if i == 0 {
			output[0] = start
			continue
		}
		output[i] = output[i-1] + h
	}
	return output
}

func transposeArray(input [][]float64) [][]float64 {

	output := make([][]float64, len(input[0]))

	// initialize slice
	for i := range output {
		output[i] = make([]float64, len(input))
	}

	for i, v := range input {
		for j := 0; j < len(input[0]); j++ {
			output[j][i] = v[j]
		}
	}

	return output
}
