package main

import (
	"fmt"
	"os"
	"os/exec"
)

func launchClaude(model string) error {
	if os.Getenv("OPENROUTER_API_KEY") == "" {
		return fmt.Errorf("OPENROUTER_API_KEY is not set")
	}

	cmd := exec.Command("claude", "--model", model)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Fprintf(os.Stderr, "Launching Claude Code with %s...\n", model)
	cmd.Env = append(os.Environ(),
		"ANTHROPIC_BASE_URL=https://openrouter.ai/api",
		"ANTHROPIC_API_KEY=",                                    // empty — suppresses key detection prompt
		"ANTHROPIC_AUTH_TOKEN="+os.Getenv("OPENROUTER_API_KEY"), // actual auth
		"ANTHROPIC_DEFAULT_OPUS_MODEL="+model,
		"ANTHROPIC_DEFAULT_SONNET_MODEL="+model,
		"ANTHROPIC_DEFAULT_HAIKU_MODEL="+model,
		"CLAUDE_CODE_SUBAGENT_MODEL="+model,
	)
	return cmd.Run()
}
