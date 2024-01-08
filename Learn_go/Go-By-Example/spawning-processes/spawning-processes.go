// Sometimes our Go programs need to spawn other, non-go processes.
package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
    // Start with a simple command that takes no arguments or input and just
    // prints something to stdout. The exec.Command helper creates an obj to
    // represent this external process.
    dateCmd := exec.Command("date")

    // The output method runs the command, waits for it to finish and collects
    // its standard output. If there were no errors, dateOut will hold bytes with
    // the date info.
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // Output and other methods of Command will return *exec.Error if there was
    // a problem executing the command(e.g. wrong path), and *Exec.ExitError if
    // the command ran but exited with a non-zero return code.
    _, err = exec.Command("date", "-x").Output() 
    if err != nil {
        switch e := err.(type) {
        case *exec.Error:
            fmt.Println("failed executing:", err)
        case *exec.ExitError:
            fmt.Println("command exit rc =", e.ExitCode())
        default:
            panic(err)
        }
    }

    // We'll look at a slightly more involved case where we pipe data to the
    // external process on its stdin and collect the result from its stdout.
    grepCmd := exec.Command("grep", "hello")

    // We explicitly grab input/output pipes, start the process, write some input
    // to it, read the resulting output, and finally wait for the process to exit
    grepIn, err := grepCmd.StdinPipe()
    if err != nil {
        panic(err)
    }
    grepOut, err := grepCmd.StdoutPipe()
    if err != nil {
        panic(err)
    }
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, err := io.ReadAll(grepOut)
    if err != nil {
        panic(err)
    }
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // Note that when spawning commands we need to provide an explicitly
    // delineated command and argument array, vs being able to just pass in one
    // command-line string. If you want to spawn a full command with a string,
    // you can use bash's -c option:
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))

}
