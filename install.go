// https://golang.org/pkg/os/exec/

package main

import (
"os/exec"
"fmt"
)

func removeDocker() (string, error) {
  removeCommand="sudo apt-get remove docker docker-engine docker.io containerd runc"
  cmd := exec.Command("/bin/sh", "-c", removeCommand)
  err := cmd.Start()
  if err != nil{
    return "Docker removed", nil
  }
  return "Docker removal failed, maybe not installed: " + err, 1
}

func installDocker() (string, error) {
  defer func(){
    if r:= recover(); r!=nil {
	err = r.(error)
    }
    return "Docker installation failed"+err, 1 
  }
  installCommand=`sudo apt-get update && 
	sudo apt-get install \
	  apt-transport-https ca-certificates \
	  curl gnupg-agent software-properties-common &&
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - &&
	sudo add-apt-repository \
	   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
	   $(lsb_release -cs) \
	   stable" &&
	sudo apt-get update &&
	sudo apt-get install docker-ce docker-ce-cli containerd.io`
  cmd := exec.Command("/bin/sh", "-c", installCommand)
  err := cmd.Start()
  if err != nil{
    return "Docker installed", nil
  }
  return "Docker installation error: " + err, 1
}


func main() {
result, err = removeDocker()
fmt.print(result)
}
