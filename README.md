# Golang Repository Pattern
Simple REST API that represents data of students. this project implements repository pattern

<div>
  <h4>Packages what has been used </h4>
  <ul>
    <li><a href="http://github.com/google/uuid">uuid</a> - UUID</li>
    <li><a href="https://github.com/lib/pq">pq</a> - Postgres Driver</li>
    <li><a href="https://github.com/julienschmidt/httprouter">httprouter</a> - HTTP Request Router </li>
    <li><a href="https://github.com/go-playground/validator">validator</a> - Golang Validator</li>
    <li><a href="https://www.docker.com/">docker</a> - Docker</li>
    <li><a href="https:///www.github.com/golang-jwt/jwt ">go-jwt</a> - Golang JSON Web Token</li>
  </ul>
</div>

## Setup Environtment

```bash
export DB_HOST = ""
export DB_PASSWORD = ""
export DB_DATABASE = ""
export DB_PORT = ""
export DB_USER = ""
```


## Endpoints

### Create User 
```bash
POST api/v1/users
Request e.g 
    {
        "username":string,
        "password":string,
    }
```
### Login 
```bash
POST api/v1/login
Request e.g 
    {
        "username":string,
        "password":string,
    }
```

### Create Student 
``` bash
POST api/v1/students
Request e.g
  {   
    "name":string,
    "identityNumber":int,
    "gender":string,
    "major":string,
    "class":string,
    "religion":string
  }
```
### Get Students
``` bash
GET api/v1/students
Result e.g 
    {
      "statusCode":int,
      "message":string,
      "data":[
         {
          "id":string,
          "name":string,
          "identityNumber":int,
          "gender":string,
          "major":string,
          "class":string,
          "religion":string,
          "createdAt":timestamp,
          "updatedAt":timestamp
         }
      ]
    }
```

### Get Student By ID
``` bash
GET api/v1/students/:id
Result e.g 
    {
        "statusCode":int,
        "message":string,
        "data":{
            "id":string,
            "name":string,
            "identityNumber":int,
            "gender":string,
            "major":string,
            "class":string,
            "religion":string,
            "createdAt":timestamp,
            "updatedAt":timestamp
        }
    }
```
### Update Student By ID
``` bash
PUT api/v1/students/:id
Request e.g
  {   
    "name":string,       
    "identityNumber":int,
    "gender":string,
    "major":string,
    "class":string,
    "religion":string
  }
```
### Delete Student By ID 
``` bash
DELETE api/v1/students/:id
Result e.g
 {
     "statusCode":int,
     "message":string,
     "data":string
 }
```


