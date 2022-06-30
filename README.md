
# deliveryApp

This project is an example of a simple backend delivery app that uses microservice pattern.

## Run Locally

Clone the project

```bash
  git clone https://github.com/nndergunov/deliveryApp
```

Go to the project directory

```bash
  cd ./deliveryApp
```

Start all the databases, services and message broker

```bash
  make docker-build-all
```

## API Reference

### Restaurant service

#### Post restaurant

```http
  curl -X POST http://localhost:8086/v1/admin/restaurants -H 'Content-Type: application/json' -d '{"Name":"your_restaurant_name", "City":"your_restaurant_city", "AcceptingOrders":true/false, "Address":"Address", "Longitude":<your_restaurant_longitude>, "Latitude":<your_restaurant_latitude>}'
```


#### Post restaurant menu

```http
  curl -X POST http://localhost:8086/v1/admin/restaurants/${id}/menu -H 'Content-Type: application/json' -d '{"RestaurantID":<Restaurant ID>,"Items":[{"ID":0,"MenuID":0,"Name":"menu_item_name","Price":<item price>,"Course":"item_course"}, ...]}'
```

| Parameter | Type  | Description                             |
| :-------- | :---- | :-------------------------------------- |
| `id`      | `int` | **Required**. Id of restaurant to fetch |


#### Get all restaurants (menus are not returned)

```http
  curl -X GET http://localhost:8086/v1/restaurants
```


#### Get restaurant by Id

```http
  curl -X GET http://localhost:8086/v1/restaurants/${id}
```

| Parameter | Type  | Description                             |
| :-------- | :---- | :-------------------------------------- |
| `id`      | `int` | **Required**. Id of restaurant to fetch |


#### Get restaurant menu

```http
  curl -X GET http://localhost:8086/v1/restaurants/${id}/menu
```

| Parameter | Type  | Description                                             |
| :-------- | :-----| :------------------------------------------------------ |
| `id`      | `int` | **Required**. Id of restaurant from which to fetch menu |


### Order service

#### Post order

```http
  curl -X POST http://localhost:8085/v1/orders -H 'Content-Type: application/json' -d '{"FromUserID":<your_user_id>, "RestaurantID":<restaurant_id>, "OrderItems":[<item_1_id>, <item_2_id>, ...]'
```


### Kitchen service

#### Get information of what to cook

```http
  curl -X GET http://localhost:8084/v1/tasks/${id}
```

| Parameter | Type  | Description                                              |
| :-------- | :-----| :------------------------------------------------------- |
| `id`      | `int` | **Required**. Id of restaurant from which to fetch tasks |
