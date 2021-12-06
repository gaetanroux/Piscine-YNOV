package def

import (
	"fmt"
	"io"
    "os"
)

// Output create a file named after the argument.
func Output(strPerso string, fileName string) {

    if fileName != "" {
        
        createFile(fileName)
        deleteFile(fileName)
        createFile(fileName)
        writeFile(strPerso, fileName) 
        readFile(fileName)
    }
}

func createFile(fileName string) {
    // check if file exists.
    var _, err = os.Stat(fileName)

    // create file if not exists.
    if os.IsNotExist(err) {
        var file, err = os.Create(fileName)
        if isError(err) {
            return
        }
        defer file.Close()
    }
}

func writeFile(test string, fileName string) {
    // Open file using READ & WRITE permission.
    var file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    // Write some text according to the parameters.
    _, err = file.WriteString(test)
    if isError(err) {
        return
    }

    // Save file changes.
    err = file.Sync()
    if isError(err) {
        return
    }
}

func readFile(fileName string) {
    // Open file for reading.
    var file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    // Read file, line by line.
    var text = make([]byte, 1024)
    for {
        _, err = file.Read(text)

        // Break if finally arrived at end of file.
        if err == io.EOF {
            break
        }

        // Break if error occured.
        if err != nil && err != io.EOF {
            isError(err)
            break
        }
    }
}

// isError check if there is an error.
func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}

func deleteFile(fileName string) {
    // Delete file.
    var err = os.Remove(fileName)
    if isError(err) {
        return
    }
}
