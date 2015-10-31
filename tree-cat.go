package main

// go run tree-cat.go <folder>

import (
    "fmt"
    "flag"
    "os"
    "io"
)

func checkerror(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func cat_file( filename string) {

    //fmt.Printf("======(File %s Begin)======\n", filename)
    fmt.Printf("{\n")

    // Print file content
    fh, err := os.Open(filename)
    checkerror(err)
    _, err = io.Copy(os.Stdout, fh)
    checkerror(err)

    //fmt.Printf("======(File %s End)======\n", filename)
    fmt.Printf("}\n")
}

func tree_cat_recursive( files []os.FileInfo, dir *os.File ) {

    for _, fileInfo := range files {

        if fileInfo.IsDir() {

            sourcedir := dir.Name() + "/" + fileInfo.Name()
            fmt.Println(sourcedir)

            new_dir, err := os.Open(sourcedir)
            checkerror(err)
            defer new_dir.Close()

            // Grab the files (list)
            new_files, err := new_dir.Readdir(0) 
            checkerror(err)

            //() Recursive
            tree_cat_recursive( new_files, new_dir )

        } else {

            filename := dir.Name() + "/" + fileInfo.Name()
            //fmt.Println( filename , "\t", fileInfo.Size(), "\t", fileInfo.Mode(), fileInfo.ModTime() )
            fmt.Println( filename )

            //() cat the filecontent
            cat_file( filename )
        }
    }
}

func main() {

    // Get the arguments from command line
    flag.Parse() 
    sourcedir := flag.Arg(0)
    if sourcedir == "" {
        sourcedir = "."
    }

    // Directory to tree
    dir, err := os.Open(sourcedir)
    checkerror(err)
    defer dir.Close()

    // Grab list of files in dir
    files, err := dir.Readdir(0) 
    checkerror(err)

    //()
    tree_cat_recursive( files, dir )
}

