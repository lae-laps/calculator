package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "strings"
)
   
var vec = vector{
    x: 1,
    y: 1,
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

func help() {
    printBold("Help for Vector - CLI\n", 15)
    fmt.Print(" - "); printBold("set <Vₓ> <Vᵧ>", 15); fmt.Print(" - sets vector from cartesian input\n - ")
    printBold("show", 15); fmt.Print(" - shows info about the current buffer\n - ")
    printBold("arc", 15); fmt.Print(" - displays the arc of the vector\n - ")
    printBold("sector", 15); fmt.Print(" - displays the sector of the vector\n - ")
    printBold("rot <angle>", 15); fmt.Print(" - rotates the vector by the quantity specified\n - ")
    printBold("invert", 15); fmt.Print(" - inverts the vector in both cartesian axis\n - ")
    printBold("flat <x/y>", 15); fmt.Print(" - Flats the vector in the corresponding coordinate\n - ")
    printBold("elong <amount>", 15); fmt.Print(" - elongates the vector without changing the angle by the ammount provided\n - ")
    printBold("clear", 15); fmt.Print(" - clears the screen\n - ")
    printBold("exit", 15); fmt.Print(" - exits the program\n - ")
    printBold("help", 15); fmt.Print(" - displays this help message\n")
}

func interactive_ui() {
    //banner()

    vec.polar_from_cartesian() // set polar coordenates from cartesian as initialized to 100 100 

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
        case "invert":
            vec.invert()
            break
        case "flat":
            if len(split) >= 2 {
                switch split[1] {
                    case "x":
                        vec.flat('x')
                    case "y":
                        vec.flat('y')
                    case "z":
                        vec.flat('z')
                    default:
                        printUserError("Please use a valid axis - (x/y/z)\n")
                }
            } else {
                printRuntimeError("Invalid number of arguments\n")
            }
            break
        case "show":
            vec.print_disposition()
            break
        case "arc":
            vec.arc()
            break
        case "sector":
            vec.sector()
            break
        case "help":
            help()
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

