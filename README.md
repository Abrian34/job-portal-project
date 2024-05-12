# job-portal-project

DATABASE ENV
========================================================================
SERVER_PORT=2000

DB_DRIVER=postgres
DB_USER=postgres          
DB_PASS=1234   
DB_NAME=JobPortalDB  
DB_HOST=localhost         
DB_PORT=5432              
CLIENT_ORIGIN=localhost:2000

STEPS
========================================================================
1. RUN AUTOMIGRATE, go run main.go migrate
2. RUN COMMAND, go run main.go
3. Hit API from postman.
4. Login to retrieve token.
5. use and input retrieved token on Authorization > bearer token.
 
