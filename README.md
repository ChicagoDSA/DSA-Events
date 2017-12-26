![](/logo.jpg)

**An API for DSA Events**

[Dgraph](https://dgraph.io/) is a powerful graph database written in [Go](https://golang.org/). The query language for DGraph is based off of [GraphQL](http://graphql.org/), called [GraphQL+-](https://docs.dgraph.io/query-language/). 

This API is a nice way to interface with Dgraph, and has the 3 following endpoints:

- `/query`
- `/mutate`
- `/alter`

To understand what mutations and alterations are, you can take a look over [here](https://tour.dgraph.io/) to understand the fundamentals of graph database interactions.

---

# Getting Started

## Install Dgraph

Your first step is to install dgraph. You can download and install from DGraph's main site over [here](https://docs.dgraph.io/get-started/#step-1-install-dgraph).

One important note is that this API uses docker as middleware, so please follow the 'From Docker Image' instructions in the link I posted previously. If you don't have docker installed on your machine, you can find the installation instructions from them over [here](https://docs.docker.com/engine/installation/).

## Start DGraph

Run the following commands to get your dgraph instance running locally.

### Start DGraph Zero (instance manager):

This will run the dgraphzero container and store temporary data in `/tmp/data`:

`docker run --rm -it -p 8080:8080 -p 9080:9080 -v /tmp/data:/dgraph --name DSA_Events dgraph/dgraph dgraph zero --port_offset -2000`

### Start DGraph Instance:

This will execute the docker image:

`docker exec -it DSA_Events dgraph server --memory_mb 2048 --zero localhost:5080`

### DEPRECATED (Dgraph prior to dgraphzero implementation)
Run DGraph instance with ports mapped:

`docker run --rm -it -p 8080:8080 -p 9080:9080 -v ~/dgraph:/dgraph --name dgraph dgraph/dgraph dgraph --bindall=true --memory_mb 2048`

---

## Install Go dependencies

While you're at the root of the project, run:

`go get ./...`

This will install the required dependencies of the whole project

## Build project

Run:

`make build`

## Run project

Run:

`make run`

---

## GraphQL+- Samples

### GraphQL+- Sample Query
#### POST - http://localhost:5000/query
##### Purpose:
This query will get an Event with all the available parameters. It will get an Event with 'uid' of 0x3.
##### Body:
```
{
	Event(func: uid(0x3)) {
		uid
		Name
		Time
		Description
		Data
		HostingChapter {
			Title
			State
			City
			Contact {
				Name
				PhoneNumber
				Email
				Facebook
				Twitter
			}
		}
		Location {
			Name
			State
			City
			ZipCode
		}
	}
}
```

### GraphQL+- Sample Mutation
#### POST - http://localhost:5000/mutate
##### Purpose:
This is a sample mutation that will create a new Event with the parameters given.
##### Body:
```
{
	"name":"Created Event",
	"date":"08-11-17",
	"time":"14:30",
	"hosting_chapter":{
		"title":"Milwaukee DSA",
		"city":"Milwaukee",
		"state":"WI",
		"contact":{
			"name":"Jeb Bush",
			"phone_number":"123-456-7890",
			"email":"jeb@hotmail.com",
			"facebook":"jebisthebest",
			"twitter":"@jebisgood"
		}
	},
	"description":"Test Description",
	"location":{
		"name":"Location Name!",
		"city":"Milwaukee",
		"state":"WI",
		"zip_code":"12345"
	}
}
```

---

#### Built using:
- [Go](https://golang.org/doc/)
- [Dgraph](https://docs.dgraph.io/)
- [Gin-Gonic](https://gin-gonic.github.io/gin/)
- [Logrus](https://github.com/Sirupsen/logrus)
