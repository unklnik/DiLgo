### DiLgo >> Utility function for SDL2 and Go

**NOTE:** That this will not work just by building, the SDL2.dll file, SDL2_image.dll file and SDL2_ttf.dll files must be in the working directory. Unfortunately the SDL2_ttf.dll file is too large to upload to GitHub so you will need to get it here https://github.com/libsdl-org/SDL_ttf/releases. Just download the package, open the archive and copy only the SDL2_ttf.dll into the working directory (where all the other files of the program are). 

Also, you need to be on the master branch of SDL2 bindings otherwise there will be build errors so once SDL2 is installed then run:<br>
```go get -v github.com/veandco/go-sdl2/sdl@master``` 

Then you should be able to build and run the program.

Created with the SDL2 bindings for Go here https://github.com/veandco/go-sdl2

DiLgo is an unfinished project that was created to make it easier to use SDL2 with Go for game development. Getting SDL2 to work can be tricky and DiLgo hopefully makes it a bit easier to draw shapes and animations and create grids. Unfortunately time and patience got the better of me and the project is incomplete however it may be of some use to someone. Note that sound is not implemented. 

https://github.com/user-attachments/assets/8031ba7d-875c-4ccc-be04-9c8f2c469218

