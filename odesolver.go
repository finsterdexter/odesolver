package odesolver

type ode func(float64, []float64) []float64

// classical Runge-Kutta
func OdeSolve(f ode, t []float64, initialConditions []float64) ([]float64, [][]float64) {
	// t is the independent variable
	y := make([][]float64, len(t)) // dependent variable y, each element could itself be an array
	var h float64 = t[1] - t[0] // step size, i.e. delta t
	var lambda float64 = 2 // a free parameter, lambda = 2 is the classical Runge-Kutta
	
	// len(y) should match len(t)
	// len(y[n]) should match len(initialConditions)

	for i, _ := range t {
		output := make([]float64, len(initialConditions))
		
		if i == 0 {
			output = initialConditions
			y[0] = output
			continue
		}
		
		k1 := make([]float64, len(output))
		k2 := make([]float64, len(output))
		k3 := make([]float64, len(output))
		k4 := make([]float64, len(output))
		
		k1 = f(t[i-1], y[i-1])
		
		// build dependent array for k2
		y2 := make([]float64, len(output))
		for n, v := range y[i-1] {
			y2[n] = v + 1/2 * k1[n] * h
		}
		k2 = f(t[i-1] + 1/2 * h, y2)
		
		// build dependent array for k3
		y3 := make([]float64, len(output))
		for n, v := range y[i-1] {
			y3[n] = v + (1/2 - 1/lambda) * k1[n] * h + 1/lambda * k2[n] * h
		}
		k3 = f(t[i-1] + 1/2 * h, y3)
		
		// build dependent array for k4
		y4 := make([]float64, len(output))
		for n, v := range y[i-1] {
			y4[n] = v + (1 - lambda/2) * k2[n] * h + lambda/2 * k3[n] * h
		}
		k4 = f(t[i-1] + h, y4)
		
		for j, _ := range initialConditions {
			output[j] = y[i-1.0][j] + 1.0/6.0 * h * (k1[j] + (4.0 - lambda) * k2[j] + lambda * k3[j] + k4[j])
		}
		y[i] = output
	
		if i+1 == len(t)  {
			break
		}
		h = t[i+1] - t[i]
	}
	
	return t, y
}