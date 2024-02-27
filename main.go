package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	r                   = initRocket()
	rocketAlpha         = 255
	lasers              []*rl.Vector2
	enemies             []*enemy
	lastSpawn           = time.Now().Unix()
	lastHit             = time.Now().Unix()
	rockTexture2D       rl.Texture2D
	ufoTexture2D        rl.Texture2D
	backgroundTexture2D rl.Texture2D
	rocketTexture2D     rl.Texture2D
	rockSmall           rl.Texture2D
	ufoSmall            rl.Texture2D
	heart               rl.Texture2D
	explosionTexture2D  rl.Texture2D

	gameOver bool
)

func drawHud() {
	rl.DrawRectangle(0, 0, 80, 122, rl.Black)
	rl.DrawTexture(heart, 0, 0, rl.White)
	rl.DrawText(fmt.Sprintf("%d", r.hearts), 42, 0, 32, rl.White)
	rl.DrawTexture(rockSmall, 0, 42, rl.White)
	rl.DrawText(fmt.Sprintf("%d", r.destroyedObj[rock]), 42, 42, 32, rl.White)
	rl.DrawTexture(ufoSmall, 0, 84, rl.White)
	rl.DrawText(fmt.Sprintf("%d", r.destroyedObj[ufo]), 42, 82, 32, rl.White)
}

func drawScreen(text ...string) {
	rl.DrawRectangle(0, 0, windowWidth, windowHeight, color.RGBA{R: 0, G: 0, B: 0, A: 187})
	offset := int32(0)
	for !rl.WindowShouldClose() && !rl.IsKeyPressed(rl.KeySpace) {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		drawHud()
		for _, v := range text {
			rl.DrawText(v, windowWidth/2-int32((len(v)/2)*35), windowHeight/2+offset, 64, rl.White)
			offset += 64
		}
		offset = 0
		rl.EndDrawing()

	}

}

func handleInput() {
	if rl.IsKeyDown(rl.KeyA) {
		if r.x > 10 {
			r.x -= rocketMoveSpeed
		}
	} else if rl.IsKeyDown(rl.KeyD) {
		if r.x < windowWidth-10-rocketRes {
			r.x += rocketMoveSpeed
		}
	}
	if rl.IsKeyPressed(rl.KeyS) {
		vec := rl.Vector2{
			X: float32(r.x + (rocketRes / 2) - (laserWidth / 2)),
			Y: float32(rocketYConst - rocketRes),
		}
		lasers = append(lasers, &vec)
	}

}

func spawnEnemies() {
	if time.Now().Unix()-lastSpawn >= 1 {
		var newWave []*enemy
		for i := 0; i < rand.Intn(4)+3; i++ {
			hp := rockHp
			t := rock
			if rand.Int31n(20) == 1 {
				hp = ufoHp
				t = ufo
			}
		retry:
			x := rand.Int31n(windowWidth-enemyRes-10) + 10
			for _, v := range newWave {
				r1 := rl.Rectangle{
					X:      float32(v.x),
					Y:      float32(v.y),
					Width:  float32(enemyRes),
					Height: float32(enemyRes),
				}
				r2 := rl.Rectangle{
					X:      float32(x),
					Y:      0,
					Width:  float32(enemyRes),
					Height: float32(enemyRes),
				}
				if rl.CheckCollisionRecs(r1, r2) {
					goto retry
				}
			}
			newWave = append(newWave, &enemy{
				t:             t,
				y:             rand.Int31n(enemyYConst),
				x:             x,
				explosionTime: explosionTime,
				hp:            hp,
			})
		}
		enemies = append(enemies, newWave...)
		lastSpawn = time.Now().Unix()
	}
}

func updatePositions() {
	for i, v := range lasers {
		v.Y -= float32(laserSpeed)
		if v.Y < 0 {
			lasers = append(lasers[:i], lasers[i+1:]...)
		}
	}
	for i, v := range enemies {
		v.y += enemyMoveSpeed
		if v.y > windowHeight {
			enemies = append(enemies[:i], enemies[i+1:]...)
		}
	}

}

func checkForCollision() {
	if r.hit {
		if rocketAlpha == 0 {
			rocketAlpha = 255
		} else {
			rocketAlpha -= 10
		}
	}
	if r.hit && (time.Now().Unix()-lastHit) >= hitDelay {
		r.hit = false
		rocketAlpha = 255
	}
outer:
	for _, e := range enemies {
		if e.invisible {
			continue
		}
		er := rl.Rectangle{
			X:      float32(e.x),
			Y:      float32(e.y),
			Width:  float32(enemyRes),
			Height: float32(enemyRes),
		}
		for i, l := range lasers {
			lr := rl.Rectangle{
				X:      l.X,
				Y:      l.Y,
				Width:  float32(laserWidth),
				Height: float32(laserHeight),
			}
			if rl.CheckCollisionRecs(er, lr) {
				e.hp -= 1
				lasers = append(lasers[:i], lasers[i+1:]...)
				if e.hp == 0 {
					e.invisible = true
					r.destroyedObj[e.t] += 1
					continue outer
				}

			}
		}
		pr := rl.Rectangle{
			X:      float32(r.x),
			Y:      float32(rocketYConst),
			Width:  float32(rocketRes),
			Height: float32(rocketRes),
		}
		if rl.CheckCollisionRecs(er, pr) && !r.hit {
			r.hit = true
			lastHit = time.Now().Unix()
			r.hearts -= 1

			if r.hearts == 0 {
				gameOver = true
			}
		}

	}
}

func draw() {
	for _, v := range lasers {
		rl.DrawRectangle(int32(v.X), int32(v.Y), laserWidth, laserHeight, rl.Red)

	}
	for _, v := range enemies {
		if v.invisible {
			if v.explosionTime >= 0 {
				rl.DrawTexture(explosionTexture2D, v.x, v.y, rl.White)
				v.explosionTime -= 1
			}
			continue
		}
		if v.t == rock {
			rl.DrawTexture(rockTexture2D, v.x, v.y, rl.White)
		} else {
			rl.DrawTexture(ufoTexture2D, v.x, v.y, rl.White)
		}
	}
	rl.DrawTexture(rocketTexture2D, r.x, rocketYConst, color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: uint8(rocketAlpha),
	})
}

func main() {
	rl.InitWindow(windowWidth, windowHeight, "RocketBoom")
	defer rl.CloseWindow()

	rocketImg := rl.LoadImage(rocketTexture)
	rocketTexture2D = rl.LoadTextureFromImage(rocketImg)
	rockImg := rl.LoadImage(rockTexture)
	rockTexture2D = rl.LoadTextureFromImage(rockImg)
	ufoImg := rl.LoadImage(ufoTexture)
	ufoTexture2D = rl.LoadTextureFromImage(ufoImg)
	backImg := rl.LoadImage(backgroundTexture)
	backgroundTexture2D = rl.LoadTextureFromImage(backImg)

	rl.ImageResize(ufoImg, 32, 32)
	ufoSmall = rl.LoadTextureFromImage(ufoImg)

	rl.ImageResize(rockImg, 32, 32)
	rockSmall = rl.LoadTextureFromImage(rockImg)

	heartImg := rl.LoadImage(heartTexture)
	heart = rl.LoadTextureFromImage(heartImg)

	explosionImg := rl.LoadImage(explosionTexture)
	explosionTexture2D = rl.LoadTextureFromImage(explosionImg)

	rl.UnloadImage(explosionImg)

	defer rl.UnloadTexture(explosionTexture2D)

	rl.UnloadImage(heartImg)
	defer rl.UnloadTexture(heart)

	rl.UnloadImage(backImg)
	defer rl.UnloadTexture(backgroundTexture2D)

	rl.UnloadImage(ufoImg)
	defer rl.UnloadTexture(ufoTexture2D)

	rl.UnloadImage(rockImg)
	defer rl.UnloadTexture(rockTexture2D)

	rl.UnloadImage(rocketImg)
	defer rl.UnloadTexture(rocketTexture2D)

	rl.SetTargetFPS(60)

	defer func() {
		drawScreen("Game Over", "Press space to exit")
	}()
	drawScreen("RocketBoom", "Press Space to start")
	for !rl.WindowShouldClose() && !gameOver {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTexture(backgroundTexture2D, 0, 0, rl.White)
		spawnEnemies()
		checkForCollision()
		handleInput()
		updatePositions()
		draw()
		drawHud()
		rl.EndDrawing()
	}

}
