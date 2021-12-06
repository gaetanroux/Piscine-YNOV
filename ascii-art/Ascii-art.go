package main

import (
"fmt"
"os"
"log"
"strings"
"bufio"
"./abc" // FS import.
"./def" // Output import.
)

var finalTabtest []string

// Return the asked text.
func printAscii(argument []byte, nbrLigne int, police string) []string {

    // The characters size is 8, the asked text will be returned line by line.
	for i := 0; i < 8; i++ {    
        
        // The characters will be added one by one.
		for index, value := range argument {

            // Open the police .txt and close it at the end of the function.
            f, err := os.Open(police)
            if err != nil {

                log.Fatal(err)
			}
			
            defer f.Close()

            scanner := bufio.NewScanner(f)      // Open a new scanner to read the .txt.

            numberLine := 1     // Initialize the scanner to the first line.
            
            // Break if there is a \n.
            if value == 92 && index < len(argument) - 1 {

                if argument[index + 1] == 110 {
                    
                    break
                }
            }

            // Scan the .txt and add the line needed to the array to return.
            for scanner.Scan() {

                if numberLine == int(value) - 30 + (int(value) - 32)*8 + i {

                    finalTabtest = append(finalTabtest, scanner.Text())
                }
                
                numberLine++
            }
        }
		
        finalTabtest = append(finalTabtest, "\n")
    }
    
    totalLigne := 1        // Total number of group of 8 lines according to the \n.

    // Check the number of \n in the argument.
    for index, value := range argument {

        if value == 92 && index < len(argument) - 1 {

            if argument[index + 1] == 110 {

                totalLigne++
                if nbrLigne <= totalLigne {
                    
                    newArg := []byte(argument[index + 2:])      // newArg is the continuation of the argument after the \n.
                    // Call the function with newArg.
                    printAscii(newArg, nbrLigne, police)
                    nbrLigne++
                    break
                }
            }
        }
    }

    return finalTabtest
}

func main() {
    
    arg := os.Args
    arr := []byte(arg[1])
    nbrLigne := 1       // Actual line in printing.
    police := "standard.txt"        // Default police.
    
    // Change the police.
    if len(arg) > 2 {

        police = abc.FS(arg[2])
    }

    print := printAscii(arr, nbrLigne, police)        // To make the code more clear.
    canPrint := true        // Print only once.

    // Create a new file with the return of text asked.
    for _, v := range arg {

    
        if len(v) > 9 && v[:9] == "--output=" {
            def.Output(strings.Join(print,""), v[9:])
            canPrint = false
        } 
    }

    // Print the text when needed.
    for range arg {

        if canPrint {
            for i := 0; i < len(print); i++ {

                fmt.Print(print[i])
                canPrint = false
            }
        }
    }
}
