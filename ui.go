package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type model struct {
	strs     []string
	states   []bool
	spinners []spinner.Model
}

type pacmanPkgMsg string
type gitRepoMsg string
type yayPkgMsg string
type yayInstallMsg string

func initialModel() model {
	spinners := make([]spinner.Model, 4)
	states := []bool{false, false, false, false}
	strs := []string{"Installing yay...", "Installing pacman packages...", "Installing yay packages...", "Cloning repos..."}

	for i := range spinners {
		spinners[i] = spinner.New()
		spinners[i].Spinner = spinner.Dot
		spinners[i].Style = lg.NewStyle().Foreground(lg.Color("205"))
	}
	return model{spinners: spinners, states: states, strs: strs}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinners[0].Tick, m.spinners[1].Tick, m.spinners[2].Tick, m.spinners[3].Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}
	case yayInstallMsg:
		m.strs[0] = string(msg)
		m.states[0] = true
	case pacmanPkgMsg:
		m.strs[1] = string(msg)
		m.states[1] = true
	case yayPkgMsg:
		m.strs[2] = string(msg)
		m.states[2] = true
	case gitRepoMsg:
		m.strs[3] = string(msg)
		m.states[3] = true
	}

	if m.states[0] && m.states[1] && m.states[2] && m.states[3] {
		return m, tea.Quit
	}

	m.spinners[0], cmd = m.spinners[0].Update(msg)
	cmds = append(cmds, cmd)
	m.spinners[1], cmd = m.spinners[1].Update(msg)
	cmds = append(cmds, cmd)
	m.spinners[2], cmd = m.spinners[2].Update(msg)
	cmds = append(cmds, cmd)
	m.spinners[3], cmd = m.spinners[3].Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var str string
	if !m.states[0] {
		str = m.spinners[0].View() + m.strs[0] + "\n"
	} else {
		str = m.strs[0] + "\n"
	}
	if !m.states[1] {
		str += m.spinners[1].View() + m.strs[1] + "\n"
	} else {
		str += m.strs[1] + "\n"
	}
	if !m.states[2] {
		str += m.spinners[2].View() + m.strs[2] + "\n"
	} else {
		str += m.strs[2] + "\n"
	}
	if !m.states[3] {
		str += m.spinners[3].View() + m.strs[3] + "\n"
	} else {
		str += m.strs[3] + "\n"
	}
	return str
}
