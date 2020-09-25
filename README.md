# Golang-Studies
Repo for studies propose

First Course: Learn How To Code: Google's Go (golang) Programming Language
    official-documentation: https://golang.org/  
    lang-ony-documentation: https://godoc.org/  
    links: https://www.udemy.com/course/learn-how-to-code/  
    resources: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit  
    playground: https://play.golang.org/  


# Install Golang:
1. Download proper install version from here: https://golang.org/dl/ .
2. Install it.
3. Type `go env` or `go help` to verify Golang env is installed in your machine.
4. Check for proper Golang version by command: `go version`.

# Recommendation
1. If you with mac, just `brew install go`, it's much easier.

# Set Go env
export GOROOT=/Users/<user>/Documents/gowrokspace  
(go workspace should include three directories: `bin`, `pkg` & `src` ).   
export PATH=$PATH:$GOROOT/bin   
export GOBIN=/Users/bars/Documents/gowrokspace/bin   

# Run scripts
`go run` <GoFile>.go 

# Commands
`go fmt` - format code.  
`go fmt ./...` - format all files.  
`go build Hello.go ` - test & make an executable file.   
`go install Hello.go` - make an executable file.  
`go get <repo>` - get github repo.
`go test` - test go files.
`go test -cover` - test coverage.
`go test -coverprofile c.out ` - export test profiler
`go tool cover -html=c.out` - generate html from c.out
`go tool cover -func=c.out` - generate output from c.out
`go test -bench .` - get the benchmark of your module.

# Working with Go modules
`go mod init` - init Go module.  
`go list -m all` - all the dependencies.  
`go list -m -versions <package-name>` - get all specified package versions. 