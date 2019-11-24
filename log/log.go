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

// Print a message to stdout
func Print(arg interface{}) {
  stdOutLogger.Print(arg)
}

// Printf prints a formatted message to stdout
func Printf(args ...interface{}) {
  stdOutLogger.Printf(args[0].(string), args[1:]...)
}

// Println prints a message to stdout followed by a line break
func Println(arg interface{}) {
  stdOutLogger.Println(arg)
}

// Fatal prints a message to stderr and exits
func Fatal(arg interface{}) {
  stdErrLogger.Fatal(arg)
}

// Fatalf prints a formatted message to stderr and exits
func Fatalf(args ...interface{}) {
  stdErrLogger.Fatalf(args[0].(string), args[1:]...)
}

// Fatalln prints a message to stderr followed by a line break and exits
func Fatalln(args ...interface{}) {
  stdErrLogger.Fatalln()
}

// Panic with a message printed to stderr
func Panic(arg interface{}) {
  stdErrLogger.Panic(arg)
}

// Panicf panics with a formatted message printed to stderr
func Panicf(args ...interface{}) {
  stdErrLogger.Panicf(args[0].(string), args...)
}

// Panicln panics with a message followed by a line break printed to stderr
func Panicln(arg interface{}) {
  stdErrLogger.Panicln(arg)
}

// Return a formatted error message
func Errorf(args ...interface{}) error {
  return errors.Errorf(args[0].(string), args[1:]...)
}
