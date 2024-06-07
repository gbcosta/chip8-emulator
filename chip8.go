package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const INSTRUCIONS_PER_TICK = 12

func checkError(e error){
    if(e != nil){
        panic(e)
    }
}

type Opcode struct{
    nibbles [4]byte
    instruction uint16
}

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

var opcode Opcode
var chip8 Chip8

func initChip8() {
    chip8.programCounter = 0x200
    chip8.stackPointer = 0
    chip8.indexRegister = 0
    chip8.soundTimer = 0
    chip8.delayTimer = 0
    chip8.draw = false

    for i := range chip8.registers{
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

    img = ebiten.NewImage(64, 32)
}

func loadROM(){    
    data, err := os.ReadFile("./roms/IBMLogo.ch8")

    checkError(err)
    nextMemoryAdress := 0

    for _, value := range data{
        chip8.memory[chip8.programCounter + uint16(nextMemoryAdress)] = value 
        nextMemoryAdress += 1
    }
}

func chip8Cycle(){
    if(ebiten.ActualTPS() >= 1){
        ticksDeviation := 60 / int(ebiten.ActualTPS())
        for i := 0; i < ticksDeviation * INSTRUCIONS_PER_TICK; i++{
            fetch()
            decode()
        }
    }
}

func fetch(){
    firstOpcodePart := chip8.memory[chip8.programCounter]
    secondOpcodePart := chip8.memory[chip8.programCounter + 1]

    opcode.instruction = uint16(firstOpcodePart) << 8 | uint16(secondOpcodePart)
    
    maskFirstPart := byte(0b11110000)
    maskSecondPart := byte(0b00001111)


    opcode.nibbles[0] = firstOpcodePart & maskFirstPart >> 4
    opcode.nibbles[1] = firstOpcodePart & maskSecondPart 
    opcode.nibbles[2] = secondOpcodePart & maskFirstPart >> 4
    opcode.nibbles[3] = secondOpcodePart & maskSecondPart 

    chip8.programCounter += 2
}

func decode(){
    switch opcode.nibbles[0] {
    case 0:
        clearScreen()
    case 1: 
        jump()
    case 6:
        setRegister()
    case 7:
        addValueToResgister()
    case 0xA:
        setIndexRegister()
    case 0xD:
        draw()
}
}

func execute(){
    
}
func clearScreen(){
    img.Clear()
    img = ebiten.NewImage(64, 32)
}

func jump(){
    chip8.programCounter = uint16(opcode.nibbles[1]) << 8 | 
    uint16(opcode.nibbles[2]) << 4 | uint16(opcode.nibbles[3])
}

func setRegister(){
    chip8.registers[uint8(opcode.nibbles[1])] = opcode.nibbles[2] << 4 | opcode.nibbles[3]
}

func addValueToResgister(){
    vx := opcode.nibbles[1]
    chip8.registers[int(vx)] += opcode.nibbles[2] << 4 | opcode.nibbles[3]
}

func setIndexRegister(){
    chip8.indexRegister = uint16(opcode.nibbles[1]) << 8 | 
    uint16(opcode.nibbles[2]) << 4 | uint16(opcode.nibbles[3])
}

func draw(){
    vx := chip8.registers[uint8(opcode.nibbles[1])] & 63
    vy := chip8.registers[uint8(opcode.nibbles[2])] & 31
    nSprites := int(opcode.nibbles[3])

    chip8.registers[0xf] = 0

    red := color.RGBA{255, 0, 0, 255}
    
    for i := 0; i < nSprites; i++{
        if int(vy) + i > 31{
            break
        }

        mask := byte(128)
        for j:= 0; j < 8; j++{
            maskApplied := chip8.memory[int(chip8.indexRegister) + i] & mask
            if (maskApplied == mask && int(vx) + j < 64){
                img.Set(int(vx) + j, int(vy) + i, red)
            }
            mask /= 2
        }
    }
    chip8.draw = true
}

