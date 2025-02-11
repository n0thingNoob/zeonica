package core

import (
	"math"
	"strconv"
	"strings"
)

type coreState struct {
	PC               uint32
	TileX, TileY     uint32
	Registers        []uint32
	Code             []string
	RecvBufHead      []uint32
	RecvBufHeadReady []bool
	SendBufHead      []uint32
	SendBufHeadBusy  []bool
}

type instEmulator struct {
}

func (i instEmulator) RunInst(inst string, state *coreState) {
	tokens := strings.Split(inst, ",")
	for i := range tokens {
		tokens[i] = strings.TrimSpace(tokens[i])
	}

	instName := tokens[0]
	if strings.Contains(instName, "CMP") {
		instName = "CMP"
	}

	instFuncs := map[string]func([]string, *coreState){
		"WAIT": i.runWait,
		"SEND": i.runSend,
		"JMP":  i.runJmp,
		"CMP":  i.runCmp,
		"JEQ":  i.runJeq,
		"DONE": func(_ []string, _ *coreState) { i.runDone() }, // Since runDone might not have parameters
	}

	if instFunc, ok := instFuncs[instName]; ok {
		instFunc(tokens, state)
	} else {
		panic("unknown instruction " + inst)
	}
}

func (i instEmulator) runWait(inst []string, state *coreState) {
	dst := inst[1]
	src := inst[2]

	i.waitSrcMustBeNetRecvReg(src)
	srcIndex, err := strconv.Atoi(strings.TrimPrefix(src, "NET_RECV_"))
	if err != nil {
		panic(err)
	}

	if !state.RecvBufHeadReady[srcIndex] {
		return
	}

	state.RecvBufHeadReady[srcIndex] = false
	i.writeOperand(dst, state.RecvBufHead[srcIndex], state)
	state.PC++
}

func (i instEmulator) waitSrcMustBeNetRecvReg(src string) {
	if !strings.HasPrefix(src, "NET_RECV_") {
		panic("the source of a WAIT instruction must be NET_RECV registers")
	}
}

func (i instEmulator) runSend(inst []string, state *coreState) {
	dst := inst[1]
	src := inst[2]

	i.sendDstMustBeNetSendReg(dst)
	dstIndex, err := strconv.Atoi(strings.TrimPrefix(dst, "NET_SEND_"))
	if err != nil {
		panic(err)
	}

	if state.SendBufHeadBusy[dstIndex] {
		return
	}

	state.SendBufHeadBusy[dstIndex] = true
	val := i.readOperand(src, state)
	state.SendBufHead[dstIndex] = val
	state.PC++
}

func (i instEmulator) sendDstMustBeNetSendReg(dst string) {
	if !strings.HasPrefix(dst, "NET_SEND_") {
		panic("the destination of a SEND instruction must be NET_SEND registers")
	}
}

func (i instEmulator) runJmp(inst []string, state *coreState) {
	dst := inst[1]

	for i := 0; i < len(state.Code); i++ {
		line := strings.Trim(state.Code[i], " \t\n")
		if strings.HasPrefix(line, dst) && strings.HasSuffix(line, ":") {
			state.PC = uint32(i)
			return
		}
	}
}

func (i instEmulator) readOperand(operand string, state *coreState) (value uint32) {
	if strings.HasPrefix(operand, "$") {
		registerIndex, err := strconv.Atoi(strings.TrimPrefix(operand, "$"))
		if err != nil {
			panic("invalid register index")
		}

		value = state.Registers[registerIndex]
	}

	return
}

func (i instEmulator) writeOperand(operand string, value uint32, state *coreState) {
	if strings.HasPrefix(operand, "$") {
		registerIndex, err := strconv.Atoi(strings.TrimPrefix(operand, "$"))
		if err != nil {
			panic("invalid register index")
		}

		state.Registers[registerIndex] = value
	}
}

func (i instEmulator) runCmp(inst []string, state *coreState) {
	Itype := inst[0]
	//Float or Integer
	switch {
	case strings.Contains(Itype, "I"):
		i.parseAndCompareI(inst, state)
	case strings.Contains(Itype, "F32"):
		i.parseAndCompareF32(inst, state)
	default:
		panic("invalid cmp")
	}
}

func (i instEmulator) parseAndCompareI(inst []string, state *coreState) {
	instruction := inst[0]
	dst := inst[1]
	src := inst[2]

	srcVal := i.readOperand(src, state)
	dstVal := uint32(0)
	imme, err := strconv.ParseUint(inst[3], 10, 32)
	if err != nil {
		panic("invalid compare number")
	}

	immeI32 := int32(uint32(imme))
	srcValI := int32(srcVal)

	conditionFuncs := map[string]func(int32, int32) bool{
		"EQ": func(a, b int32) bool { return a == b },
		"NE": func(a, b int32) bool { return a != b },
		"LE": func(a, b int32) bool { return a <= b },
		"LT": func(a, b int32) bool { return a < b },
		"GT": func(a, b int32) bool { return a > b },
		"GE": func(a, b int32) bool { return a >= b },
	}

	for key, function := range conditionFuncs {
		if strings.Contains(instruction, key) && function(srcValI, immeI32) {
			dstVal = 1
		}
	}
	i.writeOperand(dst, dstVal, state)
	state.PC++
}

func (i instEmulator) parseAndCompareF32(inst []string, state *coreState) {
	instruction := inst[0]
	dst := inst[1]
	src := inst[2]

	srcVal := i.readOperand(src, state)
	dstVal := uint32(0)
	imme, err := strconv.ParseUint(inst[3], 10, 32)
	if err != nil {
		panic("invalid compare number")
	}

	conditionFuncsF := map[string]func(float32, float32) bool{
		"EQ": func(a, b float32) bool { return a == b },
		"NE": func(a, b float32) bool { return a != b },
		"LT": func(a, b float32) bool { return a < b },
		"LE": func(a, b float32) bool { return a <= b },
		"GT": func(a, b float32) bool { return a > b },
		"GE": func(a, b float32) bool { return a >= b },
	}

	immeF32 := math.Float32frombits(uint32(imme))
	srcValF := math.Float32frombits(srcVal)

	for key, function := range conditionFuncsF {
		if strings.Contains(instruction, key) && function(srcValF, immeF32) {
			dstVal = 1
		}
	}
	i.writeOperand(dst, dstVal, state)
	state.PC++
}

func (i instEmulator) runJeq(inst []string, state *coreState) {
	src := inst[2]
	imme, err := strconv.ParseUint(inst[3], 10, 32)

	if err != nil {
		panic("invalid compare number")
	}

	srcVal := i.readOperand(src, state)

	if srcVal == uint32(imme) {
		i.runJmp(inst, state)
	} else {
		state.PC++
	}
}

func (i instEmulator) runDone() {
	// Do nothing.
}
