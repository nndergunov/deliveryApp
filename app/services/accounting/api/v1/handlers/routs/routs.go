package routs

import (
	"accounting/api/v1/handlers/consumergrp"
	"accounting/api/v1/handlers/curiergrp"
	"accounting/api/v1/handlers/restaurantgrp"
	"accounting/api/v1/web"
	"net/http"
)

// Package routs contains the full set of handler functions and routes
// supported by the routs web api.

const (
	version = "/v1"
)

// InitRoutes binds all the version 1 routes.
func InitRoutes(app *web.AppMux) {
	//consumer
	cgh := consumergrp.NewConsumerHandler(consumergrp.Params{
		Logger:          app.Log,
		ConsumerService: app.ConsumerService,
	})

	app.ServeMux.HandleFunc(version+"/status", cgh.StatusHandler).Methods(http.MethodPost)

	app.ServeMux.HandleFunc(version+"/consumers/account", cgh.InsertNewConsumerAccount).Methods(http.MethodPost)
	app.ServeMux.HandleFunc(version+"/consumers/account/{"+consumergrp.ConsumerIDKey+"}", cgh.GetConsumerAccount).Methods(http.MethodGet)
	app.ServeMux.HandleFunc(version+"/consumers/account/{"+consumergrp.ConsumerIDKey+"}", cgh.DeleteConsumerAccount).Methods(http.MethodDelete)

	app.ServeMux.HandleFunc(version+"/consumers/account/balance-add", cgh.AddToBalanceConsumerAccount).Methods(http.MethodPut)
	app.ServeMux.HandleFunc(version+"/consumers/account/balance-sub", cgh.SubFromBalanceConsumerAccount).Methods(http.MethodPut)
	//courier
	crgh := couriergrp.NewCourierHandler(couriergrp.Params{
		Logger:         app.Log,
		CourierService: app.CourierService,
	})

	app.ServeMux.HandleFunc(version+"/couriers/account", crgh.InsertNewCourierAccount).Methods(http.MethodPost)
	app.ServeMux.HandleFunc(version+"/couriers/account/{"+couriergrp.CourierIDKey+"}", crgh.DeleteCourierAccount).Methods(http.MethodDelete)
	app.ServeMux.HandleFunc(version+"/couriers/account/{"+couriergrp.CourierIDKey+"}", crgh.GetCourierAccount).Methods(http.MethodGet)

	app.ServeMux.HandleFunc(version+"/couriers/account/balance-add", crgh.AddToBalanceCourierAccount).Methods(http.MethodPut)
	app.ServeMux.HandleFunc(version+"/couriers/account/balance-sub", crgh.SubFromBalanceCourierAccount).Methods(http.MethodPut)
	//restaurant
	rgh := restaurantgrp.NewRestaurantHandler(restaurantgrp.Params{
		Logger:            app.Log,
		RestaurantService: app.RestaurantService,
	})

	app.ServeMux.HandleFunc(version+"/restaurants/account", rgh.InsertNewRestaurantAccount).Methods(http.MethodPost)
	app.ServeMux.HandleFunc(version+"/restaurants/account/{"+restaurantgrp.RestaurantIDKey+"}", rgh.GetRestaurantAccount).Methods(http.MethodGet)
	app.ServeMux.HandleFunc(version+"/restaurants/account/{"+restaurantgrp.RestaurantIDKey+"}", rgh.DeleteRestaurantAccount).Methods(http.MethodDelete)

	app.ServeMux.HandleFunc(version+"/restaurants/account/balance-add", rgh.AddBalanceRestaurantAccount).Methods(http.MethodPut)
	app.ServeMux.HandleFunc(version+"/restaurants/account/balance-sub", rgh.SubFromBalanceRestaurantAccount).Methods(http.MethodPut)

}
