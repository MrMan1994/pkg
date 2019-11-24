package log

import (
  "flag"
  "github.com/pkg/errors"
  "log"
  "os"
)

var (
  stdOutLogger = log.New(os.Stdout, "", 0)
  stdErrLogger = log.New(os.Stderr, "", 0)
)

func init() {
  flag.Parse()
}

func Print(arg interface{}) {
  stdOutLogger.Print(arg)
}

func Printf(args ...interface{}) {
  stdOutLogger.Printf(args[0].(string), args[1:]...)
}

func Println(arg interface{}) {
  stdOutLogger.Println(arg)
}

func Fatal(arg interface{}) {
  stdErrLogger.Fatal(arg)
}

func Fatalf(args ...interface{}) {
  stdErrLogger.Fatalf(args[0].(string), args[1:]...)
}

func Fatalln(args ...interface{}) {
  stdErrLogger.Fatalln()
}

func Panic(arg interface{}) {
  stdErrLogger.Panic(arg)
}

func Panicf(args ...interface{}) {
  stdErrLogger.Panicf(args[0].(string), args...)
}

func Panicln(arg interface{}) {
  stdErrLogger.Panicln(arg)
}

func Errorf(args ...interface{}) error {
  return errors.Errorf(args[0].(string), args[1:]...)
}
