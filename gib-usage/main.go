package main

import (
	"fmt"
)

const (
	viewsPerMonth   = 500 * 1000 * 1000
	viewsPerDay     = viewsPerMonth / 30
	viewsPerHour    = viewsPerDay / 24
	viewsPerMinute  = viewsPerHour / 60
	viewsPerSecond  = viewsPerMinute / 60
	bytesPerDay     = 50 * 1024 * 1024 * 1024 * 1024
	bytesPerHour    = bytesPerDay / 24
	bytesPerMinutes = bytesPerHour / 60
	bytesPerSecond  = bytesPerMinutes / 60
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func main() {
	fmt.Println("==============================================")
	fmt.Println("Bytes per day:", ByteSize(bytesPerDay))
	fmt.Println("Views per month:", viewsPerMonth)
	fmt.Println("==============================================")

	// calculating
	fmt.Println("Views per second:", viewsPerSecond)
	fmt.Println("Data per second:", ByteSize(bytesPerSecond))
	fmt.Println("Data per view:", ByteSize(bytesPerDay/viewsPerDay))
}
