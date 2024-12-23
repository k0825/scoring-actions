package main

import (
    "bytes"
    "encoding/csv"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    file, err := os.Open("./problems/1.two-sum/test.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading CSV:", err)
        return
    }

    for _, record := range records {
        inputData := strings.ReplaceAll(record[0], "\"", "")
        expectedOutput := strings.ReplaceAll(record[1], "\"", "")

        cmd := exec.Command("npx", "ts-node", "./problems/1.two-sum/main.ts")
        cmd.Stdin = bytes.NewBufferString(inputData)
        var out bytes.Buffer
        cmd.Stdout = &out

        err := cmd.Run()
        if err != nil {
            fmt.Println("Error running command:", err)
            return
        }

        actualOutput := strings.TrimSpace(out.String())
        if actualOutput == strings.TrimSpace(expectedOutput) {
            fmt.Println("Success")
        } else {
            fmt.Println("Failure")
        }
    }
}