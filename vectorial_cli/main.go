package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

// debug conf
const show_full_error_messages = true


func handle(err error) int {
    if err != nil {
        if show_full_error_messages {
            printRuntimeError(fmt.Sprintf("error: %s\n", err))
        } else {
            printRuntimeError("error: Undefined Runtime Error\n")
        }
        return 1
    } else {
        return 0
    }
}

func banner() {
    fmt.Println(`
    ╔══════════════════╗
    ║ VECTORIAL - CLI  ║
    ╚══════════════════╝
    `)
}

func interactive_ui() {
    banner()
    for {
        in := bufio.NewReader(os.Stdin)
        line, err := in.ReadString('\n')
        handle(err)
        line = strings.Replace(line, "\n", "", -1)
    }
}

func main() {
    interactive_ui()
}

