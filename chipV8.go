package main

type Chip8 struct{
    memory [4096]byte
    programCounter [2]byte
    indexRegister [2]byte
    stack* byte
    registers [16]uint8
    soundTimer byte
    delayTimer byte
}

var chip8 Chip8

func initChip8() {
    
}
