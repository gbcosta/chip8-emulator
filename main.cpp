#include "./src/chip8.h"

int main() {
  Chip8 chip8;
  chip8.loadRom();
  chip8.run();

  glfwTerminate();
}
