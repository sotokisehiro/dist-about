package main

import (
	"fmt"
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
	template := `                    ......                                    
            XK0OxxddddddddxxkOKKN                             
    X00OxxddddddddddddddddddddddddxxkO0KX                     
+*ddd==dddddddddddddddddddddddxddd00ddddxdk0K*+               
    ?XK0OkxxdddddxddddxdddddddddddxxxkO0K?                    
             -=XK0...xddddd...kO0=-                [DEVICE]
@OKXNKK:            .....              *NK00kx@    [CPU]
dddxddxk+:0KN,,                  WN*00OxxdddddN    [MEMORY]
dddddddddxddddxkO?K*+    +*K0OxddddddddddddddNN               
#/dxddx0*+K-Oxddddxd)    *dddddxddk0KK0kddxdddN               
ddxdx+/    \-W0xddxd)    (dxodxo+*/   +*NOodddN               
dddx/        **KdxdO,    (oxdxx/         \xdxdN               
dxd/          <kddxd)    (ddddx           \xddN               
ddx/           ,ddxd)    (dxdd;           [xddN               
ddd?           ?xdxd)    (dddd?           /-ddN               
dxdx?         =xddxd)    (ddddx0         /-xddX               
dddxk>      <xxddxdO)    (dxdddKL>      <KddxdN               
ddxddx*N*+?=Nxddxdxd)    (dxdddddO*+?L=XOdxdddN               
{dddddddddddddddddxd)    (dxddddddddxxddddxddxW               
  :+KKOkxddddxxddddx)    (dddoxddddddxkO0KW+:/                
           :+0Okxddx*    *ddddxkOKW+:.                        
                 @*+*    *+*@.                                
`
	vmT, vmF, vmP := deviceInfoMem()
	template = strings.Replace(template, "[DEVICE]", "OS"+gocolor.Defalt(": open.Yellow.os")+"\u001b[33m", 1)
	template = strings.Replace(template, "[MEMORY]", "Memory"+gocolor.Defalt(": "+fmt.Sprint((vmT-vmF)/1000)+"MiB/"+fmt.Sprint(vmT/1000)+"MiB ("+strconv.FormatFloat(vmP, 'f', 2, 64)+"% Used)")+"\u001b[33m", 1)
	infos, _ := cpu.Info()
	for _, info := range infos {
		template = strings.Replace(template, "[CPU]", strings.Replace("CPU"+gocolor.Defalt(": "+info.ModelName)+"\u001b[33m", "       ", "", 1), 1) // Important space
	}
	println(gocolor.Yellow(template) + "\n")
}
