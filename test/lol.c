#include <sqlite3.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX_QUERY_LEN 256
typedef enum { FALSE_B, TRUE_B } bool_t;
bool_t first_json_row;

