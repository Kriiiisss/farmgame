package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadPlayerCam() rl.Camera2D {
	var camera rl.Camera2D

	camera.Target = rl.Vector2{X: float32(player.WorldPosition.X), Y: float32(player.WorldPosition.Y)}
	camera.Offset = rl.Vector2{X: float32(rl.GetRenderWidth() / 2), Y: float32(rl.GetRenderHeight() / 2)}
	camera.Rotation = 0.0
	camera.Zoom = 2.0

	return camera
}

func LoadFreeCam() rl.Camera2D {
	var camera rl.Camera2D

	camera.Target = rl.Vector2{X: float32(gameMap.Width) * TILE_SIZE / 2, Y: float32(gameMap.Height) * TILE_SIZE / 2}
	camera.Offset = rl.Vector2{X: float32(rl.GetRenderWidth() / 2), Y: float32(rl.GetRenderHeight() / 2)}
	camera.Rotation = 0.0
	camera.Zoom = 0.5

	return camera
}

func HandleCamera(camera *rl.Camera2D) {
	if rl.IsWindowResized() {
		camera.Offset = rl.Vector2{X: float32(rl.GetRenderWidth() / 2), Y: float32(rl.GetRenderHeight() / 2)}
	}
	targetPos := rl.Vector2{X: float32(math.Round(float64(player.WorldPosition.X))), Y: float32(math.Round(float64(player.WorldPosition.Y - 0.75*player.Height*TILE_SIZE)))}
	camera.Target = targetPos
	if rl.GetMouseWheelMove() != 0 {
		camera.Offset = rl.Vector2{X: float32(rl.GetRenderWidth() / 2), Y: float32(rl.GetRenderHeight() / 2)}
		camera.Zoom = float32(math.Exp(float64(math.Log(float64(camera.Zoom)) + float64(rl.GetMouseWheelMove()*0.1))))
		camera.Zoom = Clamp(camera.Zoom, 1.5, 4.0)
	}
}
