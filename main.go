package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type collisionOBJ struct {
	positionX float32
	positionY float32
	radius    float32
	color     rl.Color
	typeOBJ   int16
}

func getCollision(collider1 collisionOBJ, collider2 collisionOBJ) bool {
	xDiff := float32(math.Abs(float64(collider1.positionX) - float64(collider2.positionX)))
	yDiff := float32(math.Abs(float64(collider1.positionY) - float64(collider2.positionY)))
	if xDiff < (collider1.radius+collider2.radius) && yDiff < (collider1.radius+collider2.radius) {
		return true
	}
	return false
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(800)

	var playerSize float32 = 10.0

	var sideOffset int32 = 120
	var movementOffset int32 = 8

	playerOBJ := collisionOBJ{positionX: 200, positionY: 200, radius: playerSize, color: rl.Blue, typeOBJ: 0}
	getCollision(playerOBJ, playerOBJ)

	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - basic shapes drawing")

	collObjects := []*collisionOBJ{&playerOBJ}
	collObjects = append(collObjects, &collisionOBJ{positionX: 300, positionY: 300, radius: 10.0, color: rl.Red, typeOBJ: 2})
	//Loop objects --> collision

	var progrss float32 = 0.0

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		//MOVEMENT
		if rl.IsKeyDown(rl.KeyD) {
			if playerOBJ.positionX+playerSize+float32(movementOffset) < float32(screenWidth)-float32(sideOffset) {
				playerOBJ.positionX += float32(movementOffset)
			}
		}
		if rl.IsKeyDown(rl.KeyA) {
			if (playerOBJ.positionX-playerSize)-float32(movementOffset) > float32(sideOffset) {
				playerOBJ.positionX -= float32(movementOffset)
			}
		}
		if rl.IsKeyDown(rl.KeyW) {
			if (playerOBJ.positionY-playerSize)-float32(movementOffset) > float32(sideOffset) {
				playerOBJ.positionY -= float32(movementOffset)
			}
		}
		if rl.IsKeyDown(rl.KeyS) {
			if playerOBJ.positionY+playerSize+float32(movementOffset) < float32(screenHeight)-float32(sideOffset) {
				playerOBJ.positionY += float32(movementOffset)
			}
		}

		rl.ClearBackground(rl.Black)

		//SIDEOFFSET
		rl.DrawLine(sideOffset, screenHeight-sideOffset, screenWidth-sideOffset, screenHeight-sideOffset, rl.White)
		rl.DrawLine(sideOffset, sideOffset, screenWidth-sideOffset, sideOffset, rl.White)
		rl.DrawLine(sideOffset, sideOffset, sideOffset, screenHeight-sideOffset, rl.White)
		rl.DrawLine(screenWidth-sideOffset, sideOffset, screenWidth-sideOffset, screenHeight-sideOffset, rl.White)

		//draw entities
		for i := 0; i < len(collObjects); i++ {
			positionUpdate(collObjects[i], progrss)
			rl.DrawCircle(int32(collObjects[i].positionX), int32(collObjects[i].positionY), collObjects[i].radius*1.5, collObjects[i].color)
		}

		//collision check
		for i := 0; i < len(collObjects); i++ {
			for j := 0; j < len(collObjects); j++ {
				if getCollision(*collObjects[i], *collObjects[j]) && i != j {
					panic("ende")
				}
			}
		}

		progrss += 0.5
		if progrss > 100.0 {
			progrss = 0.0
		}

		rl.EndDrawing()
	}
}

func positionUpdate(inputOBJ *collisionOBJ, inputProgress float32) {
	switch inputOBJ.typeOBJ {
	case 0:
		return
	case 1:
		inputOBJ.positionX = 400.0 + float32(math.Cos(float64(inputProgress))*70.0)
		inputOBJ.positionY = 400.0 + float32(math.Sin(float64(inputProgress))*70.0)
		return
	case 2:
		inputOBJ.positionX = 800 - (inputProgress*2/100)*(800)
		inputOBJ.positionY = (inputProgress * 2 / 100) * (800)
		return
	case 3:
		inputOBJ.positionX = 800 - (inputProgress*2/100)*(800)
		inputOBJ.positionY = 800 - (inputProgress*2/100)*(800)
		return
	case 4:
		inputOBJ.positionX = (inputProgress * 2 / 100) * (800)
		inputOBJ.positionY = 800 - (inputProgress*2/100)*(800)
		return
	default:
		inputOBJ.positionX = (inputProgress * 1 / 100) * (800)
		inputOBJ.positionY = (inputProgress * 1 / 100) * (800)
		return
	}
}
