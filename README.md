![Go](https://github.com/honnuanand/GoMuxPostgres/workflows/Go/badge.svg)


This is a Project to focus my learning path on ***go-lang for the web***. 
Some of the initial Ideas are from This [tutorial/blog](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql) which was very clear to get started.

Build Steps: 
1. Check out the Code from the repository
1. 
    ```
    $ cd GoMuxPostgres
    ```
1. 
    ```
    $ go build -v -o bin/go-mux-api
    ```
1. 
    ```
    $ bin/go-mux-api 
    user=postgres password=pgpass dbname=postgres sslmode=disable
    ```
1. Open a Browser and hit all the end points as per the routes created

