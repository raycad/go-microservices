# Golang Microservices

A **Microservices Architecture** consists of a collection of small, autonomous **services**. Each **service** is self-contained and should implement a single business capability. 

Below is an example of designing and implementing **Microservices** using:
* [**Gin Gonic**](https://github.com/gin-gonic/gin)
* [**Traefik**](https://github.com/containous/traefik)
* **MongoDB**

### 1. Movie Microservices Architecture
![Movie Services](./docs/images/arc_movie_microservice_architecture.png)

#### 1.1. Microservices Authentication Flow In General
| Terms | Description |
|-|-|
| **3rd App Service**| Outternal services that want to use/integrate with Movie service. |
| **API Gateway** | The single entry point for all clients. The API gateway proxied/routed requests to the appropriate service. |
| **IAM Service** | Service that authenticates user and grants authorizations (permissions).|
| **Resource Service** | Service that hosts protected resources and validates access token during resource access. |

![Communication Flows](./docs/images/seq_access_resource.png)

#### 1.2. Recommendation Service Authentication Flow

![Recommendation Service Flows](./docs/images/seq_recommendation_service.png)

#### 1.3. Class Diagram Of User Management Service

![User Management Service Class Diagram](./docs/images/cd_user_service.png)

***
### 2. Movie Service Build & Run
#### 2.1. Setup development environement
* Install Golang
* Install MongoDB

    <https://docs.mongodb.com/manual/installation>

* Start MongoDB
```sh
$ mongod --dbpath="[your_database_path]"
```
* Install neccessary Golang packages 
> $ go get -u github.com/swaggo/swag/cmd/swag github.com/swaggo/gin-swagger github.com/swaggo/gin-swagger/swaggerFiles github.com/alecthomas/template github.com/gin-gonic/gin github.com/sirupsen/logrus gopkg.in/mgo.v2/bson github.com/natefinch/lumberjack

#### 2.2. Compile & run services
##### - Generate API documentation using Swag

Guideline: https://github.com/swaggo/swag

* Checkout source code
```sh
$ git clone https://github.com/raycad/go-microservices.git
```

* Go to the main application file of the service
```sh
$ cd [main_application_directory]
```

* Generate API document to the ./docs folder using <strong>swag cli</strong>
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

##### -  Run services
* Run the <strong>Authentication</strong> service
```sh
$ cd user-microservice
$ go run main.go
>> [GIN-debug] Listening and serving HTTP on :8808
```

* <strong>Authentication Swagger</strong>

<em>http://localhost:8808/swagger/index.html</em>

**NOTE:** Using the default admin account **admin/admin** to authenticate the services

![Communication Flows](./docs/images/swagger_user_rest_api.png)

![Communication Flows](./docs/images/swagger_user_rest_models.png)

* Run the <strong>Movie</strong> service
```sh
$ cd movie-microservice
$ go run main.go
>> [GIN-debug] Listening and serving HTTP on :8809
```

* <strong>Movie Swagger</strong>

<em>http://localhost:8809/swagger/index.html</em>

**NOTE:** Using the default admin account **admin/admin** to authenticate the services

![Communication Flows](./docs/images/swagger_movie_rest_api.png)

![Communication Flows](./docs/images/swagger_movie_rest_models.png)

***
### 3. Setup API Gateway
**3.1. Download the latest traefik**

https://github.com/containous/traefik/releases/latest

**3.2. Launch traefik**
* Copy `traefik.toml` from [go-microservices]/traefik to traefik execution folder and update parameters.
```sh
defaultEntryPoints = ["http", "https"]
[entryPoints]
    [entryPoints.http]
    address = ":7777"

[api]
    entryPoint = "http"
    dashboard = true

[file]

[frontends]
    [frontends.usermanagement]
		entrypoints = ["http"]		
		backend="usermanagement"
		[frontends.usermanagement.routes.matchUrl]
			rule="PathPrefixStrip:/seedotech.usermanagement"

	[frontends.moviemanagement]
		entrypoints = ["http"]		
		backend="moviemanagement"
		[frontends.moviemanagement.routes.matchUrl]
			rule="PathPrefixStrip:/seedotech.moviemanagement"

[backends]
    [backends.usermanagement]
        [backends.usermanagement.servers.main1]
			url = "http://192.168.1.9:8808"
			weight = 3

		[backends.usermanagement.servers.main2]
			url = "http://192.168.1.10:8808"
			weight = 1

	[backends.moviemanagement]
        [backends.moviemanagement.servers.main1]
			url = "http://192.168.1.9:8809"
			weight = 3

		[backends.moviemanagement.servers.main2]
			url = "http://192.168.1.10:8809"
			weight = 1

		[backends.moviemanagement.servers.main3]
			url = "http://192.168.1.12:8809"
			weight = 2				            
```
* Start **traefik**
```sh
$ ./traefik -c traefik.toml
```

**Traefik dashboard**

![Traefik Dashboard](./docs/images/traefik_api_gateway.png)

Now, assume that the **Gateway IP** is **192.168.1.8**, then, when sending a request to the gateway like:
`GET http://192.168.1.8:7777/seedotech.usermanagement/api/v1/movies/list`

**Gateway** will route request to
`http://192.168.1.9:8808/api/v1/movies/list`
**OR**
`http://192.168.1.10:8808/api/v1/movies/list`

***
### 4. REST API Response Format
<strong>4.1.</strong> Successful response returns the application data through **HTTP 200 OK Message**
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

<strong>4.2.</strong> Failed response returns the application error information (**NOT HTTP 200 OK Message**)
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

***
### 5. Coding Convention
#### 5.1. MongoDB Naming Convention
###### COLLECTION NAMES
* The name should be a <strong>plural</strong> of the types of documents to be saved.
* Use <strong>camelCase</strong>. Normally you shouldn’t have to because the collection name will be one word (plural of the types of documents saved).
###### DATABASE NAMES
* Try to have the database named after the project and one database per project.
* Use <strong>camelCase</strong>.
###### FIELD NAMES
* Use <strong>camelCase</strong>.
* Don’t use <strong>_ underscore</strong> as the starting character of the field name. The only field with <strong>_ undescore</strong> should be <strong>_id</strong>.

#### 5.2. Golang Coding Convention
* Foler naming: lowercase with the <strong>hyphens</strong> seperated between words. e.g. user-microservice
* Package naming: no underscore between words, <strong>everything is lowercase</strong>. e.g. controller
* Go file naming: no underscore between words, <strong>everything is lowercase</strong>. e.g. moviegenre.go
* Files with the suffix <strong>_test.go</strong> are only compiled and run by the `go test` tool
* Files with os and architecture specific suffixes automatically follow those same constraints, e.g. `name_linux.go` will only build on linux, `name_amd64.go` will only build on amd64. This is the same as having a `//+build amd64` line at the top of the file

***
### 6. Terminology, Concepts
**Discovery Service** e.g. Eureka

**Load Balancing** e.g. Ribbon

**Circuit Breaker** e.g. Hystrix

#### 6.1. Service Registry
**Why do we need service registry?**
* A microservice ecosystem may consist of a high number of services that need to know each other’s location. If we have multiple environments (dev, qa, uat, prod) then configuring all these services can be very time consuming and error prone.
* In a cloud environment the ip address and port of the services are not known in advance. Based on demand new service instances can be added or removed on the fly.

The solution is to create the “telephone book” of services, called service registry.

#### 6.2. Circuit Breaker
**Why are they used?**

A circuit breaker is used to provide stability and prevent cascading failures in distributed systems. These should be used in conjunction with judicious timeouts at the interfaces between remote systems to prevent the failure of a single component from bringing down all components.

As an example, we have a web application interacting with a remote third party web service.
Let’s say the third party has oversold their capacity and their database melts down under load.
Assume that the database fails in such a way that it takes a very long time to hand back an error to the third party web service. This in turn makes calls fail after a long period of time. Back to our web application, the users have noticed that their form submissions take much longer seeming to hang. Well the users do what they know to do which is use the refresh button, adding more requests to their already running requests. This eventually causes the failure of the web application due to resource exhaustion. This will affect all users, even those who are not using functionality dependent on this third party web service.

Introducing circuit breakers on the web service call would cause the requests to begin to fail-fast, letting the user know that something is wrong and that they need not refresh their request. This also confines the failure behavior to only those users that are using functionality dependent on the third party, other users are no longer affected as there is no resource exhaustion. Circuit breakers can also allow savvy developers to mark portions of the site that use the functionality unavailable, or perhaps show some cached content as appropriate while the breaker is open.

**Avoiding cascading failures with circuit breaker**

Cascading failure can happen when multiple services are calling each other. For example if we have the below service topology

![Zuul API Gateway](./docs/images/cascading_failures.png)

If service 1 fails then all other services will fail too. Additional risk is that Service 2 and 4 will keep sending requests to service 1, so it will be completely overloaded and will never be able to recover. The circuit breaker pattern addresses this problem. Just like an electrical switch, when closed electrons (requests) can flow through it, when open the flow is stopped. Hystrix can wrap methods with circuit breaker that uses the following logic :

1. By default the circuit is closed and the original method is executed.
1. The original method throws an exception, the fallback is executed.
1. Error rate hits the threshold, the circuit opens.
1. Until the circuit is open the original method is not executed anymore, only the fallback.
1. After a predefined amount of time the circuit is closed, and the flow starts from the beginning.

**What do they do?**

**1. During normal operation, a circuit breaker is in the Closed state:**

* Exceptions or calls exceeding the configured callTimeout increment a failure counter
* Successes reset the failure count to zero
* When the failure counter reaches a maxFailures count, the breaker is tripped into Open state

**2. While in Open state:**

* All calls fail-fast with a CircuitBreakerOpenException
* After the configured resetTimeout, the circuit breaker enters a Half-Open state

**3. In Half-Open state:**

* The first call attempted is allowed through without failing fast
* All other calls fail-fast with an exception just as in Open state
* If the first call succeeds, the breaker is reset back to Closed state and the resetTimeout is reset
* If the first call fails, the breaker is tripped again into the Open state (as for exponential backoff circuit breaker, the resetTimeout is multiplied by the exponential backoff factor)

**4. State transition listeners:**

* Callbacks can be provided for every state entry via onOpen, onClose, and onHalfOpen
* These are executed in the ExecutionContext provided.

**5. Calls result listeners:**

* Callbacks can be used eg. to collect statistics about all invocations or to react on specific call results like success, failures or timeouts.
* Supported callbacks are: onCallSuccess, onCallFailure, onCallTimeout, onCallBreakerOpen.
* These are executed in the ExecutionContext provided.

![Circuit Breaker States](./docs/images/circuit_breaker_states.png)

A typical scenario is when a REST client is wrapped with circuit breaker. The client makes a call to the backend, after subsequent exceptions the circuit opens and the fallback method is called that is serving a static content or an error message. Circuit breaker can be used in any place where there is a chance that operation will fail, but most often this is used in clients of external systems.

#### 6.3. Zuul API Gateway
The Zuul API Gateway is part of the Netflix OSS package. It is very lightweight and integrates well with Eureka. API Gateway is the single entry point into the microservice ecosystem from the outside world.

![Zuul API Gateway](./docs/images/arc_zuul_api_gateway.png)

**Main features of Zuul**

* Provides a unified access to multiple different microservices
* Hides internal details of the microservice ecosystem
* Load balances across multiple service instances
* Allowes access to services
* Restricts access to internal only services
* Looks up services from Eureka
* Implements filters for authentication or logging purposes

Zuul and NGINX are very similar the key differences are

* Unlike for NGINX Zuul can’t be downloaded as a binary distribution, instead it is running inside a Spring Boot application using the Spring Cloud integration
* Zuul is written in Java, therefore integration with the Java, Spring Boot, Spring Cloud and Netflix OSS ecosystem is quite easy

**Eureka** is the **Service Registry** and part of **Netflix OSS** and Spring Cloud stack just like **Zuul**. 

#### 6.4. Concurrency vs. Parallelism
Concurrency and parallelism are related concepts, but there are small differences. Concurrency means that two or more tasks are making progress even though they might not be executing simultaneously. This can for example be realized with time slicing where parts of tasks are executed sequentially and mixed with parts of other tasks. Parallelism on the other hand arise when the execution can be truly simultaneous.

#### 6.5. Asynchronous vs. Synchronous
A method call is considered synchronous if the caller cannot make progress until the method returns a value or throws an exception. On the other hand, an asynchronous call allows the caller to progress after a finite number of steps, and the completion of the method may be signalled via some additional mechanism (it might be a registered callback, a Future, or a message).

A synchronous API may use blocking to implement synchrony, but this is not a necessity. A very CPU intensive task might give a similar behavior as blocking. In general, it is preferred to use asynchronous APIs, as they guarantee that the system is able to progress. Actors are asynchronous by nature: an actor can progress after a message send without waiting for the actual delivery to happen.

#### 6.6. Non-blocking vs. Blocking
We talk about blocking if the delay of one thread can indefinitely delay some of the other threads. A good example is a resource which can be used exclusively by one thread using mutual exclusion. If a thread holds on to the resource indefinitely (for example accidentally running an infinite loop) other threads waiting on the resource can not progress. In contrast, non-blocking means that no thread is able to indefinitely delay others.

Non-blocking operations are preferred to blocking ones, as the overall progress of the system is not trivially guaranteed when it contains blocking operations.

#### 6.7. Microservices Best Practices
* Model services around the business domain.

* Decentralize everything. Individual teams are responsible for designing and building services. Avoid sharing code or data schemas.

* Data storage should be private to the service that owns the data. Use the best storage for each service and data type.

* Services communicate through well-designed APIs. Avoid leaking implementation details. APIs should model the domain, not the internal implementation of the service.

* Avoid coupling between services. Causes of coupling include shared database schemas and rigid communication protocols.

* Offload cross-cutting concerns, such as authentication and SSL termination, to the gateway.

* Keep domain knowledge out of the gateway. The gateway should handle and route client requests without any knowledge of the business rules or domain logic. Otherwise, the gateway becomes a dependency and can cause coupling between services.

* Services should have loose coupling and high functional cohesion. Functions that are likely to change together should be packaged and deployed together. If they reside in separate services, those services end up being tightly coupled, because a change in one service will require updating the other service. Overly chatty communication between two services may be a symptom of tight coupling and low cohesion.

* Isolate failures. Use resiliency strategies to prevent failures within a service from cascading. See Resiliency patterns and Designing resilient applications.

***
### 7. References
https://doc.akka.io/docs/akka/2.5.4/scala/general/terminology.html

https://doc.akka.io/docs/akka/2.5.4/scala/common/circuitbreaker.html

https://docs.microsoft.com/en-us/azure/architecture/guide/architecture-styles/microservices

https://docs.microsoft.com/en-us/azure/architecture/microservices/microservice-boundaries

**API Gateway**

https://docs.microsoft.com/en-us/azure/architecture/microservices/gateway

**Interservice Communication**

https://docs.microsoft.com/en-us/azure/architecture/microservices/interservice-communication

http://howtocookmicroservices.com/communication

https://dzone.com/articles/building-microservices-inter-process-communication-2

