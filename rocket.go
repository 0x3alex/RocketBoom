package main

type rocket struct {
	hearts       int
	x            int32
	destroyedObj map[string]int
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
	}
}
