package main

import (
    "fmt"
    "io"
    "os"
    "io/ioutil"
    "sort"
    "path"
    "strconv"
    // "path/filepath"
    // "path/filepath"
    // "strings"
)

func listFolder(out io.Writer, folder_path string, prevOffset string, printFiles bool) error{
    folders, _ := ioutil.ReadDir(folder_path)
    
    sort.Slice(folders, func(i, j int) bool {
        return folders[i].Name() < folders[j].Name()
    })
    totalFoldersCount := 0
    totalFilesCount := 0
    for _, f := range folders{
        if f.IsDir(){
            totalFoldersCount += 1
        }else{
            totalFilesCount += 1
        }
    }
    currentFolderCount := 0
    currentFileCount := 0
    if !printFiles{
        currentFileCount = totalFilesCount
    }
    for _, f := range folders{
        offset := ""
        nextOffset := prevOffset
        if f.IsDir(){
            currentFolderCount += 1
            if (currentFolderCount == totalFoldersCount && currentFileCount == totalFilesCount){
                offset = prevOffset + "└───"
                nextOffset += "    "
            }else{
                offset = prevOffset + "├───"
                nextOffset += "│   "
            }
            fmt.Fprintf(out, offset+f.Name()+"\n")
            
            listFolder(out, path.Join(folder_path,f.Name()), nextOffset, printFiles)
        }else{
            if printFiles{
                currentFileCount += 1
                if(currentFolderCount == totalFoldersCount && currentFileCount == totalFilesCount){
                    offset = prevOffset + "└───"
                }else{
                    offset = prevOffset + "├───"
                }
                sizePart := ""
                if f.Size() == 0{
                    sizePart = " (empty)"
                }else{
                    sizePart = " ("+ strconv.FormatInt(f.Size(),10) + "b)"
                }
                fmt.Fprintf(out, offset+f.Name() + sizePart+"\n")
            }
        }
    }
    return nil
}

func dirTree(out io.Writer, startPath string, printFiles bool) error{
    listFolder(out, startPath, "", printFiles)
    return nil
}

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

