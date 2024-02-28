package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type enemy struct {
	t             string
	hp            int
	x, y          int32
	invisible     bool
	explosionTime int
}

var (
	enemies  []*enemy
	hpLookUp = map[string]int{
		rock: rockHp,
		ufo:  ufoHp,
	}
	lastSpawn = time.Now().Unix()
)

func drawEnemies() {
	for _, v := range enemies {

		if v.invisible && v.explosionTime != 0 {
			v.explosionTime -= 1
			rl.DrawTexture(explosionTexture2D, v.x, v.y, rl.White)
			continue
		}
		if v.invisible {
			continue
		}
		switch v.t {
		case rock:
			rl.DrawTexture(rockTexture2D, v.x, v.y, rl.White)
			break
		case ufo:
			rl.DrawTexture(ufoTexture2D, v.x, v.y, rl.White)
			break
		}
	}
}

func updateEnemyPos() {
	for i, v := range enemies {
		if v.invisible {
			continue
		}
		v.y += enemyMoveSpeed
		if v.y > windowHeight {
			t := enemies[:i]
			if i+1 < len(enemies) {
				t = append(t, enemies[i+1:]...)
			}
			enemies = t
		}
	}
}

func spawnEnemyWave() {
	if time.Now().Unix()-lastSpawn < 1 {
		return
	}

	var newWave []*enemy
	for i := 0; i < rand.Intn(5)+4; i++ {
	retry:
		x := rand.Int31n(windowWidth-10-enemyRes) + 10
		r1 := rl.Rectangle{
			X:      float32(x),
			Y:      0,
			Width:  float32(enemyRes),
			Height: float32(enemyRes),
		}
		for _, v := range newWave {
			r2 := rl.Rectangle{
				X:      float32(v.x),
				Y:      float32(v.y),
				Width:  float32(enemyRes),
				Height: float32(enemyRes),
			}
			if rl.CheckCollisionRecs(r1, r2) {
				goto retry
			}
		}

		n := rand.Int31n(20)
		t := rock
		if n <= 2 {
			t = ufo
		}

		newWave = append(newWave, &enemy{
			x:             x,
			y:             rand.Int31n(enemyYConst),
			t:             t,
			hp:            hpLookUp[t],
			explosionTime: 3,
		})
	}
	enemies = append(enemies, newWave...)
	lastSpawn = time.Now().Unix()

}
