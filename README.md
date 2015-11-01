# tree-cat

Printing the file-name and file-content of all files in a directory (recursively).
Simple golang example.

## Usage:
### Build:
    go build tree-cat.go
### Execute:
    ./tree-cat <folder>
    ./tree-cat test/


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


## ./tree-cat test 
    test/asdf.txt
    {
    asdf
    asdf
    asdf
    }
    test/lol
    test/lol/loller.txt
    {
    asdf
    ffffl
    }
    test/lol/lol2.txt
    {
    asdf
    asdf
    asdf
    asdf
    
    }
    test/lol.c
    {
    #include <sqlite3.h>
    #include <stdio.h>
    #include <stdlib.h>
    #include <string.h>
    #define MAX_QUERY_LEN 256
    typedef enum { FALSE_B, TRUE_B } bool_t;
    bool_t first_json_row;
    
    }
