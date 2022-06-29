
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

| Parameter         | Type     | Description                                   |
| :---------------- | :------- | :-------------------------------------------- |
| `Name`            | `string` | Restaurant name                               |
| `City`            | `string` | City in which restaurant is located           |
| `AcceptingOrders` | `bool`   | Whether restaurant is accepting orders or not |
| `Address`         | `string` | Restaurant address                            |
| `Longitude`       | `float`  | Restaurant coordinates longitude              |
| `Latitude`        | `float`  | Restaurant coordinates latitude               |


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
