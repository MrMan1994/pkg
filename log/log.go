package log

import(
  "flag"
  "log"
  "os"
  "github.com/pkg/errors"
)

var (
  stdOutLogger = log.New(os.Stdout, "", 0)
  stdErrLogger = log.New(os.Stderr, "", 0)
)

func init() {
  flag.Parse()
}

func Print() {

}

func Printf() {

}

func Println() {

}

func Fatal() {

}

func Fatalf() {

}

func Fatalln() {

}

func Panic() {

}

func Panicf() {

}

func Panicln() {

}

func Errorf() error {
  return
}
