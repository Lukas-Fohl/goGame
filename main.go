package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type collisionOBJ struct {
	positionX float64
	positionY float64
	width     float64
	height    float64
}

func getCollision(collider1 collisionOBJ, collider2 collisionOBJ) bool {
	xDiff := math.Abs(collider1.positionX - collider2.positionX)
	yDiff := math.Abs(collider1.positionY - collider2.positionY)
	if xDiff < (collider1.width+collider2.width) || yDiff < (collider1.height+collider2.height) {
		return true
	}
	return false
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
