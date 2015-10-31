# tree-cat

Printing the file-contents of all files in a directory (recursive).
Simple golang example. 

## Usage:
### Build:
    go build tree-cat.go
### Execute:
    ./tree-cat.go <folder>
    ./tree-cat.go test/

## Similar to:

### find
    find <folder> -type f -print -exec cat {} \;
    find test/ -type f -print -exec cat {} \;

### find , xargs, echo, cat
    find <folder> -type f -print0 | xargs -0 -I % sh -c 'echo %; cat %'
    find test/ -type f -print0 | xargs -0 -I % sh -c 'echo %; cat %'

### grep
    grep -r '.*' <folder>
    grep -r '.*' test/

### cat (bad)
    cat test/* test/*/*

### tree (Can't show file-content)
    tree <folder>
    tree test/


