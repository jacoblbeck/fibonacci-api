# fibonacci-api

To run the server clone this repo and run `make up` from the base of the file. This will create the Go binary and run the docker compose commands to run the server and postgres database.

The endpoints available are:
```/health
   /api/v1/fibonacci/current
   /api/v1/fibonacci/next  
   /api/v1/fibonacci/previous
   /api/v1/fibonacci/reset```

The health endpoint returns the status of the server, the current returns the current value in the fibonacci sequence, next returns the next value, previous returns the previous value, and reset returns to the beginning of the sequence.  