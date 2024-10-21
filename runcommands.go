package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	cp "github.com/otiai10/copy"
)

var destination = "C:\\Users\\kenne\\codebase"

func CopyTempate(projectName, sourceTemplatePath string) string {
	projectFolder := destination + "\\" + projectName
	log.Printf("creating project in %s\n", projectFolder)
	if err := cp.Copy(sourceTemplatePath, projectFolder); err != nil {
		log.Fatalf("failed to copy template %s", err.Error())
	}
	return projectFolder
}

func Install(projectFolder string) {
	log.Printf("running 'yarn install'")
	cmd := exec.Command("yarn", "install")
	cmd.Dir = projectFolder
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("error running command %s", err.Error())
	}
	fmt.Println(string(out))
}

func OpenCode(projectFolder string) {
	log.Printf("opening in code")
	cmd := exec.Command("code", ".")
	cmd.Dir = projectFolder
	_, err := cmd.Output()
	if err != nil {
		log.Fatalf("error opening vs code %s", err.Error())
	}
}

func RunCommands(projectName, sourceTemplatePath string) {
	projectFolder := CopyTempate(projectName, sourceTemplatePath)
	time.Sleep(time.Second * 2)
	Install(projectFolder)
	OpenCode(projectFolder)
}
