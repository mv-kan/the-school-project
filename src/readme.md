# Golang source code

## Clean Architecture
The main ideology I applied to this project is clean architecture one. It is ok and I used to use it so no surprise I used it again. 
## Transaction vs Repositories
I exposed db implementation in service to be able create transaction between repository actions. 
## Unit testing with containers
I use testcontainers go to create 