package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Model struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const cacheTTL = 6 * time.Hour

func configDir() string {
	base, _ := os.UserConfigDir()
	return filepath.Join(base, "free-launch")
}

func EnsureModels() error {
	dir := configDir()
	cpPath := filepath.Join(dir, "checkpoint.txt")

	data, err := os.ReadFile(cpPath)
	if err == nil {
		ts, err := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64)
		if err == nil && time.Since(time.Unix(ts, 0)) < cacheTTL {
			return nil
		}
	}

	return fetchAndCache()
}

func fetchAndCache() error {
	resp, err := http.Get("https://openrouter.ai/api/v1/models")
	if err != nil {
		return fmt.Errorf("fetching models: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}

	var result struct {
		Data []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Pricing struct {
				Prompt     string `json:"prompt"`
				Completion string `json:"completion"`
			} `json:"pricing"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("parsing response: %w", err)
	}

	var free []Model
	for _, m := range result.Data {
		if m.Pricing.Prompt == "0" && m.Pricing.Completion == "0" {
			free = append(free, Model{ID: m.ID, Name: m.Name})
		}
	}

	dir := configDir()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	out, err := json.Marshal(free)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dir, "data.json"), out, 0o644); err != nil {
		return err
	}

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	return os.WriteFile(filepath.Join(dir, "checkpoint.txt"), []byte(ts), 0o644)
}

func LoadModels() ([]Model, error) {
	data, err := os.ReadFile(filepath.Join(configDir(), "data.json"))
	if err != nil {
		return nil, err
	}
	var models []Model
	return models, json.Unmarshal(data, &models)
}

func FindModel(query string, models []Model) (Model, error) {
	// Exact match
	for _, m := range models {
		if m.ID == query {
			return m, nil
		}
	}

	// Prefix match
	var matches []Model
	for _, m := range models {
		if strings.HasPrefix(m.ID, query) {
			matches = append(matches, m)
		}
	}

	switch len(matches) {
	case 0:
		return Model{}, fmt.Errorf("no model matching %q", query)
	case 1:
		return matches[0], nil
	default:
		ids := make([]string, len(matches))
		for i, m := range matches {
			ids[i] = m.ID
		}
		return Model{}, fmt.Errorf("ambiguous query %q, matches: %s", query, strings.Join(ids, ", "))
	}
}
