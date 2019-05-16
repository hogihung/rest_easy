# README

REST Easy is a health check application, written in Golang, which takes a file,
in JSON format, of target URLs and associated settings.  The application will
read this file on initalization and execute REST GET requests to the target
endpoints.  It will display the results to the screen, and optionally print to
a log file.

## Installation
1) Install Go, following instructions at golang.org/doc/install or, for Mac, you can just use homebrew:
```bash
brew install go &&
export GOPATH=$HOME/go; export PATH=$PATH:$GOPATH/bin &&
go get -u github.com/hogihung/rest_easy &&
cd $GOPATH/src/github.com/hogihung/rest_easy &&
go build
```

## Example Usage:
```bash
rest_easy list --targets ./targets.json --log ./rest_easy.log &&
cat rest_easy.log
```

```
➜  rest_easy>  rest_easy 
REST Easy is a command line utility, which takes a JSON formatted configuration
file and performs REST GET requests against the defined target endpoints. 

Using this app, with JSON formatted config file, one can run n number of GET requests to the
defined target endpoints and display the response to the terminal (default) and/or write the
responses to a file.

Usage:
  rest_easy [command]

Available Commands:
  adhoc       Use the 'adhoc' sub-command to run GET requests against single endpoint
  help        Help about any command
  list        Use the 'list' sub-command to list the defined endpoints to be tested
  test        Use the 'test' sub-command to run GET requests to defined endpoints

Flags:
  -h, --help         help for rest_easy
      --log string   log file (default is ./my_app.log)

Use "rest_easy [command] --help" for more information about a command.
➜  rest_easy>
```
