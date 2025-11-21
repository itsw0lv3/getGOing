package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
  // User input pkgs
  fmt.Println("Please pkgs")
  var userPkgs string

  scanner := bufio.NewScanner(os.Stdin)
  if scanner.Scan() {
    userPkgs = scanner.Text()
  }
  userPkgArr := strings.Split(userPkgs, " ")
  for _, pkg := range userPkgArr {
    fmt.Println(pkg)
  }

  // Check for pkgs; need to check for pkg manager!
  cmd := exec.Command("sudo", "-Ss", userPkgArr[0])
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    fmt.Fprintf(os.Stderr, "failed")
  }

}
