# Assignment 1

This is a pretty minimal solution to the assignment, nothing too special going on. Errors are reported over http or by logging them to stdout.

The server is very picky about how to format the url path (all lower case, dates must be yyyy-mm-dd), if anything is wrong, you get a 404, which may not be entirely appropriate. Other error codes are more conventional.
