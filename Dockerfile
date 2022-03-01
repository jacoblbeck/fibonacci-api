FROM scratch 

ADD release/fibonacci-api /bin/fibonacci-api

ENTRYPOINT ["/bin/fibonacci-api"]