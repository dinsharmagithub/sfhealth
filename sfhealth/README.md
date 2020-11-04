# San Fransisco Health Department Restaurant Score Service (SFRS)
## About
The work is an academic exercise completed in a limited time. Thus it needs a lot on the design and implementation part for being good.
## Objective
- The objective is get the following done in given time
  - Get the implementation done the CRUD operations using GRPC
  - Provide REST based APIs for these operation
## Assumptions
- The the database schema is already flattened and there is no effort done on normalizing it or thinking on any improvements on it. The table column types too thus are not put very thoughtfully.
- The implementation of every aspect cannot be done in short duration so placeholders are kept for further code changes or enhancements at several places.

## Design
![Block Diagram](./img/blockdiagram.png)
The design thought for the requirement is to create two servers.
GRPC server interacts with the database and provides the CRUD operations to the clients within the system. 
The REST server is client for the GPRS server but provides the REST services for these operations to its clients outside the system.
The database used is Postgres based on its ease of configurtion and the focus as explained in the Assumptions section above.
## Needed Improvements
- TODO sections in the code and unit tests for completion
- Use of config package and config.dat file for all hardcoded information
  - REST server details
  - Granular DB details
  - log file path
- Code restructuring and naming improvements 
- Build script creation and better packaging
- Addressing Security aspects
- Nice to have in plan: Putting in a container image

## Using the work
Build script is yet to be created so, in order to run, the following can be done 
1. Create restaurant_scores table a Postgres database using sql/ddl.sql 
1. Provide the Postgres details in config/config.dat file
1. Build and execute grpcServer/main.go for running GRPC server with sfhealth as working directory
1. Build and execute restServer/main.go for running REST server
1. Make REST calls with /list /update /insert options e.g. http://localhost:8080/update
