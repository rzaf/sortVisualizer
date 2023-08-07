package sortv

import (
	"time"

	ray "github.com/gen2brain/raylib-go/raylib"
)

func calculateBarWidth() {
	l := len(sortingSlice)
	if l > 25 {
		barDist = 2
	} else {
		barDist = 6
	}
	x := ray.GetScreenWidth() - l*barDist
	barWidth = x / l
	barXOffset = ray.GetScreenWidth() - barWidth*l - l*barDist
	barXOffset /= 2
}

func drawBars() {
	for i := 0; i < visibleSlice.Len(); i++ {
		if i == stepI || i == stepJ {
			continue
		}
		drawBar(i, visibleSlice[i])
	}
	if stepI != -1 {
		drawBar(stepI, visibleSlice[stepI])
	}
	if stepJ != -1 {
		drawBar(stepJ, visibleSlice[stepJ])
	}
}

func drawBar(i int, value int) {
	color := ray.Blue
	x := barX(i)
	if stepIndex > 0 && stepIndex < len(steps) {
		if i == stepI || i == stepJ {
			color = ray.Red
		}
	}
	if isSwapping {
		if i == stepI {
			percentage := float32(time.Now().UnixMilli()-startTime) / float32(delay)
			if percentage > 0.5 {
				percentage = 0.5
			}
			x2 := barX(stepJ)
			x = lerp(float32(x2), float32(x), percentage)
		}
		if i == stepJ {
			percentage := float32(time.Now().UnixMilli()-startTime) / float32(delay)
			if percentage > 0.5 {
				percentage = 0.5
			}
			x2 := barX(stepI)
			x = lerp(float32(x2), float32(x), percentage)
		}
	}
	ray.DrawRectangle(
		x,
		int32(ray.GetScreenHeight()-value),
		int32(barWidth),
		int32(value), //int32(visibleSlice[i]*10),
		color,
	)
}

func barX(index int) int32 {
	return int32(barXOffset + index*barWidth + index*barDist)
}

func lerp(v0 float32, v1 float32, t float32) int32 {
	// return int32(v0 + t*(v1-v0))
	return int32((1-t*2)*v0 + t*2*v1)
}
