package main

import (
	"time"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/flan6/data-algo/algorithms/sorting"
)

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	length := 25
	data := make([]float64, length)
	for i := range data {
		data[i] = float64(length - i)
	}

	draw := func() {
		spline1 := widgets.NewSparkline()
		spline1.Data = data
		spline1.LineColor = termui.ColorCyan
		spline1.Title = "SelectionSort"

		group := widgets.NewSparklineGroup(spline1)
		group.SetRect(0, 0, 30, 15)
		group.Title = "Sorting Visualization"

		termui.Render(group)
	}

	draw()

	sorting.SelectionSortVisual(data, func() {
		draw()
		time.Sleep(500 * time.Millisecond)
	})

	termuiEvents := termui.PollEvents()
	for {
		e := <-termuiEvents
		if e.Type == termui.KeyboardEvent {
			break
		}
	}
}
