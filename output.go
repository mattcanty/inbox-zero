package main

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

type result struct {
	Name        string
	Description string
	Action      string
}

type output struct {
	results []result
}

func (o *output) writeTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Description", "Action"})

	for _, o := range o.results {
		t.AppendRows([]table.Row{
			{o.Name, o.Description, o.Action},
		})
	}

	t.Render()
}
