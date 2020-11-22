# 99Movies

99Movies read in movie reviews that employees have written, and then compose "tweets" that can be shared through company account.

### Implementation

The system is designed for taking any input format, currently it is supporting to take input from file, but using a reader Interface it can support other input formats like kafka, tcp etc.
Currently, it is printing the tweets in console output but can be dumped to a file or moved to other systems like kafka, tcp etc.

### Steps to run

- To test 

`make test`

- To run

```
# build binary
make build

# commands to run
./99movies process -r {review_file_path} -m {movies_file_path}
```


