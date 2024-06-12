package main

import (
	"math"
	"math/rand/v2"

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
		if collider1.typeOBJ == 0 || collider2.typeOBJ == 0 {
			return true
		}
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
	collObjects = append(collObjects, &collisionOBJ{positionX: 300, positionY: 300, radius: 20.0, color: rl.Red, typeOBJ: 2})
	//Loop objects --> collision

	var progrss float32 = 0.0

	var iterations int32 = 0

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
			positionUpdate(collObjects[i], collObjects[0], progrss, iterations)
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

		if progrss > 100.0 {
			iterations++
			progrss = 0.5
			for i := 0; i < len(collObjects)-1; i++ {
				collObjects = collObjects[:len(collObjects)-1]
			}
			var typeNum int16 = int16(rand.IntN(8)) + 1
			switch typeNum {
			case 7:
				collObjects = append(collObjects, &collisionOBJ{
					positionX: 0.0,
					positionY: 0.0,
					radius:    20.0 * (1.0 + (float32(iterations) / 100.0)),
					typeOBJ:   2,
					color:     rl.Red,
				})
				collObjects = append(collObjects, &collisionOBJ{
					positionX: 0.0,
					positionY: 800.0,
					radius:    20.0 * (1.0 + (float32(iterations) / 100.0)),
					typeOBJ:   3,
					color:     rl.Red,
				})
			default:
				collObjects = append(collObjects, &collisionOBJ{
					positionX: 0.0,
					positionY: 0.0,
					radius:    20.0 * (1.0 + (float32(iterations) / 100.0)),
					typeOBJ:   typeNum,
					color:     rl.Red,
				})
			}
		}
		for i := 0; i < len(collObjects); i++ {
			switch collObjects[len(collObjects)-1].typeOBJ {
			case 0:
				continue
			case 6:
				progrss += 0.2
				break
			default:
				progrss += 0.3
				break
			}
		}

		rl.EndDrawing()
	}
}

func positionUpdate(inputOBJ *collisionOBJ, player *collisionOBJ, inputProgress float32, iterations int32) {
	switch inputOBJ.typeOBJ {
	case 0:
		return
	case 1:
		inputOBJ.positionX = 800 - (inputProgress*2/100)*(800)
		inputOBJ.positionY = (inputProgress * 2 / 100) * (800)
		return
	case 2:
		inputOBJ.positionX = 800 - (inputProgress*2/100)*(800)
		inputOBJ.positionY = 800 - (inputProgress*2/100)*(800)
		return
	case 3:
		inputOBJ.positionX = (inputProgress * 2 / 100) * (800)
		inputOBJ.positionY = 800 - (inputProgress*2/100)*(800)
		return
	case 4:
		inputOBJ.positionX = (inputProgress * 2 / 100) * (800)
		inputOBJ.positionY = 400.0 + float32(math.Sin(float64(inputProgress))*40.0)
		return
	case 5:
		inputOBJ.positionY = (inputProgress * 2 / 100) * (800)
		inputOBJ.positionX = 400.0 + float32(math.Sin(float64(inputProgress))*40.0)
		return
	case 6:
		inputOBJ.color = rl.Green
		/*
			yDiffn**2 + xDiffn**2 = 25
			xDiff / yDiff = relation
			xDiff = relation * yDiff
			yDiffn**2 + (relation * yDiff) = 25
			xDiffn**2 = 25 - ydiffn**2
		*/
		xDiff := player.positionX - inputOBJ.positionX
		yDiff := player.positionY - inputOBJ.positionY
		relation := float32(math.Abs(float64(xDiff / yDiff)))
		if relation > 1 {
			relation = 1
		} else if relation < -1 {
			relation = -1
		}

		if xDiff > 0 {
			inputOBJ.positionX += float32(relation * 5.0 * float32(1+(iterations/100)))
		} else {
			inputOBJ.positionX -= float32(relation * 5.0 * float32(1+(iterations/100)))
		}

		if yDiff > 0 {
			inputOBJ.positionY += float32(5 * (1 + (iterations / 100)))
		} else {
			inputOBJ.positionY -= float32(5 * (1 + (iterations / 100)))
		}
		if inputProgress > 98.0 {
			inputOBJ.positionX = -100
			inputOBJ.positionY = -100
		}
		return
	default:
		inputOBJ.positionX = (inputProgress * 1 / 100) * (800)
		inputOBJ.positionY = (inputProgress * 1 / 100) * (800)
		return
	}
}
