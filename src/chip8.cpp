#include <fstream>
#include <ios>
#include <iostream>
#include <iterator>
#include <vector>

#include "./chip8.h"

Chip8::Chip8() {
  resetChip8();
  opengl = new Opengl();
}

void Chip8::resetChip8() {
  this->PC = 0x200;
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
    std::cerr << "Not possible open the rom" << std::endl;
    return 1;
  }

  std::vector<uint8_t> rom((std::istreambuf_iterator<char>(romFile)),
                           (std::istreambuf_iterator<char>()));

  for (int i = 0; i < rom.size(); i++) {
    this->memory[i + 0x200] = rom[i];
  }

  for (int i = 0; i < 5; i++) {
    fetch();
    decode();
  }

  return 0;
}

void Chip8::fetch() {
  uint8_t opcodePartOne = this->memory[this->PC];
  uint8_t opcodePartTwo = this->memory[this->PC + 1];

  this->opcode.instruction = (opcodePartOne << 8) | opcodePartTwo;

  this->opcode.nibbles[0] = (opcodePartOne & 0b11110000) >> 4;
  this->opcode.nibbles[1] = opcodePartOne & 0b00001111;
  this->opcode.nibbles[2] = (opcodePartOne & 0b11110000) >> 4;
  this->opcode.nibbles[3] = opcodePartOne & 0b00001111;

  this->PC += 2;
}

void Chip8::run() {
  while (!glfwWindowShouldClose(opengl->getWindow())) {
    opengl->run();
  }
}

void Chip8::decode() {

  switch (this->opcode.nibbles[0]) {
  case 0:
    std::cout << "clear" << std::endl;
    break;
  case 0xd:
    std::cout << "draw" << std::endl;
  case 0xa:
    std::cout << "set register i" << std::endl;
    break;
  case 0x6:
    std::cout << "set register 6" << std::endl;
    break;
  case 0x7:
    std::cout << "set register 7" << std::endl;
    break;
  }
}
