package main

type ChipV8 struct{
    memory [4096]byte
    programCounter [2]byte
    indexRegister [2]byte
    stack* byte
    registers [16]uint8
    soundTimer byte
    delayTimer byte
}

var chipV8 ChipV8

func initChipV8() {
    
}
