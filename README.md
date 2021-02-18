# go-apigateway router
## Router for the Existing Zaiserve API.
[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)
----
### Bit about files and its usages
- _main.go_ - Starting point of this project is from main.go file, where it is just calling a __RequestHandler__ function which it will handle the routes.

- _handler.go_ - __RequestHandler__ function resides in the _handler.go_ file.

- In _handler.go_ we are importing _jwt-go_ , _gin_ (HTTP Go Web framework), _utilities_ (Its just a another go file where it contains HOST API Url and another details about api), and _helper.go_ where it contains few functions which are found repetition.

- In __RequestHandler()__ initializing gin with default middleware (where it contains logger and recovery middleware attached).

- _utilities/endpoint.go_ - Other than helpers and thing, we have utilities/endpoint.go forgive me for naming convention, but this one will gives essential information about the host and which url should we need to hit like that. 

- _helper/error.go_ - Which it haves a function which it will just write out the response for the requests which is going to handle.

### Actual Story of this gateway

From utilities.HOST + url string which in results gives the actual server url. But before that we are performing a check which is whether the URL should contains jwt token or not. 

> Note: In terms of our application, we don't need jwt authentication or signed token for login, registration etc., But initally we are checking only for the login alone, but in future we can include registration also, but as for in our existing microservices we don't have any endpoints for registration so the code adopts only for login alone.

If it isn't login it will checks for jwt token, with that jwt token, we are verifying whether the jwt token is alive and not expired, and it isn't malformed by using the jwt package. If jwt verification succesful, then it is good to go for our server.

In terms of passing our request to server, we don't change any of the request format for the api url, we are using ReverseProxy as for now our local and our server is in different environments or different servers, so I used ReversProxy technique, But SingleHostReverseProxy is like the request is actually hitting a endpoint but it proxies the request to another endpoint within same server.

Example: 
http://gogogo.com/api/v1/users  get hits in browser go will proxy this request to http://gogogo.com/api/v1/verifiedusers It is in terms of single host reverse proxy.
```sh
http://gogogo.com/api/v1/users
```
to
```sh
http://gogogo.com/api/v1/verifiedusers
```


In our scenario we are using ReverseProxy which isn't single host.

Example: 
http://localhost:3000/auth/login get hits in browser our api gateway will proxy this request to http://actualserver.com/api/auth/login. 

```sh
http://localhost:3000/auth/login
```
to
```sh
http://actualserver.com/api/auth/login
```

