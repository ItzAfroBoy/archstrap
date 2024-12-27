package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type SCHEMA struct {
	Pacman []string
	Git    []string
	Yay    []string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readJSON(s *SCHEMA) {
	data, err := os.ReadFile("install.json")
	check(err)
	check(json.Unmarshal(data, s))
}

func installYay(p *tea.Program) {
	if _, err := exec.Command("yay", "-V").Output(); err != nil {
		os.Chdir("Documents")
		os.Chdir("yay")
		check(exec.Command("makepkg", "-si", "--noconfirm").Run())
		p.Send(yayInstallMsg("✓ Yay installed"))
	} else {
		p.Send(yayInstallMsg("✓ Yay already installed"))
	}
}

func pacman(p *tea.Program, pkgs []string) {
	for _, pkg := range pkgs {
		check(exec.Command("sudo", "pacman", "-S", pkg, "--noconfirm").Run())
	}
	p.Send(pacmanPkgMsg(fmt.Sprintf("✓ %d Packages installed", len(pkgs))))
}

func git(p *tea.Program, repos []string) {
	os.Chdir("Documents")
	for _, repo := range repos {
		check(exec.Command("git", "clone", repo).Run())
	}
	p.Send(gitRepoMsg(fmt.Sprintf("✓ %d repos cloned", len(repos))))
}

func yay(p *tea.Program, pkgs []string) {
	for _, pkg := range pkgs {
		check(exec.Command("yay", "-S", pkg, "--noconfirm").Run())
	}
	p.Send(yayPkgMsg(fmt.Sprintf("✓ %d Packages installed", len(pkgs))))
}

func main() {
	skipPacman := flag.Bool("skip-pacman", false, "Skip pacman package installation")
	skipGit := flag.Bool("skip-git", false, "Skip git repo cloning")
	skipYay := flag.Bool("skip-yay", false, "Skip yay package installation")
	flag.Parse()

	var packages SCHEMA
	p := tea.NewProgram(initialModel())

	readJSON(&packages)
	go func() {
		installYay(p)

		if *skipGit {
			p.Send(gitRepoMsg("✓ Git repo cloning skipped"))
		} else {
			go git(p, packages.Git)
		}

		if *skipYay {
			p.Send(yayPkgMsg("✓ Yay package installation skipped"))
		} else {
			go yay(p, packages.Yay)
		}

		if *skipPacman {
			p.Send(pacmanPkgMsg("✓ Pacman package installation skipped"))
		} else {
			pacman(p, packages.Pacman)
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
