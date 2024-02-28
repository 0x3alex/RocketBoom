package main

import (
	"fmt"
	"image/color"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	rocketPlayer         = initRocket()
	rocketAlpha          = 255
	lastHit              = time.Now().Unix()
	rockTexture2D        rl.Texture2D
	ufoTexture2D         rl.Texture2D
	backgroundTexture2D  rl.Texture2D
	rocketTexture2D      rl.Texture2D
	rockSmall            rl.Texture2D
	ufoSmall             rl.Texture2D
	missleSmall          rl.Texture2D
	heartTexture2D       rl.Texture2D
	explosionTexture2D   rl.Texture2D
	laserTexture2D       rl.Texture2D
	clusterBombTexture2D rl.Texture2D
	missleTexture2D      rl.Texture2D
	minigunTexture2D     rl.Texture2D
	explosionSFX         rl.Sound
	laserSFX             rl.Sound
	hitSFX               rl.Sound
	gameOver             bool
)

func drawHud() {
	rl.DrawRectangle(0, 0, 80, 122, rl.Black)
	rl.DrawTexture(heartTexture2D, 0, 0, rl.White)
	rl.DrawText(fmt.Sprintf("%d", rocketPlayer.hearts), 42, 0, 32, rl.White)
	rl.DrawTexture(rockSmall, 0, 42, rl.White)
	rl.DrawText(fmt.Sprintf("%d", rocketPlayer.destroyedObj[rock]), 42, 42, 32, rl.White)
	rl.DrawTexture(ufoSmall, 0, 84, rl.White)
	rl.DrawText(fmt.Sprintf("%d", rocketPlayer.destroyedObj[ufo]), 42, 82, 32, rl.White)
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
		if rocketPlayer.x > 10 {
			rocketPlayer.x -= rocketMoveSpeed
		}
	} else if rl.IsKeyDown(rl.KeyD) {
		if rocketPlayer.x < windowWidth-10-rocketRes {
			rocketPlayer.x += rocketMoveSpeed
		}
	}
	if rl.IsKeyPressed(rl.KeyS) || rocketPlayer.powerUps["minigun"].active {
		if !rocketPlayer.powerUps["minigun"].active {
			rl.PlaySound(laserSFX)
		}
		t := laser
		if rocketPlayer.powerUps["missle"].active {
			t = missle
		}
		spawnProjectile(rocketPlayer.x+(rocketRes/2)-(projectileRes/2), rocketYConst-rocketRes, t)
	}

}

func draw() {
	drawProjectiles()
	drawEnemies()
	drawPowerUps()
	rl.DrawTexture(rocketTexture2D, rocketPlayer.x, rocketYConst, color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: uint8(rocketAlpha),
	})
}

func main() {
	rl.InitWindow(windowWidth, windowHeight, "RocketBoom")
	rl.InitAudioDevice()
	loadRessources()

	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()
	defer unloadRessources()
	rl.SetTargetFPS(60)

	defer func() {
		drawScreen("Game Over", "Press space to exit")
	}()
	drawScreen("RocketBoom", "Press Space to start")
	for !rl.WindowShouldClose() && !gameOver {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTexture(backgroundTexture2D, 0, 0, rl.White)
		spawnEnemyWave()
		updateEnemyPos()
		spawnPowerUp()
		updatePowerUpPos()
		updateProjectilePos()
		checkForProjectileHit()
		checkForEnemyHit()
		checkForPowerUpHit()
		updatePowerUps()
		handleInput()

		draw()
		drawHud()
		rl.EndDrawing()
	}

}

func unloadRessources() {
	rl.UnloadSound(laserSFX)
	rl.UnloadSound(explosionSFX)
	rl.UnloadSound(hitSFX)
	rl.UnloadTexture(explosionTexture2D)
	rl.UnloadTexture(heartTexture2D)
	rl.UnloadTexture(backgroundTexture2D)
	rl.UnloadTexture(ufoTexture2D)
	rl.UnloadTexture(rockTexture2D)
	rl.UnloadTexture(rocketTexture2D)
	rl.UnloadTexture(laserTexture2D)
	rl.UnloadTexture(missleTexture2D)
	rl.UnloadTexture(clusterBombTexture2D)

}

func loadRessources() {
	explosionWave := rl.LoadWave(explosionSound)
	explosionSFX = rl.LoadSoundFromWave(explosionWave)

	hitWave := rl.LoadWave(hitSound)
	hitSFX = rl.LoadSoundFromWave(hitWave)

	laserWave := rl.LoadWave(laserSound)
	laserSFX = rl.LoadSoundFromWave(laserWave)

	rl.UnloadWave(hitWave)
	rl.UnloadWave(laserWave)
	rl.UnloadWave(explosionWave)

	missleImg := rl.LoadImage(missleTexture)
	missleTexture2D = rl.LoadTextureFromImage(missleImg)

	rl.ImageResize(missleImg, 16, 16)
	missleSmall = rl.LoadTextureFromImage(missleImg)

	rl.UnloadImage(missleImg)

	laserImg := rl.LoadImage(laserTexture)
	laserTexture2D = rl.LoadTextureFromImage(laserImg)

	rl.ImageResize(laserImg, 32, 32)
	minigunTexture2D = rl.LoadTextureFromImage(laserImg)

	rl.UnloadImage(laserImg)

	clusterBombImg := rl.LoadImage(clusterBombTexture)
	clusterBombTexture2D = rl.LoadTextureFromImage(clusterBombImg)

	rl.UnloadImage(clusterBombImg)

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
	heartTexture2D = rl.LoadTextureFromImage(heartImg)

	explosionImg := rl.LoadImage(explosionTexture)
	explosionTexture2D = rl.LoadTextureFromImage(explosionImg)

	rl.UnloadImage(explosionImg)
	rl.UnloadImage(heartImg)
	rl.UnloadImage(backImg)
	rl.UnloadImage(ufoImg)
	rl.UnloadImage(rockImg)
	rl.UnloadImage(rocketImg)
}
