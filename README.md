![](/logo.jpg)

**An API for DSA Events**

[![GoDoc](https://godoc.org/github.com/ChicagoDSA/DSA-Events?status.svg)](https://godoc.org/github.com/ChicagoDSA/DSA-Events)
[![GitHub version](https://badge.fury.io/gh/ChicagoDSA%2FDSA-Events.svg)](https://badge.fury.io/gh/ChicagoDSA%2FDSA-Events)
[![Build Status](https://travis-ci.org/ChicagoDSA/DSA-Events.svg?branch=master)](https://travis-ci.org/ChicagoDSA/DSA-Events)

[Dgraph](https://dgraph.io/) is a powerful graph database written in [Go](https://golang.org/). The query language for DGraph is based off of [GraphQL](http://graphql.org/), called [GraphQL+-](https://docs.dgraph.io/query-language/). 

This API is a nice way to interface with Dgraph, and has the 3 following endpoints:

- `/query`
- `/mutate`
- `/alter`

To understand what mutations and alterations are, you can take a look over [here](https://docs.dgraph.io/master/query-language/) to understand the fundamentals of graph database interactions.

---

# Getting Started

## Install Dgraph

Your first step is to install dgraph. You can download and install from DGraph's main site over [here](https://docs.dgraph.io/get-started/#step-1-install-dgraph).

One important note is that this API uses docker as middleware, so please follow the 'From Docker Image' instructions in the link I posted previously. If you don't have docker installed on your machine, you can find the installation instructions from them over [here](https://docs.docker.com/engine/installation/).

## Start DGraph

Run the following command to get your dgraph instance running locally via the included docker compose file:

`docker-compose up -d`

This will start Dgraph Server, Zero (Dgraph Instance Manager) and Ratel (Dgraph Console UI). 

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
	    name
	    description
	    location
	    time
	    working_group {
	      name
	      description
	    }
	    chapter {
	      name
	      location
	      contact {
	        name
	        email
	        facebook
	        twitter
	      }
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
	"name":"Name of event",
	"description":"Event's description.",
	"location": {
		"type": "Point", 
		"coordinates": [17.8803304,-245.6662756]
	},
	"time":"2018-05-01T15:30:00Z",
	"working_group": {
		"name": "Chapter working group",
		"description": "Chapter working group "
	},
	"chapter": {
		"name": "Hosting chapter",
		"location": {
			"type": "Point", 
			"coordinates": [41.8803304,-87.6662756]
		},
		"contact": {
			"name":"John Doe",
			"email":"doe@hotmail.com",
			"facebook":"John Doe",
			"twitter":"@johndoe"
		}
	},
}
```

---

#### Built using:
- [Go](https://golang.org/doc/)
- [Dgraph](https://dgraph.io/)
- [Gin-Gonic](https://gin-gonic.github.io/gin/)
- [Logrus](https://github.com/Sirupsen/logrus)
