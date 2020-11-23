# 99Movies

99Movies read in movie reviews that employees have written, and then compose "tweets" that can be shared through company account.

### Implementation

Keeping in mind the single responsibility and open/close principles of System design, the system is designed for taking any input format, 
currently it is supporting to take input from file, but using a reader Interface it can support other input formats like kafka, tcp etc.
Just need to Implement ReaderInterface and override the methods to support other input source and 
similarly need to implement WriterInterface to support writing to various other output destination.

Channels and GoRoutines are used to read reviews and write tweets in real time rather than storing in memory.
The application first reads movies files and stores movies and its release date in local cache so that look up happens in O(1) time.
Ratings are calculated based on weighted scores.

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


