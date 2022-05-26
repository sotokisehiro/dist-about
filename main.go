package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/code-raisan/gocolor"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func deviceInfoMem() (int, int, float64) {
	vm, _ := mem.VirtualMemory()
	vmT := vm.Total / 100000
	vmF := vm.Free / 100000
	return int(vmT), int(vmF), vm.UsedPercent
}

func main() {
	bytes, err := ioutil.ReadFile("./aa.txt")
	if err != nil {
		panic(err)
	}
	template := string(bytes)
	vmT, vmF, vmP := deviceInfoMem()
	template = strings.Replace(template, "[DEVICE]", "OS"+gocolor.Defalt(": open.Yellow.os")+"\u001b[33m", 1)
	template = strings.Replace(template, "[MEMORY]", "Memory"+gocolor.Defalt(": "+fmt.Sprint(vmT-vmF)+"MiB/"+fmt.Sprint(vmT)+"MiB ("+strconv.FormatFloat(vmP, 'f', 2, 64)+"% Used)")+"\u001b[33m", 1)
	infos, _ := cpu.Info()
	for _, info := range infos {
		template = strings.Replace(template, "[CPU]", strings.Replace("CPU"+gocolor.Defalt(": "+info.ModelName)+"\u001b[33m", "       ", "", 1), 1) // Important space
	}
	println(gocolor.Yellow(template) + "\n")
}
