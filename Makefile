build:
	go build -o ./bin/sacct-observer ./cmd/sacct-observer

install:
	mkdir -p ~/.local/bin
	cp ./bin/sacct-observer ~/.local/bin

get-autocomplete:
	mkdir -p ~/.local/share/bash-completion/completions
	cp bash_autocomplete/sacct-observer ~/.local/share/bash-completion/completions
