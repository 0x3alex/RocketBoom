package main

var (
	windowWidth       = int32(800)
	windowHeight      = int32(1200)
	rocketMoveSpeed   = int32(5)
	rocketRes         = int32(64)
	enemyRes          = int32(64)
	enemyYConst       = int32(enemyRes / 2)
	rocketYConst      = windowHeight - (6 * rocketRes)
	laserSpeed        = int32(20)
	laserWidth        = int32(20)
	laserHeight       = int32(40)
	rockHp            = 3
	ufoHp             = 7
	enemyMoveSpeed    = int32(3)
	hitDelay          = int64(3)
	explosionTime     = 4
	rock              = "rock"
	ufo               = "ufo"
	rocketTexture     = "./ressources/rocket.png"
	backgroundTexture = "./ressources/background.png"
	ufoTexture        = "./ressources/ufo.png"
	rockTexture       = "./ressources/rock.png"
	heartTexture      = "./ressources/heart.png"
	explosionTexture  = "./ressources/explosion.png"
	explosionSound    = "./ressources/explosion.wav"
	laserSound        = "./ressources/laser.wav"
	hitSound          = "./ressources/hit.wav"
)
