# Code test


## Description & expectation

We have three endpoints

* https://5c3ce12c29429300143fe570.mockapi.io/api/registeredusers returns all the registered users in the system
* https://5c3ce12c29429300143fe570.mockapi.io/api/unregisteredusers returns all the unregistered users in the system
* https://5c3ce12c29429300143fe570.mockapi.io/api/projectmemberships returns all the project memberships for all users, registered and unregistered, in the system

A new service in Go has been created, with a single endpoint that will return all the users, registered and unregistered, including the project ids the users belong to. If the users do not belong to a project, the project ids attribute would contain an empty array in the response payload.

Example of json payload to be returned from the new `/users` endpoint
```
[
  {
    "id": "1",
    "city": "Jaydashire",
    "company": "Goyette - Renner",
    "country": "South Africa",
    "firstName": "firstName 1",
    "lastName": "lastName 1",
    "organizationType": "organizationType 1",
    "phone": "524.276.1570 x487",
    "state": "SD",
    "zipCode": "68048",
    "disclaimerAccepted": false,
    "languageCode": "en",
    "emailAddress": "last1@mail.com",
    "projectIds": ["1","2"]
  },
  {  
    "id":"21",
    "emailAddress":"email1@somewhere.com",
    "languageCode":"en",
    "registrationId":"jwsMJNOk3oM3hVM5bGcF1",
    "registrationIdGeneratedTime":"156165026851",
    "projectIds": []
  }   
]
```  


## Prereqs

Make sure you have docker installed.

## Usage

To run the tests:

 `docker-compose build test`

 `docker-compose run test`

To spin up the Docker container run

`docker-compose up user-projects-image`

Then visit

 <http://localhost:7777/users>

 in your browser to view the API.
