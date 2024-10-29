// Красивый вывод списка файлов и директорий
package main

import (
	"fmt"
	"log"
	"os"

	goprettyv6 "github.com/jedib0t/go-pretty/v6/list"

	prettylist "github.com/av-belyakov/golang_structures_and_algorithms/thirdpartylibraries/go-pretty/prettylist/cmd"
)

func main() {
	fmt.Println("Example pretty list library")

	fmt.Println("The file list:")
	if err := prettyListExample(); err != nil {
		log.Fatalln(err)
	}
}

func prettyListExample() error {
	rootPath, err := prettylist.GetRootPath(prettylist.ROOT_DIR)
	if err != nil {
		return err
	}

	countFiles, countFolders, newList, err := prettylist.CreateFilesAndFoldersList(rootPath)
	if err != nil {
		return err
	}

	w := goprettyv6.NewWriter()
	w.SetOutputMirror(os.Stdout)
	w.SetStyle(goprettyv6.StyleConnectedLight)

	getList(newList, w)

	w.Render()

	fmt.Printf("Search count:\n \tfiles = %d\n\tfolders = %d\n", countFiles, countFolders)

	return nil
}

func getList(list []prettylist.Element, w goprettyv6.Writer) {
	files := []interface{}(nil)

	for _, v := range list {
		if v.Name == ".git" {
			continue
		}

		if len(v.Next) == 0 {
			//w.AppendItem(v.Name)
			files = append(files, v.Name)
		} else {
			w.AppendItem(v.Name)
			w.Indent()

			getList(v.Next, w)

			w.UnIndent()
		}
	}

	w.AppendItems(files)
}
