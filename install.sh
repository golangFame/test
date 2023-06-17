#!/bin/bash

zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer) &
source /home/codespace/.gvm/scripts/gvm
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.17.13
gvm use go1.17.13
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.20
gvm use go1.20

gvm use go1.20 --default

# sudo apt-get install curl git mercurial make binutils bison gcc build-essential

# go install golang.org/dl/go1.14@latest & go1.14 download ;
# bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer);


curl -sLk https://raw.githubusercontent.com/kevincobain2000/gobrew/master/git.io.sh | sh

go install github.com/kevincobain2000/gobrew/cmd/gobrew@latest


export PATH="$HOME/.gobrew/current/bin:$HOME/.gobrew/bin:$PATH"
# export GOROOT="$HOME/.gobrew/current/go"
