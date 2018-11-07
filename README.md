# GoQuiz

[![Golang](https://golang.org/doc/gopher/appenginegophercolor.jpg)](https://golang.org/doc/code.html)

GoQuiz is a program of Q & A received by a CSV file. Too, it's possible to pass a timer to answer the questions.

## Install

To run this program, you will need to:

Example:

```sh
export GOPATH=$(go env GOPATH)
go get github.com/vinixavier/goquiz
cd $GOPATH/vinixavier/goquiz
go install
```

## Test

To test this program:

```sh
cd $GOPATH/github.com/vinixavier/goquiz
go test
```

## Using

The program have following help:

```sh
Usage of goquiz:
  -filepath string
        Required - Specifies a file with content of the questions and answers for Quiz.
  -timer int
        Optional - Specifies a value to timer of the Quiz.
```

See examples of CSV files in the `data/*.csv` directory.