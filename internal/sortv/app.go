package sortv

import (
	"os"
	"path"

	ray "github.com/gen2brain/raylib-go/raylib"
)

func findResourceDirectory() {
	binaryPath, error := os.Executable()
	if error == nil {
		binaryPath = path.Dir(binaryPath)
		if _, error = os.Stat(binaryPath + "/resources"); error == nil {
			os.Chdir(binaryPath)
			return
		}
	}
	cmdPath, error := os.Getwd()
	if error == nil {
		if _, error = os.Stat(cmdPath + "/resources"); error == nil {
			os.Chdir(cmdPath)
			return
		}
	}
	panic("unable to find resources directory !")
}

func Init() {
	findResourceDirectory()
	ray.SetTraceLog(ray.LogError)
	ray.InitWindow(1280, 720, "Sort Visualizer")
	ray.SetExitKey(0)
	ray.SetWindowState(ray.FlagWindowResizable)
	ray.SetWindowPosition(100, 100)
	ray.SetTargetFPS(60)
	OnWindowResized(ray.GetScreenWidth(), ray.GetScreenHeight())
}

func Run() {
	for !ray.WindowShouldClose() {
		if ray.IsWindowResized() {
			OnWindowResized(ray.GetScreenWidth(), ray.GetScreenHeight())
		}
		Update()
		ray.BeginDrawing()
		Draw()
		ray.EndDrawing()
	}
	close()
}

func close() {
	UnloadUi()
	ray.CloseWindow()
}
