package main

import (
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	date()
}

//time
func date() {
	buffer := make([]float64, 0)
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "date"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			buffer = append(buffer, float64(time.Now().Unix()))

			err = plotutil.AddLinePoints(p,
				"raw", randomPoints(len(buffer), nil, buffer),
			)
			if err != nil {
				panic(err)
			}

			// Save the plot to a PNG file.
			if err := p.Save(4*vg.Inch, 4*vg.Inch, "date.png"); err != nil {
				panic(err)
			}
		}
	}

}

// //fft
// func fft() {
// 	rand.Seed(int64(0))
// 	rawData := make([]float64, 20)
// 	for i, _ := range rawData {
// 		rawData[i] = 10 * rand.Float64()
// 	}

// 	fmt.Println("raw data")
// 	fmt.Println(rawData)

// 	pRaw, err := plot.New()
// 	if err != nil {
// 		panic(err)
// 	}

// 	pRaw.Title.Text = "Raw data"
// 	pRaw.X.Label.Text = "X"
// 	pRaw.Y.Label.Text = "Y"

// 	err = plotutil.AddLinePoints(pRaw,
// 		"raw", randomPoints(len(rawData), nil, rawData),
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Save the plot to a PNG file.
// 	if err := pRaw.Save(4*vg.Inch, 4*vg.Inch, "raw.png"); err != nil {
// 		panic(err)
// 	}

// 	// FFT DATA
// 	fftDataRaw := fft.FFTReal(rawData)
// 	fftData := make([]float64, len(fftDataRaw))
// 	for i, _ := range fftDataRaw {
// 		fftData[i] = cmplx.Phase(fftDataRaw[i])
// 	}

// 	pFFT, err := plot.New()
// 	if err != nil {
// 		panic(err)
// 	}

// 	pFFT.Title.Text = "ft data"
// 	pFFT.X.Label.Text = "X"
// 	pFFT.Y.Label.Text = "Y"

// 	err = plotutil.AddLinePoints(pFFT,
// 		"fft", randomPoints(len(fftData), nil, fftData),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Save the plot to a PNG file.
// 	if err := pFFT.Save(4*vg.Inch, 4*vg.Inch, "fft.png"); err != nil {
// 		panic(err)
// 	}
// }

// randomPoints returns some random x, y points.
func randomPoints(n int, x, y []float64) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {

		pts[i].X = float64(i)
		pts[i].Y = y[i]

	}
	return pts
}
