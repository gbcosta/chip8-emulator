#ifndef CHIP8_H
#define CHIP8_H

#include <cstdint>

struct OpCode {
  int8_t byte[4];
  int16_t instruction;
};

class Chip8 {
private:
  int8_t memory[4096];
  int16_t programCounter;
  int16_t indexRegister;
  int16_t stackPointer;
  int8_t registers[16];
  int8_t stack[16];
  int8_t soundTimer;
  int8_t delayTimer;
  bool draw;

public:
  Chip8();
  int loadRom();
  void fetch();
  void decode();
  void resetChip8();
};

#endif
