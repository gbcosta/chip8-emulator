#include <cstdint>
#include <fstream>
#include <ios>
#include <iostream>
#include <iterator>
#include <vector>

#include "./chip8.h"

Chip8::Chip8() {}

void Chip8::resetChip8() {
  this->programCounter = 0x200;
  this->stackPointer = 0;
  this->indexRegister = 0;
  this->soundTimer = 0;
  this->delayTimer = 0;
  this->draw = false;

  for (int i = 0; i < 15; i++) {
    this->registers[i] = 0;
    this->stack[i] = 0;
  }
  int16_t fonts[80] = {
      0xF0, 0x90, 0x90, 0x90, 0xF0, 0x20, 0x60, 0x20, 0x20, 0x70, 0xF0, 0x10,
      0xF0, 0x80, 0xF0, 0xF0, 0x10, 0xF0, 0x10, 0xF0, 0x90, 0x90, 0xF0, 0x10,
      0x10, 0xF0, 0x80, 0xF0, 0x10, 0xF0, 0xF0, 0x80, 0xF0, 0x90, 0xF0, 0xF0,
      0x10, 0x20, 0x40, 0x40, 0xF0, 0x90, 0xF0, 0x90, 0xF0, 0xF0, 0x90, 0xF0,
      0x10, 0xF0, 0xF0, 0x90, 0xF0, 0x90, 0x90, 0xE0, 0x90, 0xE0, 0x90, 0xE0,
      0xF0, 0x80, 0x80, 0x80, 0xF0, 0xE0, 0x90, 0x90, 0x90, 0xE0, 0xF0, 0x80,
      0xF0, 0x80, 0xF0, 0xF0, 0x80, 0xF0, 0x80, 0x80,
  };

  for (int i = 0; i < 80; i++) {
    this->memory[i] = fonts[i];
  }
}

int Chip8::loadRom() {
  std::ifstream romFile("./roms/IBMLogo.ch8", std::ios::binary);
  if (!romFile.is_open()) {
    std::cout << "Not possible open the rom" << std::endl;
    return 1;
  }
  std::vector<char> rom((std::istreambuf_iterator<char>(romFile)),
                        (std::istreambuf_iterator<char>()));

  return 0;
}