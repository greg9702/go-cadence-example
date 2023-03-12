# **go-cadence-example**

## **Overview**

go-cadence-example is a project for testing / demonstrating capabilities of 
[candence](https://github.com/uber/cadence), focusing on using it as a saga ochestrator. Cadence
introduces concept of workflows: *A fault-oblivious stateful function that orchestrates activities. 
A workflow has full control over which activities are executed, and in which order.* Let's see how
we could apply it for implementing [orchestration-based saga pattern](https://microservices.io/patterns/data/saga.html).

> :warning: *This project is just playground. It contains lot of bad practices (including lack of 
> interfaces, hard coded values etc.) which should  not be followed. Maybe one day I will find some 
> time to fix*


## **Architecture**

We are going to implement simple system consisting of services:
- order - receving requests for a new bookings
- payment - processing the payment oprations
- orchestrator - defining, and controlling workflows execution

The execution chain will look as follow:
```
creating order -> processing payment -> marking order as created / failed
```
Order service will create an order in a `INITATED` state. Then, payment service validates the payment.
if the payment value is greater than `100` the payment validation will fail. Based on the payment service
response, order service is obligated to mark event as `CREATED` or `FAILED`.

The enum for order state look as follow:
```
UNDEFINED OrderStatus = 0
INITATED  OrderStatus = 1
CREATED   OrderStatus = 2
FAILED    OrderStatus = 3
```

Order service acts as a entry point to the system, exposing two REST endpoints:
```
POST /v1/order - creating the order
GET /v1/order/{id} - retreiving the order
```
> :warning: Order service for simplicity stores the orders in simple in memory mapping, so time it is restarted 
> data is reseted.

In addition to these services we will need to start Cadence services:
- cassandra
- prometheus
- node-exporter
- cadence
- cadence-web
- grafana


## **Toolsets**

We will use the golang for services implementation, utilising the [go-kit](https://github.com/go-kit/kit)
library.

## **Setup**

### **Requirements**

```
docker-compose
docker
```

### **Setup**

All services are created using docker-compose.
```
docker-compose up
```

It will start all services. Cadence services take some time to start so please be patient. Once all services are up
and running we need to register (create) a domain for our application.

> *Cadence is backed by a multitenant service. The unit of isolation is called a domain. 
> Each domain acts as a namespace for task list names as well as workflow IDs.*

```
docker run --network=host --rm ubercadence/cli:master --domain samples-domain domain register
```

Now we can trigger our order creation:
```
curl --request POST \
  --url http://localhost:8082/v1/order \
  --header 'Content-Type: application/json' \
  --data '{
	"total_cost": 10,
	"vehicle_no": "123"
}'
```

> To trigger the alternative flow, please change the input total_cost input filed to value > 100


And we check the order object using query:
```
curl --request GET \
  --url http://localhost:8082/v1/order/{id} \
  --header 'Content-Type: application/json'
```


## **Solution analysis**

