package main

import rl "github.com/gen2brain/raylib-go/raylib"

type projectile struct {
	x, y   int32
	damage int
	t      string
}

var (
	projectiles  []*projectile
	damageLookUp = map[string]int{
		laser:  laserDamage,
		missle: missleDamage,
	}
)

func drawProjectiles() {
	for _, v := range projectiles {
		txt := laserTexture2D
		switch v.t {
		case missle:
			txt = missleTexture2D
			break
		}
		rl.DrawTexture(txt, v.x, v.y, rl.White)
		//rl.DrawRectangle(v.x, v.y, projectileRes, projectileRes, rl.Red)
	}
}

func spawnProjectile(x, y int32, t string) {
	projectiles = append(projectiles, &projectile{
		x:      x,
		y:      y,
		damage: damageLookUp[t],
		t:      t,
	})
}

func updateProjectilePos() {
	for i, v := range projectiles {
		v.y -= laserSpeed
		if v.y < 0 && !rocketPlayer.powerUps["minigun"].active {
			t := projectiles[:i]
			if i+1 < len(projectiles) {
				t = append(t, projectiles[i+1:]...)
			}
			projectiles = t
		}
	}
}

func checkForProjectileHit() {
	for i, projectile := range projectiles {
		pBox := rl.Rectangle{
			X:      float32(projectile.x),
			Y:      float32(projectile.y),
			Width:  float32(projectileRes),
			Height: float32(projectileRes),
		}
		for _, enemy := range enemies {
			if enemy.invisible {
				continue
			}
			eBox := rl.Rectangle{
				X:      float32(enemy.x),
				Y:      float32(enemy.y),
				Width:  float32(enemyRes),
				Height: float32(enemyRes),
			}
			if rl.CheckCollisionRecs(pBox, eBox) {
				t := projectiles[:i]
				if i+1 < len(projectiles) {
					t = append(t, projectiles[i+1:]...)
				}
				projectiles = t
				enemy.hp -= projectile.damage
				if enemy.hp <= 0 {
					enemy.invisible = true
					rocketPlayer.destroyedObj[enemy.t] += 1
				}
			}
		}
	}
}
