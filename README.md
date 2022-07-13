# deliveryApp
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![CodeFactor](https://www.codefactor.io/repository/github/nndergunov/deliveryApp/badge)](https://www.codefactor.io/repository/github/nndergunov/deliveryApp)
[![Go Report Card](https://goreportcard.com/badge/github.com/nndergunov/deliveryApp)](https://goreportcard.com/report/github.com/nndergunov/deliveryApp)
[![Repo Size](https://img.shields.io/github/repo-size/nndergunov/deliveryApp)](https://github.com/nndergunov/deliveryApp)

This project is an example of a simple backend delivery app that uses microservice pattern.

## Run Locally

Clone the project

```shell
  git clone https://github.com/nndergunov/deliveryApp
```

Go to the project directory

```shell
  cd ./deliveryApp
```

Start all the databases, services and message broker

```shell
  make docker-build-all
```

## API Reference

### Restaurant service

#### Post restaurant

```shell
curl -X POST http://localhost:8086/v1/admin/restaurants -H 'Content-Type: application/json' -d '{"Name":"your_restaurant_name", "City":"your_restaurant_city", "AcceptingOrders":true/false, "Address":"Address", "Longitude":<your_restaurant_longitude>, "Latitude":<your_restaurant_latitude>}'
```

#### Post restaurant menu

```shell
curl -X POST http://localhost:8086/v1/admin/restaurants/${id}/menu -H 'Content-Type: application/json' -d '{"RestaurantID":<Restaurant ID>,"Items":[{"ID":0,"MenuID":0,"Name":"menu_item_name","Price":<item price>,"Course":"item_course"}, ...]}'
```

| Parameter | Type  | Description                             |
| :-------- | :---- | :-------------------------------------- |
| `id`      | `int` | **Required**. Id of restaurant to fetch |

#### Get all restaurants (menus are not returned)

```shell
curl -X GET http://localhost:8086/v1/restaurants
```

#### Get restaurant by Id

```shell
curl -X GET http://localhost:8086/v1/restaurants/${id}
```

| Parameter | Type  | Description                             |
| :-------- | :---- | :-------------------------------------- |
| `id`      | `int` | **Required**. Id of restaurant to fetch |

#### Get restaurant menu

```shell
curl -X GET http://localhost:8086/v1/restaurants/${id}/menu
```

| Parameter | Type  | Description                                             |
| :-------- | :-----| :------------------------------------------------------ |
| `id`      | `int` | **Required**. Id of restaurant from which to fetch menu |

### Order service

#### Post order

```shell
curl -X POST http://localhost:8085/v1/orders -H 'Content-Type: application/json' -d '{"FromUserID":<your_user_id>, "RestaurantID":<restaurant_id>, "OrderItems":[<item_1_id>, <item_2_id>, ...]'
```

### Kitchen service

#### Get information of what to cook

```shell
curl -X GET http://localhost:8084/v1/tasks/${id}
```

| Parameter | Type  | Description                                              |
| :-------- | :-----| :------------------------------------------------------- |
| `id`      | `int` | **Required**. Id of restaurant from which to fetch tasks |

#### Courier service

##### Post courier

```shell
curl -X POST http://localhost:8082/v1/couriers
```

##### Post courier location

```shell
curl -X POST http://localhost:8082/v1/locations
```

##### Get courier

```shell
curl -X GET http://localhost:8082/v1/couriers/{id}
```

| Parameter | Type  | Description                                    |
| :-------- | :-----|:-----------------------------------------------|
| `id`      | `int` | **Required**. Id of courier to fetch           |

##### Get courier location

```shell
curl -X GET http://localhost:8082/v1/locations/{id}
```

| Parameter | Type  | Description                                    |
| :-------- | :-----|:-----------------------------------------------|
| `id`      | `int` | **Required**. Id of courier to fetch           |

#### Consumer service

##### Post consumer

```shell
curl -X POST http://localhost:8082/v1/consumers
```

##### Post consumer location

```shell
curl -X POST http://localhost:8082/v1/locations
```

##### Get consumer

```shell
curl -X GET http://localhost:8082/v1/couriers/{id}
```

| Parameter | Type  | Description                           |
| :-------- | :-----|:--------------------------------------|
| `id`      | `int` | **Required**. Id of consumer to fetch |

##### Get consumer location

```shell
curl -X GET http://localhost:8082/v1/locations/{id}
```

| Parameter | Type  | Description                           |
| :-------- | :-----|:--------------------------------------|
| `id`      | `int` | **Required**. Id of consumer to fetch |

#### Accounting service

##### Post account

```shell
curl -X GET http://localhost:8080/v1/accounts
```

##### Get account

```shell
curl -X GET http://localhost:8082/v1/accounts/{id}
```

| Parameter | Type  | Description                           |
| :-------- | :-----|:--------------------------------------|
| `id`      | `int` | **Required**. Id of consumer to fetch |

##### Get account list

```shell
curl -X GET http://localhost:8082/v1/accounts
```

| query Parameter | Type     | Description                              |
|:-----------------|:---------|:-----------------------------------------|
| `user_id`        | `string` | **Required**. Id of user to fetch        |
| `user_type`      | `string` | **Required**. user type of user to fetch |

##### Transaction adding to balance

```shell
curl -X POST http://localhost:8082/v1/transactions
```

| body              | Type      | Description                   |
|:------------------|:----------|:------------------------------|
| `to_account_id`   | `int`     | **Required**. to account id   |
| `amount`          | `float64` | **Required**.amount           |

##### Transaction sub from balance

```shell
curl -X POST http://localhost:8082/v1/transactions
```

| body              | Type      | Description                   |
|:------------------|:----------|:------------------------------|
| `from_account_id` | `int`     | **Required**. from account id |
| `amount`          | `float64` | **Required**.amount           |

##### Transaction from balance to balance

```shell
curl -X POST http://localhost:8082/v1/transactions
```

| body              | Type      | Description                   |
|:------------------|:----------|:------------------------------|
| `from_account_id` | `int`     | **Required**. from account id |
| `to_account_id`   | `int`     | **Required**. to account id   |
| `amount`          | `float64` | **Required**.amount           |

#### Delivery service

##### Get estimate values

```shell
curl -X GET http://localhost:8083/v1/estimate
```

| query param     | Type  | Description                 |
|:----------------|:------|:----------------------------|
| `consumer_id`   | `int` | **Required**. consumer id   |
| `restaurant_id` | `int` | **Required**. restaurant id |

##### Assign order to available courier near restaurant

```shell
curl -X POST http://localhost:8083/v1/orders/{id}/assing
```

| body            | Type  | Description                 |
|:----------------|:------|:----------------------------|
| `from_user_id`  | `int` | **Required**. from user id  |
| `restaurant_id` | `int` | **Required**. restaurant id |

