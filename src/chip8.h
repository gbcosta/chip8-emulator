#ifndef CHIP8_H
#include <cstdint>
#define CHIP8_H

struct Opcode {
  uint8_t byte[4];
  uint16_t instruction;
};

class Chip8 {
private:
  uint8_t memory[4096];
  uint16_t PC;
  uint16_t indexRegister;
  uint16_t stackPointer;
  uint8_t registers[16];
  uint16_t stack[16];
  uint8_t soundTimer;
  uint8_t delayTimer;
  bool draw;
  Opcode opcode;

public:
  Chip8();
  int loadRom();
  void fetch();
  void decode();
  void execute();
  void resetChip8();
};

#endif
