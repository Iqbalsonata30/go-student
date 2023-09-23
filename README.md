# Golang Repository Pattern
Simple REST API that represents data of students. this project implements repository pattern

<div>
  <h4>Packages what has been used </h4>
  <ul>
    <li><a href="http://github.com/google/uuid">uuid</a> - UUID</li>
    <li><a href="https://github.com/lib/pq">pq</a> - Postgres Driver</li>
    <li><a href="https://github.com/julienschmidt/httprouter">httprouter</a> - HTTP Request Router </li>
  </ul>
</div>

## Endpoints

### Create Student 
``` bash
POST api/v1/students
e.g
  {   
    "name":"Iqbal Sonata",
    "identityNumber":1234556789,
    "gender":"Male",
    "major":"Computer Engineering",
    "class":"5-PTK-1",
    "religion":"Islam"
  }
```
