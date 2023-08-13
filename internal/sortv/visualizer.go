package sortv

import (
	"fmt"
	"github.com/rzaf/sortVisualizer/pkg/algos"
	"github.com/rzaf/sortVisualizer/pkg/core"
	"math/rand"
	"time"

	ray "github.com/gen2brain/raylib-go/raylib"
)

const (
	minSliceLength = 10
	maxSliceLength = 100
	minDelay       = 5
	maxDelay       = 800
)

var (
	sortingSlice = make(Sortable, 20, maxSliceLength)
	visibleSlice = make(Sortable, 20, maxSliceLength)

	sortAlgos = []*sortAlgo{
		NewSortAlgo(algos.BubbleSort, "Bubble Sort"),
		NewSortAlgo(algos.InsertionSort, "Insertion Sort"),
		NewSortAlgo(algos.SelectionSort, "Selection Sort"),
		NewSortAlgo(algos.QuickSort, "Quick Sort(sort.Sort)"),
		NewSortAlgo(algos.HeapSort, "Heap Sort"),
	}
	currentAlgoIndex = 0

	barWidth   = 50
	barDist    = 6
	barXOffset = 5

	stepIndex    = -1
	stepI, stepJ = -1, -1
	startTime    int64
	delay        int64 = 50

	isSwapping bool
	nightMode  bool = true
)

func init() {
	rand.Seed(time.Now().Unix())
	randomize()
	sizeSlider = core.NewSlider(0, 0, 180, 15, minSliceLength, maxSliceLength, float32(sortingSlice.Len()))
	delaySlider = core.NewSlider(0, 0, 230, 15, minDelay, maxDelay, float32(delay))
}

func OnWindowResized(width int, height int) {
	calculateBarWidth()
	sizeSlider.SetPos(int32(ray.GetScreenWidth())-250, 42)
	delaySlider.SetPos(int32(ray.GetScreenWidth())-250, 92)
}

func randomize() {
	for i := 0; i < len(sortingSlice); i++ {
		sortingSlice[i] = 50 + rand.Intn(400)
	}
	copy(visibleSlice, sortingSlice)
}

func startSorting() {
	startButton.SetTexture(stopTexture)
	steps = make([]*Step, 0, 10)
	stepIndex = 0
	stepI, stepJ = -1, -1
	copy(sortingSlice, visibleSlice)
	sortAlgos[currentAlgoIndex].Sort(sortingSlice)
	fmt.Println("starting ")
	startTime = time.Now().UnixMilli()
}

func stopSorting() {
	startButton.SetTexture(playTexture)
	fmt.Println("stoping")
	stepIndex = -1
	startTime = 0
	stepI, stepJ = -1, -1
}

func Update() {
	UpdateUi()
	if stepIndex >= 0 {
		if t := time.Now().UnixMilli(); startTime+delay < t {
			stepI, stepJ = steps[stepIndex].i, steps[stepIndex].j
			if steps[stepIndex].StepType == StepSwap {
				i := steps[stepIndex].i
				j := steps[stepIndex].j
				visibleSlice[i], visibleSlice[j] = visibleSlice[j], visibleSlice[i]
				isSwapping = true
			} else {
				isSwapping = false
			}
			stepIndex++
			startTime = time.Now().UnixMilli()
			if stepIndex == len(steps) {
				stopSorting()
			}
		}
	}
}

func Draw() {
	if nightMode {
		ray.ClearBackground(ray.Black)
	} else {
		ray.ClearBackground(ray.RayWhite)
	}
	DrawUi()
	drawBars()
}
