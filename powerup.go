package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type powerup struct {
	active    bool
	duration  int64
	set       int64
	t         string
	x, y      int32
	invisible bool
	damage    int
}

var (
	powerups     []*powerup
	powerUpsList = []string{clusterBomb, minigun, missle}
	timeLookUp   = map[string]int64{
		missle:  missleTime,
		minigun: minigunTime,
	}
)

func updatePowerUpPos() {
	for _, v := range powerups {
		v.y += enemyMoveSpeed
	}
}

func drawPowerUps() {
	for _, v := range powerups {
		txt := missleTexture2D
		switch v.t {
		case clusterBomb:
			txt = clusterBombTexture2D
			break
		case minigun:
			txt = minigunTexture2D
			break
		}
		rl.DrawTexture(txt, v.x, v.y, rl.White)
		//rl.DrawRectangle(v.x, v.y, powerUpRes, powerUpRes, rl.Red)
	}
}

func spawnPowerUp() {
	if rand.Intn(200) != 1 {
		return
	}

	p := powerUpsList[rand.Intn(len(powerUpsList))]
	damage := 0
	switch p {
	case missle:
		damage = missleDamage
		break
	}
	x := rand.Int31n(windowWidth-10-powerUpRes) + 10
	y := rand.Int31n(enemyYConst)
	powerups = append(powerups, &powerup{
		t:        p,
		x:        x,
		y:        y,
		duration: timeLookUp[p],
		damage:   damage,
	})

}
