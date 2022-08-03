package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "strings"
)
   
var vec = vector{
    x:      100,
    y:      100,
}

// CLI conf
const show_changes_in_vectors = true
const show_full_error_messages = true
const print_ansi_escapes = true

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
    print(`
    ╔══════════════════╗
    ║ VECTORIAL - CLI  ║
    ╚══════════════════╝
    
    `, 15)
}

func interactive_ui() {
    //banner()

    vec.polar_from_cart() // set polar coordenates from cartesian as initialized to 100 100 

    for {
        print("v > ", 15)
        in := bufio.NewReader(os.Stdin)
        line, err := in.ReadString('\n')
        handle(err)
        line = strings.Replace(line, "\n", "", -1)
        parse(line)
    }
}

func parse(expression string) {
    split := strings.Split(expression, " ")
    switch split[0] {
        case "set":
            if len(split) >= 3 {
                x, err := strconv.ParseFloat(split[1], 32)
                handle(err)
                y, err := strconv.ParseFloat(split[2], 32)
                handle(err)
                vec.set_cart(x, y)
            } else {
                printRuntimeError("Invalid number of arguments\n")
            }
            break
        case "rot":
            if len(split) >= 2 {
                difference, err := strconv.ParseFloat(split[1], 32)
                handle(err)
                vec.rot(difference)
            } else {
                printRuntimeError("Invalid number of arguments\n")
            }
            break
        case "elong":
            if len(split) >= 2 {
                ammount, err := strconv.ParseFloat(split[1], 32)
                handle(err)
                vec.elong(ammount)
            } else {
                printRuntimeError("Invalid number of arguments\n")
            }
        case "show":
            vec.print_disposition()
            break
        case "clear":
            clearScreen()
            break
        case "exit":
            os.Exit(0)
        default:
            printRuntimeError("Invalid Expression\n")
    }
}

func main() {
    interactive_ui()
}

