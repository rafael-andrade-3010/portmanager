# Port Manager

Saving and getting ports

## Solution

The solution was divided in 2 services and 3 modules:

- portapi - Service responsible to have HTTP Apis that can be consumed by clients
- portservice - Service responsible to store the Port data into the database
- proto - Repository with the proto spec used by both services


## portapi

This service has a very basic a simple HTTP api built using plain golang. it is currently handling only 2 requests

    /POST /client-api/port 
    Saves the body into the database
    body: {  "AEAJM": { "name": "Ajman", "city": "Ajman", "country": "United Arab Emirates", "alias": [], "regions": [], "coordinates": [ 55.5136433, 25.4052165 ], "province": "Ajman", "timezone": "Asia/Dubai", "unlocs": [ "AEAJM" ], "code": "52000" }, "AEAUH": { "name": "Abu Dhabi", "coordinates": [ 54.37, 24.47 ], "city": "Abu Dhabi", "province": "Abu Z¸aby [Abu Dhabi]", "country": "United Arab Emirates", "alias": [], "regions": [], "timezone": "Asia/Dubai", "unlocs": [ "AEAUH" ], "code": "52001" }}
    
    /GET /client-api/port 
    List all saved ports in the database

## portservice
Service that has a very simple GPRC API that is being consumed by portapi service.
This service is responsible to save/update the ports into database and also responsible to fetch data from the database
The database chosen for this example is MongoDB, It was mostly chosen because it was easy to use and demonstrate in the demo. Also it has a good performance for these scenarios.

## proto
Proto Spec implemented by portapi and porservice to integrate with each other


## How to run
To run the test using docker, you can go into main folder (where the docker-compose.yaml is located) and run `docker-compose up --build`  
This will build both services and spin up a Mongodb and connect each other.
You can use the postman collection located in the postman folder of the project to test the apis.


## How to build
To build the project, you need go installed in your, version 1.15. Also you must have modules enabled.
You can build and run all projects separated and can run the same set postman collection to consume it.

## Remarks
This code was done in 2 hours (like the constraint)
If I had more time I would improve a lot of things, among them:

- More tests in all layers of the app.. I made some very basic tests, it can be improved drastically
- More resiliency in the code, to check dependencies
- I would love to use openapi in the API project (I'm really used to it, just didn't had the time)
- I would add some level of security
- I would add a lot of clean coding (The code was made mostly to be performant, not pretty)
- There are some missing parts of the implementation (Update objects instead of just saving, pagination in the GET) - I couldn't finish it in time
- I could use stream in GRPC to improve communication between services

  

 


 
