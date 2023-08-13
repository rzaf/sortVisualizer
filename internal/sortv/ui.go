package sortv

import (
	"fmt"
	"github.com/rzaf/sortVisualizer/pkg/core"
	"math/rand"

	ray "github.com/gen2brain/raylib-go/raylib"
)

var (
	circleTexture   *core.Texture
	playTexture     *core.Texture
	stopTexture     *core.Texture
	shuffleTexture  *core.Texture
	darkmodeTexture *core.Texture
	daymodeTexture  *core.Texture
	sortTexture     *core.Texture

	startButton      *core.Button
	shuffleButton    *core.Button
	darkmodeButton   *core.Button
	changeSortButton *core.Button

	sizeSlider  *core.Slider
	delaySlider *core.Slider
	isLoaded    = false
)

func LoadUi() {
	circleTexture = core.NewTexture("resources/circle.png", ray.NewRectangle(0, 0, 120, 120))
	playTexture = core.NewTexture("resources/play.png", ray.NewRectangle(0, 0, 100, 100)).SetDst(ray.NewRectangle(54, 20, 50, 50))
	stopTexture = core.NewTexture("resources/stop.png", ray.NewRectangle(0, 0, 100, 100)).SetDst(ray.NewRectangle(50, 20, 50, 50))
	shuffleTexture = core.NewTexture("resources/shuffle.png", ray.NewRectangle(0, 0, 100, 100)).SetDst(ray.NewRectangle(140, 20, 50, 50))
	darkmodeTexture = core.NewTexture("resources/darkmode.png", ray.NewRectangle(0, 0, 100, 100)).SetDst(ray.NewRectangle(55, 88, 50, 50))
	daymodeTexture = core.NewTexture("resources/daymode.png", ray.NewRectangle(0, 0, 100, 100)).SetDst(ray.NewRectangle(55, 88, 50, 50))
	sortTexture = core.NewTexture("resources/sort.png", ray.NewRectangle(0, 0, 100, 100)).SetDst(ray.NewRectangle(55, 158, 50, 50))
	circleTexture.OnDraw = func() {
		if nightMode {
			circleTexture.DrawAt(startButton.Boundary)
			circleTexture.DrawAt(darkmodeButton.Boundary)
			if stepIndex == -1 {
				circleTexture.DrawAt(shuffleButton.Boundary)
				circleTexture.DrawAt(changeSortButton.Boundary)
			}
		}
	}

	startButton = core.NewButton(playTexture, ray.NewRectangle(45, 15, 60, 60))
	shuffleButton = core.NewButton(shuffleTexture, ray.NewRectangle(135, 15, 60, 60))
	darkmodeButton = core.NewButton(daymodeTexture, ray.NewRectangle(50, 85, 60, 60))
	changeSortButton = core.NewButton(sortTexture, ray.NewRectangle(50, 152, 60, 60))
	startButton.OnClick = func() {
		if startTime == 0 {
			startSorting()
		} else {
			stopSorting()
		}
	}
	shuffleButton.OnClick = func() {
		if stepIndex == -1 {
			randomize()
		}
	}
	darkmodeButton.OnClick = func() {
		nightMode = !nightMode
		if nightMode {
			sizeSlider.SetColors(ray.Gray, ray.NewColor(230, 230, 230, 250), ray.White)
			delaySlider.SetColors(ray.Gray, ray.NewColor(230, 230, 230, 250), ray.White)
			darkmodeButton.SetTexture(daymodeTexture)
		} else {
			darkmodeButton.SetTexture(darkmodeTexture)
			sizeSlider.SetColors(ray.Gray, ray.Black, ray.Black)
			delaySlider.SetColors(ray.Gray, ray.Black, ray.Black)
		}
	}
	changeSortButton.OnClick = func() {
		if stepIndex == -1 {
			currentAlgoIndex = (currentAlgoIndex + 1) % len(sortAlgos)
		}
	}
	isLoaded = true
}

func UnloadUi() {
	playTexture.Unload()
	stopTexture.Unload()
	shuffleTexture.Unload()
}

func UpdateUi() {
	if !isLoaded {
		LoadUi()
	}
	if l := sizeSlider.GetValueInt(); l != sortingSlice.Len() {
		stopSorting()
		if l <= sortingSlice.Len() {
			sortingSlice = sortingSlice[:l]
			visibleSlice = visibleSlice[:l]
		} else {
			l0 := sortingSlice.Len()
			sortingSlice = sortingSlice[:l]
			visibleSlice = visibleSlice[:l]
			for i := l0; i < l; i++ {
				sortingSlice[i] = 50 + rand.Intn(400)
				visibleSlice[i] = 50 + rand.Intn(400)
			}
		}
		copy(visibleSlice, sortingSlice)
		calculateBarWidth()
	}
	if d := int64(delaySlider.GetValueInt()); delay != d {
		stopSorting()
		delay = d
	}

	startButton.Update()
	darkmodeButton.Update()
	shuffleButton.Update()
	sizeSlider.Update()
	delaySlider.Update()
	changeSortButton.Update()
}

func DrawUi() {
	sizeSlider.Draw()
	delaySlider.Draw()
	if nightMode {
		ray.DrawText(fmt.Sprintf("len: %d", sizeSlider.GetValueInt()), int32(ray.GetScreenWidth())-230, 15, 20, ray.White)
		ray.DrawText(fmt.Sprintf("delay: %d ms", delaySlider.GetValueInt()), int32(ray.GetScreenWidth())-230, 65, 20, ray.White)
	} else {
		ray.DrawText(fmt.Sprintf("len: %d", sizeSlider.GetValueInt()), int32(ray.GetScreenWidth())-230, 15, 20, ray.Black)
		ray.DrawText(fmt.Sprintf("delay: %d ms", delaySlider.GetValueInt()), int32(ray.GetScreenWidth())-230, 65, 20, ray.Black)
	}

	circleTexture.Draw()
	if stepIndex == -1 {
		if nightMode {
			ray.DrawText(fmt.Sprint(sortAlgos[currentAlgoIndex].name), 125, 165, 28, ray.White)
		} else {
			ray.DrawText(fmt.Sprint(sortAlgos[currentAlgoIndex].name), 125, 165, 28, ray.Black)
		}
		shuffleButton.Draw()
		changeSortButton.Draw()
	} else {
		ray.DrawText(fmt.Sprint(sortAlgos[currentAlgoIndex].name), int32(ray.GetScreenWidth()/2-115), 100, 35, ray.Blue)
	}
	startButton.Draw()
	darkmodeButton.Draw()
}
