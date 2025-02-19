compiler = gcc
main: main.cpp
	 ${compiler} main.cpp -o main ./src/glad.c ./src/opengl.cpp ./src/chip8.cpp -lopengl32 -lglfw3 -lgdi32 -lstdc++
