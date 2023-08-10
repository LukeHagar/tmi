/*
Copyright © 2023 Luke Hagar lukeslakemail@gmail.com
*/

package root

import (
	_ "embed"
	"fmt"
	"log"
	"lukehagar/tmi/root/tui"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

//go:embed long.md
var LongMD string

//go:embed example.md
var ExampleMD string

func RenderMarkdown(in string) string {
	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)

	out, err := renderer.Render(in)
	if err != nil {
		log.Panic(err)
	}
	return out
}

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "tmi",
		Short:   "A Terminal Music Interface",
		Long:    RenderMarkdown(LongMD),
		Example: RenderMarkdown(ExampleMD),
		RunE: func(cmd *cobra.Command, args []string) error {

			rand.Seed(time.Now().UTC().UnixNano())

			if _, err := tea.NewProgram(tui.NewModel()).Run(); err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}

			return nil
		},
	}
	return rootCmd
}
