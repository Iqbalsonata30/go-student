# Golang Repository Pattern
Simple REST API that represents data of students. this project implements repository pattern

<div>
  <h4>Packages what has been used </h4>
  <ul>
    <li><a href="http://github.com/google/uuid">uuid</a> - UUID</li>
    <li><a href="https://github.com/lib/pq">pq</a> - Postgres Driver</li>
    <li><a href="https://github.com/julienschmidt/httprouter">httprouter</a> - HTTP Request Router </li>
    <li><a href="https://github.com/go-playground/validator">validator</a> - Golang Validator</li>
  </ul>
</div>

## Endpoints

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
        "id":string,
        "name":string,
        "identityNumber":int,
        "gender":string,
        "major":string,
        "class":string,
        "religion":string,
        "createdAt":timestamp,
        "updatedAt":timestamp
      ]
    }
```
