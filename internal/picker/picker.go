package picker

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/newtoallofthis123/free-launch/internal/models"
)

func PickModel(items []models.Model) (models.Model, error) {
	if _, err := exec.LookPath("fzf"); err == nil {
		return pickWithFzf(items)
	}
	return pickWithList(items)
}

func pickWithFzf(items []models.Model) (models.Model, error) {
	cmd := exec.Command("fzf", "--delimiter=\t", "--with-nth=1,2")
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return models.Model{}, err
	}

	go func() {
		for _, m := range items {
			fmt.Fprintf(stdin, "%s\t%s\n", m.ID, m.Name)
		}
		stdin.Close()
	}()

	out, err := cmd.Output()
	if err != nil {
		return models.Model{}, fmt.Errorf("fzf cancelled")
	}

	id := strings.Split(strings.TrimSpace(string(out)), "\t")[0]
	for _, m := range items {
		if m.ID == id {
			return m, nil
		}
	}
	return models.Model{}, fmt.Errorf("selected model not found")
}

func pickWithList(items []models.Model) (models.Model, error) {
	for i, m := range items {
		fmt.Fprintf(os.Stderr, "%3d) %s\t%s\n", i+1, m.ID, m.Name)
	}
	fmt.Fprintf(os.Stderr, "Select model [1-%d]: ", len(items))

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return models.Model{}, fmt.Errorf("no input")
	}

	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || n < 1 || n > len(items) {
		return models.Model{}, fmt.Errorf("invalid selection")
	}
	return items[n-1], nil
}
