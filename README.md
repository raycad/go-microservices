# Golang Microservices

Examples of designing and implementing <strong>Microservices</strong> using <strong>Golang</strong>.

### 1. Movie Microservices Architecture
![Movie Services](./docs/images/movie_microservice_architecture.png)

#### 1.1. Microservices Communication Flow In General
| Terms | Description |
|-|-|
| **3rd App Service**| Outternal services that want to use/integrate with Movie service. |
| **API Gateway** | The single entry point for all clients. The API gateway proxied/routed requests to the appropriate service. |
| **IAM Service** | Service that authenticates user and grants authorizations (permissions).|
| **Resource Service** | Service that hosts protected resources and validates access token during resource access. |

![Communication Flows](./docs/images/access_resource.png)

#### 1.2. Flows Of Recommendation Service
Flows of Recommendation Service with other services
![Recommendation Service Flows](./docs/images/recommendation_service.png)

### 2. Movie Service Example
#### 2.1. Setup development environement
* Install Golang
* Install MongoDB
<em>https://docs.mongodb.com/manual/installation</em>
* Start MongoDB
```sh
$ mongod --dbpath="[your_database_path]"
```
* Install neccessary Golang packages 
> $ go get -u github.com/swaggo/swag/cmd/swag github.com/swaggo/gin-swagger github.com/swaggo/gin-swagger/swaggerFiles github.com/alecthomas/template github.com/gin-gonic/gin github.com/sirupsen/logrus gopkg.in/mgo.v2/bson github.com/natefinch/lumberjack

#### 2.1. Create DB information
* Create <strong>go-microservices</strong> database
* Create collections <strong>users</strong> and <strong>movies</strong>
* Create admin account for <strong>users</strong> collection
```sh
{
    "name" : "admin",
    "password" : "admin"
}
```
#### 2.2. Compile and Run services
##### - Generate API documentation using Swag

<em>Guideline: https://github.com/swaggo/swag</em>

* Go to the main application file of the service
```sh
$ cd [main_application_directory]
```

* Generate API document to the ./docs folder using swag cli
```sh
$ swag init
```
<em><strong>NOTE:</strong> You can change host and basePath of the service by editting the file <strong>./docs/docs.go</strong></em>

##### - Change configuration file
<em>Update values of file ./config/config.json</em>
```sh
{
    "port": ":8808",
    "enableGinConsoleLog": true,
    "enableGinFileLog": false,

    "logFilename": "logs/server.log",
    "logMaxSize": 10,
    "logMaxBackups": 10,
    "logMaxAge": 30,

    "mgAddrs": "127.0.0.1:27017",
    "mgDbName": "go-microservices",
    "mgDbUsername": "",
    "mgDbPassword": "",

    "jwtSecretPassword": "raycad",
    "issuer": "seedotech"
}
```
##### -  Run the service
```sh
$ go run main.go
>> [GIN-debug] Listening and serving HTTP on :8808
```
##### - Swagger
<em>http://localhost:8808/swagger/index.html</em>

![Communication Flows](./docs/images/swagger_rest_api.png)

![Communication Flows](./docs/images/swagger_rest_models.png)

### 3. REST API Response Format
<strong>3.1.</strong> Successful response returns the application data through **HTTP 200 OK Message**
###### Example
```sh
* Returning User array (Code 200)
[
  {
    "id": "5bbc4dd782ebac0a8d2c02fe",
    "name": "user name",
    "password": "user password"
  },
  {
    "id": "5bbdadf782ebac06a695a8e7",
    "name": "raycad",
    "password": "raycad"
  }
]

* Returning token string (Code 200)
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicmF5Y2FkIiwicm9sZSI6IiIsImV4cCI6MTUzOTI0OTc3OSwiaXNzIjoic2VlZG90ZWNoIn0.lVHq8J_0vfvECbplogAOCDCNh63ivTBOhya8KE6Ew_E"
}
```

<strong>3.2.</strong> Failed response returns the application error information (**NOT HTTP 200 OK Message**)
```sh
type Error struct {
	Code    int    `json:"code" example:"27"`
	Message string `json:"message" example:"Error message"`
}

Example:
* Returning HTTP StatusBadRequest (Code 400)
{
    "code": -1,
    "message": "User name is empty"
}

* Returning HTTP StatusNotFound (Code 404)
{
    "code": -1,
    "message": "User not found"
}
```

### 4. Coding Convention
##### 4.1. MongoDB Naming Convention
###### COLLECTION NAMES
* The name should be a <strong>plural</strong> of the types of documents to be saved.
* Use <strong>camelCase</strong>. Normally you shouldn’t have to because the collection name will be one word (plural of the types of documents saved).
###### DATABASE NAMES
* Try to have the database named after the project and one database per project.
* Use <strong>camelCase</strong>.
###### FIELD NAMES
* Use <strong>camelCase</strong>.
* Don’t use <strong>_ underscore</strong> as the starting character of the field name. The only field with <strong>_ undescore</strong> should be <strong>_id</strong>.

##### 4.2. Golang Coding Convention
https://github.com/felixge/node-style-guide
