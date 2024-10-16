## Checks

- [x] Get rabbitmq working
  - [x] Connect to localhost port http://localhost:5672/
  - ~Create a init() function that will return the rabbit instance?~
  - [x] Instead we made a variable that lives in the package rabbit

## TODO

- The rabbit instance will send two types of requests to the queue:

  - On connect or at startup that will send the documents that are already on the database to the SolR service
  - When a course is created, updated or deleted that it will ~notify~ send the solr service those changes

- It's easy to send the changes on update/insert/delete, you just publish the new document (in the case of insert/update) on the service package and|or delete the old one
- To send all the existing documents(?)

docker run --hostname rabbitmq --name rabbit-mq -p 15672:15672 -p 5672:5672 rabbitmq:3-management
