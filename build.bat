@echo off
set CGO_CXXFLAGS="-IC:/Users/u/go/pkg/mod/github.com/del-xiong/miniblink@v0.0.0-20220609045710-8ec016458de6/"
go build -ldflags="-H windowsgui" -o build/basic.exe main.go