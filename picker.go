package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func pickModel(models []Model) (Model, error) {
	if _, err := exec.LookPath("fzf"); err == nil {
		return pickWithFzf(models)
	}
	return pickWithList(models)
}

func pickWithFzf(models []Model) (Model, error) {
	cmd := exec.Command("fzf", "--delimiter=\t", "--with-nth=1,2")
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return Model{}, err
	}

	go func() {
		for _, m := range models {
			fmt.Fprintf(stdin, "%s\t%s\n", m.ID, m.Name)
		}
		stdin.Close()
	}()

	out, err := cmd.Output()
	if err != nil {
		return Model{}, fmt.Errorf("fzf cancelled")
	}

	id := strings.Split(strings.TrimSpace(string(out)), "\t")[0]
	for _, m := range models {
		if m.ID == id {
			return m, nil
		}
	}
	return Model{}, fmt.Errorf("selected model not found")
}

func pickWithList(models []Model) (Model, error) {
	for i, m := range models {
		fmt.Fprintf(os.Stderr, "%3d) %s\t%s\n", i+1, m.ID, m.Name)
	}
	fmt.Fprintf(os.Stderr, "Select model [1-%d]: ", len(models))

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return Model{}, fmt.Errorf("no input")
	}

	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || n < 1 || n > len(models) {
		return Model{}, fmt.Errorf("invalid selection")
	}
	return models[n-1], nil
}
