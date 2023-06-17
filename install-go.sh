#!/bin/sh

go install golang.org/dl/go1.14@latest & go1.14 download ;
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer);
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer);
source /home/codespace/.gvm/scripts/gvm
