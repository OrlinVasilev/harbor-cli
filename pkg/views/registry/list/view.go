package list

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/utils"
	"github.com/goharbor/harbor-cli/pkg/views"
)

type model struct {
	table table.Model
}

var columns = []table.Column{
	{Title: "ID", Width: 6},
	{Title: "Name", Width: 12},
	{Title: "Status", Width: 12},
	{Title: "Endpoint URL", Width: 26},
	{Title: "Provider", Width: 12},
	{Title: "Creation Time", Width: 24},
	// {Title: "Verify Remote Cert", Width: 12},
	// {Title: "Description", Width: 12},
}

func (m model) Init() tea.Cmd {
	return tea.Quit
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return views.BaseStyle.Render(m.table.View()) + "\n"
}

func ListRegistry(registry []*models.Registry) {
	var rows []table.Row
	for _, regis := range registry {
		createdTime, _ := utils.FormatCreatedTime(regis.CreationTime.String())
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", regis.ID),
			regis.Name,
			regis.Status,
			regis.URL,
			regis.Type,
			createdTime,
			// regis.Description,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(rows)),
	)

	// Set the styles for the table
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(s.Cell.GetForeground()).
		Background(s.Cell.GetBackground()).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
