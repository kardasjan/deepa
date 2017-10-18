package cctv

import (
	"log"
	"strings"
)

const (
	intruderMessage  = "Intruder detected"
	videoLostMessage = "Video Lost"
	fullDiskMessage  = "Disk storage space low detected."
)

func isIntruder(msg string) bool {
	if strings.Contains(msg, intruderMessage) {
		log.Println("Intruder!")
		return true
	}
	return false
}

func isVideoLost(msg string) bool {
	if strings.Contains(msg, videoLostMessage) {
		log.Println("Video Lost!")
		return true
	}
	return false
}

func isFullDisk(msg string) bool {
	if strings.Contains(msg, fullDiskMessage) {
		log.Println("Full Disk!")
		return true
	}
	return false
}
