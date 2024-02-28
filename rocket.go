package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type rocket struct {
	hearts       int
	x            int32
	destroyedObj map[string]int
	powerUps     map[string]powerup
	hit          bool
}

func initRocket() *rocket {
	return &rocket{
		hearts: 3,
		x:      (windowWidth / 2) - (rocketRes / 2),
		destroyedObj: map[string]int{
			"rock": 0,
			"ufo":  0,
		},
		powerUps: map[string]powerup{
			"missle": powerup{
				active: false,
			},
			"minigun": powerup{
				active: false,
			},
		},
	}
}

func updatePowerUps() {
	for k, v := range rocketPlayer.powerUps {
		if v.active {

			if time.Now().Unix()-v.set >= v.duration {
				if v.t == minigun {
					projectiles = nil
				}
				rocketPlayer.powerUps[k] = powerup{}
			}
		}
	}
}

func checkForPowerUpHit() {

	rBox := rl.Rectangle{
		X:      float32(rocketPlayer.x),
		Y:      float32(rocketYConst),
		Width:  float32(rocketRes) - 10,
		Height: float32(rocketRes) - 10,
	}

	for i, p := range powerups {
		pBox := rl.Rectangle{
			X:      float32(p.x),
			Y:      float32(p.y),
			Width:  float32(powerUpRes),
			Height: float32(powerUpRes),
		}
		if rl.CheckCollisionRecs(rBox, pBox) {
			switch p.t {
			case clusterBomb:
				for _, v := range enemies {
					v.hp = 0
					v.invisible = true
				}
				rl.PlaySound(explosionSFX)
				break
			default:
				rocketPlayer.powerUps[p.t] = powerup{
					active:   true,
					t:        p.t,
					damage:   p.damage,
					duration: p.duration,
					set:      time.Now().Unix(),
				}
				break
			}
			powerups = append(powerups[:i], powerups[i+1:]...)
		}
	}
}

func checkForEnemyHit() {
	if time.Now().Unix()-lastHit >= hitDelay {
		rocketPlayer.hit = false
		rocketAlpha = 255
	}
	if rocketPlayer.hit {
		if rocketAlpha == 0 {
			rocketAlpha = 255
			return
		}
		rocketAlpha -= 10
		return
	}
	rBox := rl.Rectangle{
		X:      float32(rocketPlayer.x),
		Y:      float32(rocketYConst),
		Width:  float32(rocketRes),
		Height: float32(rocketRes),
	}
	for _, v := range enemies {
		if v.invisible {
			continue
		}
		eBox := rl.Rectangle{
			X:      float32(v.x),
			Y:      float32(v.y),
			Width:  float32(enemyRes),
			Height: float32(enemyRes),
		}
		if rl.CheckCollisionRecs(rBox, eBox) {
			rl.PlaySound(hitSFX)
			lastHit = time.Now().Unix()
			rocketPlayer.hit = true
			rocketPlayer.hearts -= 1
			if rocketPlayer.hearts == 0 {
				gameOver = true
			}
		}
	}
}
