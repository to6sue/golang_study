
go env


#terminal
nano ~/.bash_profile

eexport GOROOT=/usr/local/go
export GOPATH="/Users/dev/Documents/workspace-vscode/go_study"
export GOBIN=$GOPATH/bin

export PATH="$GOPATH:$GOBIN:$PATH:$GOROOT/bin"

source ~/.bash_profile