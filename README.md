# RocketBoom
a very simple but fun shooter written in go using raylib
![image](https://github.com/0x3alex/RocketBoom/assets/90933044/bcfc6e4a-acdc-47d5-9d7f-31ece456ecb5)

## Build
### Linux w. Wayland
On wayland (depending on version, compositor,..) GLFW is not working correctly, which will make the game unplayable. 

If you face this situation, just run the `build.sh` file and execute the compiled program.
### Other
`go build .` or `go run .` is in most cases enough. 

## Changing assets
You can replace every image in the `ressources` folder with your custom images (the file name must stay the same!)
- Background is 800x1200
- Rocket, Rock, Ufo are 64x64
- Heart is 32x32

