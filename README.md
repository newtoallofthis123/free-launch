# free-launch

Run [Claude Code](https://docs.anthropic.com/en/docs/claude-code) with free models from [OpenRouter](https://openrouter.ai/).

`free-launch` fetches the current list of free models from OpenRouter, lets you pick one (with fzf or a numbered list), and launches `claude` with the right environment variables so it talks to OpenRouter instead of the Anthropic API.

## Install

```
go install github.com/newtoallofthis123/free-launch@latest
```

Or build from source:

```
git clone https://github.com/newtoallofthis123/free-launch.git
cd free-launch
go build -o free-launch .
```

## Prerequisites

- [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI (`claude`) installed and on your PATH
- An [OpenRouter](https://openrouter.ai/) API key (free tier works)
- (Optional) [fzf](https://github.com/junegunn/fzf) for fuzzy model selection

## Usage

Set your OpenRouter API key:

```
export OPENROUTER_API_KEY=sk-or-...
```

Pick a model interactively:

```
free-launch claude
```

Or specify a model directly (exact ID or prefix):

```
free-launch claude google/gemma-3
```

The model list is cached locally for 6 hours. Only models with `$0` prompt and completion pricing are shown.

## How it works

1. Fetches available models from the OpenRouter API
2. Filters to free-tier models only
3. Lets you pick via fzf (if installed) or a numbered list
4. Launches `claude` with environment variables pointing at OpenRouter

## License

[MIT](LICENSE)
