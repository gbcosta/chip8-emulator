package main

import (
  "os"
)

type Chip8 struct{
  memory [4096]byte
  programCounter uint16
  indexRegister uint16
  stack[16]uint16
  stackPointer uint16
  registers [16]byte
  soundTimer byte
  delayTimer byte
  draw bool
}

func checkError(e error){
  if(e != nil){
    panic(e)
  }
}

func clearScreen(){
}

func jump(){
}

func setRegister(){
}

func addValueToResgister(){
}

func setIndexRegiste(){
}

func draw(){
}

//init 8
//load game
//instructions
//v8 cycle
//draw
//get keys



var chip8 Chip8

func initChip8() {
  chip8.programCounter = 0x200
  chip8.stackPointer = 0
  chip8.indexRegister = 0
  chip8.soundTimer = 0
  chip8.delayTimer = 0
  chip8.draw = false

  for i, _ := range chip8.registers{
    chip8.registers[i] = 0
    chip8.stack[i] = 0
  }

  fonts := [80]byte{
    0xF0, 0x90, 0x90, 0x90, 0xF0, 
    0x20, 0x60, 0x20, 0x20, 0x70,
    0xF0, 0x10, 0xF0, 0x80, 0xF0,
    0xF0, 0x10, 0xF0, 0x10, 0xF0,
    0x90, 0x90, 0xF0, 0x10, 0x10,
    0xF0, 0x80, 0xF0, 0x10, 0xF0,
    0xF0, 0x80, 0xF0, 0x90, 0xF0,
    0xF0, 0x10, 0x20, 0x40, 0x40,
    0xF0, 0x90, 0xF0, 0x90, 0xF0,
    0xF0, 0x90, 0xF0, 0x10, 0xF0,
    0xF0, 0x90, 0xF0, 0x90, 0x90,
    0xE0, 0x90, 0xE0, 0x90, 0xE0,
    0xF0, 0x80, 0x80, 0x80, 0xF0,
    0xE0, 0x90, 0x90, 0x90, 0xE0,
    0xF0, 0x80, 0xF0, 0x80, 0xF0,
    0xF0, 0x80, 0xF0, 0x80, 0x80,
  }

  for i, font := range fonts{
    chip8.memory[i] = font
  }
}

func loadROM(){    
  data, err := os.ReadFile("./roms/IBMLogo.ch8")

  checkError(err)
  nextMemoryAdress := 0

  for value := range data{
    chip8.memory[chip8.programCounter + uint16(nextMemoryAdress)] = byte(value)
    nextMemoryAdress += 1
  }
}


func chip8Cycle(){
  
}

