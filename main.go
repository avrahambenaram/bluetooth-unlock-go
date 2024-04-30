package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
  runned := false
  for ;; {
    isConnected := isBluetoothDeviceConnected()
    if isConnected && !runned {
      runned = true
      typePassword()
    }
    if !isConnected {
      runned = false
    }
    time.Sleep(time.Duration(500) * time.Millisecond)
  }
}

func isBluetoothDeviceConnected() bool {
  deviceMatch := os.Getenv("BLUETOOTH_DEVICE")
  devicesCmd := exec.Command("bluetoothctl", "devices", "Connected")
  devices, devicesErr := devicesCmd.Output()
  if devicesErr != nil {
    fmt.Println(devicesErr)
    return false
  }
  devicesSep := strings.Split(string(devices), "\n")
  var isDeviceConnected bool
  for _, device := range(devicesSep) {
    isDeviceConnected = strings.Contains(device, deviceMatch)
    if isDeviceConnected {
      break
    }
  }
  return isDeviceConnected
}

func typePassword() {
  password := os.Getenv("PASSWORD")
  typingPassword := fmt.Sprintf("%s\n", password)
  typingCmd := exec.Command("wtype", typingPassword)
  typingCmd.Run()
}
