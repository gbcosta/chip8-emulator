#ifndef OPENGL_H

#define OPENGL_H
#include <glad/glad.h>

#include <GLFW/glfw3.h>
#include <cstdlib>
#include <iostream>

class Opengl {
private:
  GLFWwindow *window;

  void processInput();

public:
  Opengl();
  void run();
  GLFWwindow *getWindow();
};

#endif
