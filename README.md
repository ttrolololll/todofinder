# TODO Finder

Finds TODOs from given directory/file path.

## Installation

`go install`

## Usage

```
todofinder <path>

// Or to run without installation:
// go run main <path>
```

Eg. `todofinder ~/myproject`

### Args

path: (optional) defaults to current working directory

### Enhancements / Known Caveats

1. Does not consider code context (comment vs. string). Eg. `[]byte("// TODO: must play dota")` will be captured as a valid result
1. Does not allow exclusion of files/directories via CLI flag/config. file
1. Should perhaps process directories/files with bounded goroutines
