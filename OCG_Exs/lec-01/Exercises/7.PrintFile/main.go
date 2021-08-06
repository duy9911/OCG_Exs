package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "duy.txt"
	data := "Hello world"

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644) //-Append&Write

	if err == nil {
		fmt.Println("\n Opened")
		_, err1 := fmt.Fprintln(file, "\n"+data) //write to file
		if err1 == nil {
			fmt.Println("\n Wrote")
		} else {
			fmt.Println("\n Write error")
		}
		file.Close()
	} else {

		fileNew, err2 := os.Create(filename)
		if err2 == nil {
			fmt.Println("\n Created")
			_, err3 := fileNew.WriteString(data)
			if err3 == nil {
				fmt.Println("\n Wrote to new file")
				fileNew.Close()
				return
			} else {
				fmt.Println("\n Can not write to new file")
				fileNew.Close()
			}
		} else {
			fmt.Println("\n Can not create file")
			fileNew.Close()
		}

	}
}
