# TaskManagementSystem

Project Title:
- TaskManagementSystem

Project descp:
- A simple task management system with basic crud apis .

Dependencies:

- Go (download from official package)
- Postgres running on port: 5432 , username: postgres , password: postgres , dbname: taskdb
To install through docker use:
'''
docker run --name task_postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=taskdb \
  -p 5432:5432 \
  -d postgres
'''


BreakDown of design decisions:

- Below is the basic flow of user request.
user -> routers -> handlers -> services -> dao -> db.

Here we can also add controller between handler and services , for manging the workflow . Controller can also be used for connecting to other microservices.

- curenlty we have task router , which connects to resp task handlers ,
- In services we basically have business logics but in this case it is simple calls to dao . 
- DAO(data acces object) is where we defined all the functions required to interact with db.
- Models dir contains the object of resp tables/schema.
- DTO is how the system interacts with the user.
- pkg dir contains basic resuable codes like logging and custom errors(if required).
- repositories dir contains the gorm connections to the db.
- cmd dir has files which are executables.



How to Run:
pwd : TaskManagementSystem
commands :
1) go mod tidy
2) go run cmd/server/main.go



what can be done better:

- we can use some load balancer 
- some database concepts like indexing , shrading , connectionpool
- use cache for freq asked .


Since every tasks is loosely coupled , we can easily integrate comunication to other microservices ,preferbly at controller. And also we can containerise this application .


Output:

1) http://localhost:8080/tasks
{"status_code":"SUCCESS","message":"Success","data":[{"ID":2,"Title":"test_title","Description":"test_desc","Status":"test_status"}]}

http://localhost:8080/tasks?status=pending
{"status_code":"SUCCESS","message":"Success","data":[]}

http://localhost:8080/tasks?page=1&page_size=10&status=test_status
{"status_code":"SUCCESS","message":"Success","data":[{"ID":2,"Title":"test_title","Description":"test_desc","Status":"test_status"}]}


2) http://localhost:8080/tasks/1
{"status_code":"DATA_NOT_FOUND","message":"Data Not Found"}

http://localhost:8080/tasks/2
{"status_code":"SUCCESS","message":"Success","data":{"ID":2,"Title":"test_title","Description":"test_desc","Status":"test_status"}}



For testing rest of the endpoint , use postman or user curl for put , delete, post operations
