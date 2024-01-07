// Log package for free-form and the log/slog package for structured output
package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
    // Invoking function like Println from the log package uses standard logger,
    // which is already pre-configured for reasonable logging output to os.Stderr.
    // Additional methods like *Fatatl or Panic* will exit the program after
    // logging.
    log.Println("standard logger")

    // Loggers can be configured with flags to set their output format.
    // By default, standard logger has the log.Ldate and log.Ltime flags
    // set, and these are collected in log.LstdFlags. Can change its flags
    // to emit time with microsecond accuracy, for example
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    log.Println("with micro")

    // It also supports emitting the file name and line from which the log
    // function is called.
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("with file/line")

    // May be useful to create a custom logger and pass it around. When creating
    // a new logger, we can set a prefix to distinguish its output from other loggers.
    mylog := log.New(os.Stdout, "my:", log.LstdFlags)
    mylog.Println("from my log")

    // Can set the prefix on existing loggers (including the standard one)
    // with the SetPrefix method.
    mylog.SetPrefix("ohmy:")
    mylog.Println("from my log")

    // Loggers can have custom output targets; any io.Writer works.
    var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)

    // This call writes the log output into buf.
    buflog.Println("hello")

    // This will actually show it on standard output.
    fmt.Print("from buflog:", buf.String())

    // The slog package provides structured log output. For example, logging in
    // JSON format is straightforward.
    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)
    myslog.Info("hi there")

    // In addition to the msg, slog output can contain an arbitrary number of
    // key=value pairs.
    myslog.Info("hello again", "key", "val", "age", 25)
}
