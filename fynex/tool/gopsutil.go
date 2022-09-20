package tool

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"

	// "github.com/shirou/gopsutil/v3/winservices" // only for windows
	"github.com/tidwall/pretty"
	// "github.com/shirou/gopsutil/mem"  // to use v2
)

func goPsUtil() {
	w := fyne.CurrentApp().NewWindow("Html to txt")

	opts := []string{"CPU", "VirtualMemory", "SwapDevices", "SwapMemory",
		"Net", "Disk", "DiskUsage", "IOCounter",
		"Host", "Process",
		// "WinServices"
	}
	mapp := make(map[string]string, len(opts))
	entry := widget.NewMultiLineEntry()
	v, _ := mem.VirtualMemory()
	vm := string(pretty.Pretty([]byte(v.String())))
	hostInfo, _ := host.Info()
	hostInfoStr := string(pretty.Pretty([]byte(hostInfo.String())))

	cpuInfo, _ := cpu.Info()
	cpuString := ""
	for _, c := range cpuInfo {
		cpuString = cpuString + string(pretty.Pretty([]byte(c.String())))
	}
	buf := bytes.NewBuffer(nil)
	parts, _ := disk.Partitions(true)
	diskUsageStr := ""
	ioCountersStr := ""
	for _, p := range parts {
		du, err := disk.Usage(p.Device)
		if err != nil {
			diskUsageStr = "can not get disk usage"
			break
		}
		diskUsageStr = diskUsageStr + string(pretty.Pretty([]byte(du.String())))

		dm, _ := disk.IOCounters(p.Device)
		ioCountersStr = ioCountersStr + string(pretty.Pretty([]byte(dm[p.Device].String())))
		buf.Write(pretty.Pretty([]byte(p.String())))
	}
	mapp["Disk"] = buf.String()
	mapp["DiskUsage"] = diskUsageStr
	mapp["IOCounter"] = ioCountersStr

	netList, _ := net.Interfaces()
	netListStr := pretty.Pretty([]byte(netList.String()))

	procList, _ := process.Processes()
	buf.Reset()
	for _, p := range procList {
		buf.Write(pretty.Pretty([]byte(p.String())))
	}
	mapp["Process"] = buf.String()

	// services, _ := winservices.ListServices()
	// buf.Reset()
	// for _, v := range services {
	// 	buf.Write(pretty.Pretty([]byte(gconv.String(v))))
	// }
	// mapp["WinServices"] = buf.String()

	swapDevices, _ := mem.SwapDevices()
	buf.Reset()
	for _, v := range swapDevices {
		buf.Write(pretty.Pretty([]byte(v.String())))
	}
	mapp["SwapDevices"] = buf.String()

	swapMem, _ := mem.SwapMemory()
	mapp["SwapMemory"] = string(pretty.Pretty([]byte(swapMem.String())))

	mapp["CPU"] = cpuString
	mapp["VirtualMemory"] = vm
	mapp["Host"] = hostInfoStr
	mapp["Net"] = string(netListStr)
	lst := widget.NewList(func() int {
		return len(opts)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("infoinfoinfoinfoinfo")
	}, func(lii widget.ListItemID, co fyne.CanvasObject) {
		lbl := co.(*widget.Label)
		lbl.SetText(opts[lii])
	})
	lst.OnSelected = func(id widget.ListItemID) {
		entry.SetText(mapp[opts[id]])
	}

	cc := container.NewBorder(nil, nil, lst, nil, entry)
	w.SetContent(cc)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
