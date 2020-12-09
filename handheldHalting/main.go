package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := `acc +48
nop +308
acc +33
acc +48
jmp +379
acc +27
acc +23
acc +29
acc +3
jmp +326
acc -12
nop +248
nop +146
jmp +293
acc +8
acc -16
acc +18
jmp +255
jmp +390
jmp +442
acc +38
jmp +451
jmp +499
acc +41
acc +44
nop +298
acc +31
jmp +337
acc -3
acc +34
nop +266
acc -19
jmp +281
jmp +232
acc -9
acc +25
acc -8
nop +228
jmp +79
jmp +58
acc +38
nop +318
jmp +17
acc +44
jmp +298
acc -4
jmp +196
jmp +249
acc -9
acc -3
acc +17
jmp +556
jmp +1
acc -15
jmp +166
acc +23
acc +10
acc +39
jmp +475
acc +28
jmp +41
jmp +322
acc +48
acc -11
acc -13
acc +12
jmp +92
jmp +1
acc +9
jmp -26
acc +4
acc -8
nop +484
acc -14
jmp +526
acc -13
acc -2
acc -13
jmp +419
acc +22
acc -17
jmp +1
jmp +151
acc +3
acc +2
acc -13
acc -11
jmp +352
acc +24
jmp -27
nop +206
acc +22
nop +225
jmp +360
acc +48
jmp +105
jmp +80
acc -16
jmp +89
acc +0
jmp +339
acc +37
acc +41
nop +156
jmp +452
jmp +208
nop +60
jmp +155
acc -16
jmp +274
nop -77
acc -16
acc +21
jmp +508
acc -7
acc -1
jmp -83
acc -11
jmp +28
acc +3
acc +32
acc -18
jmp -89
acc +4
acc -14
acc -1
acc +33
jmp -75
nop +185
acc +8
acc +22
acc +26
jmp +164
acc -19
jmp +260
jmp +174
acc +2
acc +34
jmp -120
acc -15
acc +48
jmp +165
acc +5
nop +240
jmp -121
jmp +114
jmp -36
nop +432
jmp +1
acc -18
jmp +429
acc +44
nop +110
jmp +198
acc -18
acc +32
acc +41
jmp +102
nop +177
acc +35
acc +24
acc +46
jmp +121
acc +20
jmp +1
jmp +407
jmp +1
acc +35
acc +46
nop -18
jmp +6
jmp +422
acc -19
jmp -85
acc +33
jmp -116
nop +79
jmp +284
acc +3
acc +49
nop +317
jmp +6
acc +6
jmp +295
nop -141
acc -4
jmp -44
nop +155
acc +48
acc -17
jmp +188
acc +22
jmp +286
nop +103
acc -2
acc +45
jmp +20
acc +21
acc +37
jmp +235
jmp +42
acc -11
nop -15
acc -5
jmp +235
jmp +178
acc +12
acc -15
jmp +25
acc +9
acc +11
nop +389
acc +50
jmp +146
acc +26
jmp +144
acc -14
jmp +304
nop +254
jmp +337
jmp +17
jmp +1
acc +6
acc -4
acc +42
jmp +117
acc +25
acc +50
acc +45
jmp -112
acc +14
acc +27
jmp +347
nop +15
jmp +14
acc +29
jmp +236
jmp -71
acc -11
acc +21
nop +32
jmp -162
acc -15
jmp +322
acc -4
acc +16
jmp +1
jmp +100
jmp +1
jmp -77
acc +21
nop -199
acc +49
acc -1
jmp -231
jmp +230
acc -19
jmp +1
jmp -49
jmp -11
acc +6
jmp -110
jmp +331
acc +44
jmp +292
acc -7
acc -18
acc +50
jmp +221
acc +33
acc +7
jmp -45
jmp +342
acc -19
acc +36
acc +15
jmp -229
nop -5
jmp +57
acc +26
acc +43
nop -175
jmp +82
acc +45
jmp -161
acc -16
acc +35
acc +46
acc +43
jmp +1
nop +195
acc +39
acc +27
acc +32
jmp +227
jmp -272
nop +201
acc +6
acc +13
acc +12
jmp -177
acc -9
acc +46
nop +199
acc -1
jmp +1
jmp +3
acc +42
jmp +75
jmp +305
acc +49
acc -16
jmp -92
acc +3
nop +279
jmp +54
jmp +31
acc +50
jmp -125
acc +21
nop -178
acc +40
jmp +193
acc +39
acc -5
jmp +261
nop -3
acc -13
jmp -310
acc +6
acc -17
acc +12
acc +38
jmp +267
jmp -311
acc -2
jmp -7
nop +77
acc -2
acc +39
acc -16
jmp +10
nop +59
jmp -296
acc -4
acc +41
jmp -249
acc +43
nop +35
jmp +95
jmp +171
acc +10
nop +169
acc -17
jmp +47
acc +49
acc +38
nop +199
jmp +249
jmp -53
nop -194
acc +19
acc +18
jmp -16
acc +33
jmp +194
nop -194
acc +49
jmp +85
acc +50
nop -318
acc -7
jmp -49
acc -6
acc +48
acc -13
acc -14
jmp +67
acc +12
acc +19
acc +3
jmp -371
jmp -149
acc +49
nop -202
jmp -315
acc -6
jmp -171
acc -7
jmp +113
acc +34
acc +36
acc +17
jmp -97
acc +3
jmp -244
acc +25
acc +30
acc +25
acc -19
jmp +44
nop +84
jmp +124
nop +17
acc -11
acc -8
acc +4
jmp +193
jmp -388
acc +36
acc +17
jmp +1
acc -5
jmp +166
acc +39
acc -10
jmp -280
acc +15
jmp +1
jmp -396
jmp +113
acc +37
acc +13
jmp -35
nop +109
acc +8
acc +6
acc +19
jmp +39
jmp +1
jmp +1
acc +39
acc -14
jmp -291
acc +39
acc +31
jmp -231
acc +41
jmp -55
nop -167
jmp +105
acc -8
acc +34
jmp -114
nop +58
jmp +1
nop -270
acc +31
jmp -135
acc +8
acc +33
jmp +1
jmp -64
acc +24
acc +16
jmp +117
acc +35
acc -11
nop +44
acc +18
jmp +1
jmp +88
acc +29
nop +34
jmp +1
nop -118
jmp -404
jmp -144
acc -12
nop -372
acc -14
jmp -209
acc +12
jmp +1
acc -1
jmp +132
nop -93
jmp -130
acc +23
acc +30
acc +3
jmp -209
nop -381
acc -19
jmp +23
nop +87
jmp -277
acc +39
jmp -391
acc +14
acc +18
acc +24
nop -459
jmp -267
acc +35
nop +84
jmp -231
acc +5
acc +0
acc +45
jmp -210
jmp -211
acc +7
acc +8
nop -249
jmp -8
jmp -105
nop -455
acc -19
acc +36
jmp -368
acc +33
acc +10
acc +9
jmp -259
nop +41
acc -14
acc +2
jmp -336
acc +46
jmp -261
nop -284
acc +21
nop -154
jmp -485
jmp -505
acc +32
nop -327
acc +1
nop +43
jmp -23
acc -7
jmp -88
acc +10
jmp -440
acc +12
jmp -430
jmp +1
acc +46
nop -105
jmp -87
acc +49
acc -10
acc -6
jmp -411
jmp -268
acc +35
acc +15
jmp +45
acc -14
acc +0
jmp -234
nop -67
acc +32
acc +1
jmp -476
jmp -297
nop -274
jmp -435
acc +36
acc -2
acc +33
acc +9
jmp -26
jmp +50
acc +23
jmp -172
jmp +1
acc +34
acc +32
acc -4
jmp -312
acc -11
acc +26
jmp -150
acc +41
nop -79
acc +25
jmp -76
acc -2
acc +29
acc +12
jmp -549
nop -357
nop -438
jmp -320
acc +7
acc -6
jmp -149
nop -74
acc +7
acc +45
jmp -383
acc -8
acc +5
acc +12
jmp -463
acc -17
acc +5
jmp -34
jmp +1
acc +18
jmp -523
acc +38
acc +35
nop -222
jmp -424
nop -365
jmp -188
acc +0
acc +7
nop -167
acc +27
jmp -351
acc +26
jmp -201
nop -208
nop -466
acc +50
nop -531
jmp -273
acc +21
acc +25
nop -397
acc +29
jmp +1`

	instructions := createInstructions(input)

	acc, _ := breakInfiniteLoop(instructions)
	fmt.Println(acc)

	for i := 1; ; i++ {
		modifiedInstructions := replaceInstruction(instructions, i)
		acc, found := breakInfiniteLoop(modifiedInstructions)
		if !found {
			fmt.Println(acc)
			break
		}
	}
}

func createInstructions(input string) []Instruction {
	var result []Instruction
	instructions := strings.Split(input, "\n")
	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		argument, _ := strconv.Atoi(parts[1])
		result = append(result, Instruction{parts[0], argument})
	}
	return result
}

type Instruction struct {
	name     string
	argument int
}

func replaceInstruction(instructions []Instruction, i int) []Instruction {
	relativeIndex := 0
	result := make([]Instruction, len(instructions))
	copy(result, instructions)
	for absoluteIndex, instruction := range result {
		if instruction.name == "nop" || instruction.name == "jmp" {
			relativeIndex++
		}
		if relativeIndex == i {
			switch result[absoluteIndex].name {
			case "nop":
				result[absoluteIndex].name = "jmp"
			case "jmp":
				result[absoluteIndex].name = "nop"
			}
			break
		}
	}
	return result
}

func breakInfiniteLoop(instructions []Instruction) (int, bool) {
	acc := 0
	index := 0
	visited := map[int]bool{}

	found := false

	for {
		_, ok := visited[index]
		if ok {
			found = true
			break
		}
		visited[index] = true
		acc, index = interpretInstruction(acc, index, instructions)
		if index >= len(instructions) {
			break
		}
	}
	return acc, found
}

func interpretInstruction(acc int, index int, instructions []Instruction) (int, int) {
	switch instructions[index].name {
	case "acc":
		acc += instructions[index].argument
		index += 1
	case "jmp":
		index += instructions[index].argument
	case "nop":
		index += 1
	}

	return acc, index
}