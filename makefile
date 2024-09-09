compiler = gcc
main: main.cpp
	 ${compiler} main.cpp -o main ./src/glad.c -lopengl32 -lglfw3 -lgdi32 -lstdc++
