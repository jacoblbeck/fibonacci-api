# fibonacci-api

To run the server clone this repo and run `make up` from the base of the repo. This will create the Go binary and run the docker compose commands to run the server and postgres database.

The endpoints available are:
```/health -> reutrns health of the server
   /api/v1/fibonacci/current -> gets current value of fibonacci sequence
   /api/v1/fibonacci/next  -> gets next value of fibonacci sequence
   /api/v1/fibonacci/previous -> gets previous value of fibonacci sequence
   /api/v1/fibonacci/reset -> resets fibonacci sequence to 0 
 ```

 ***Note***
 
 The server can take up to 10 seconds to start upon first setup as it has to wait for the database to be ready, there is a retry mechanism built in. Once the Postgres container is able to take connections the server can restart and be ready almost immeditely.
