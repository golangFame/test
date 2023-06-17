#!/bin/sh



export GOBIN="/workspaces/bin"

go install github.com/cosmtrek/air@latest

export PATH="$PATH:$GOBIN"

echo 'export PATH="$PATH:$GOBIN"' >> ~/.zshrc
echo 'export PATH="$PATH:$GOBIN"' >> ~/.bashrc
source ~/.zshrc
source ~/.bashrc

sudo apt update -y
sudo apt install tmux -y



# export GOPATH="/workspaces/gopath"
# export GOBIN="$GOPATH/bin"