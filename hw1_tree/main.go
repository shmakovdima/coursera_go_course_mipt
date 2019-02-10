package main

import (
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"strconv"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printSise bool) error {
	err := printDir(out, path, printSise, "")

	if err != nil {
		return err
	}

	return nil
}

func printDir(out io.Writer, path string, printSise bool, item string) error {
	files, err := ioutil.ReadDir(path)

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for index, file := range files {

		isLastElement := index == getLastElementIndex(files, !printSise)
		prefix, nestedLevelItem := getGraphics(item, isLastElement)

		if file.IsDir() {
			nestedPath := path + string(os.PathSeparator) + file.Name()
			fmt.Fprintf(out, "%v%v\n", item+prefix, file.Name())
			printDir(out, nestedPath, printSise, nestedLevelItem)
		} else if printSise {
			fmt.Fprintf(out, "%v%v\n", item+prefix, getFileNameWithSize(file))
		}

	}

	return err
}

func getLastElementIndex(files []os.FileInfo, onlyDir bool) int {
	lastIndex := len(files) - 1

	if onlyDir {
		for i := lastIndex; i >= 0; i-- {
			if files[i].IsDir() {
				return i
			}
		}
	}

	return lastIndex
}

func getGraphics(item string, isLastElement bool) (string, string) {
	var prefix string
	var nestedLevelItem string

	if !isLastElement {
		prefix = "├───"
		nestedLevelItem = item + "│\t"
	} else {
		prefix = "└───"
		nestedLevelItem = item + "\t"
	}

	return prefix, nestedLevelItem

}

func getFileNameWithSize(file os.FileInfo) string {
	if file.Size() > 0 {
		return file.Name() + " (" + strconv.FormatInt(file.Size(), 10) + "b" + ")"
	} else {
		return file.Name() + " (empty)"
	}
}