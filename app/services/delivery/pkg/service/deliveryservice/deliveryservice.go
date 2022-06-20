package deliveryservice

import (
	"errors"
	"strconv"
	"time"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/service"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/service/deliveryservice/tools"
)

// DeliveryService is the interface for the user service.
type DeliveryService interface {
	GetEstimateDelivery(estimate *domain.EstimateDeliveryRequest) (*domain.EstimateDeliveryResponse, error)
	AssignOrder(orderID string, order *domain.Order) (*domain.AssignOrder, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DeliveryStorage  service.DeliveryStorage
	Logger           *logger.Logger
	RestaurantClient service.RestaurantClient
	CourierClient    service.CourierClient
	ConsumerClient   service.ConsumerClient
}

type deliveryService struct {
	deliveryStorage  service.DeliveryStorage
	logger           *logger.Logger
	restaurantClient service.RestaurantClient
	courierClient    service.CourierClient
	consumerClient   service.ConsumerClient
}

// NewDeliveryService constructs a new NewDeliveryService.
func NewDeliveryService(p Params) DeliveryService {
	deliveryServiceItem := &deliveryService{
		deliveryStorage:  p.DeliveryStorage,
		logger:           p.Logger,
		restaurantClient: p.RestaurantClient,
		courierClient:    p.CourierClient,
		consumerClient:   p.ConsumerClient,
	}

	return deliveryServiceItem
}

// GetEstimateDelivery get delivery time and cost based on location from and to. Calculating based on average delivery time in the city by car
func (c *deliveryService) GetEstimateDelivery(estimate *domain.EstimateDeliveryRequest) (*domain.EstimateDeliveryResponse, error) {
	consumerLocation, err := c.consumerClient.GetConsumerLocation(estimate.ConsumerID)
	if err != nil {
		return nil, systemErr
	}

	restaurant, err := c.restaurantClient.GetRestaurant(estimate.RestaurantID)
	if err != nil {
		return nil, systemErr
	}

	var estimateDelivery domain.EstimateDeliveryResponse
	// check latitude and longitude
	if consumerLocation == nil || restaurant == nil {
		return nil, errWrongLocData
	}
	if consumerLocation.Latitude != "" || consumerLocation.Longitude != "" ||
		restaurant.Location.Latitude != "" || restaurant.Location.Longitude != "" {

		fromLocationLatF, err := strconv.ParseFloat(consumerLocation.Latitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongFromLocLatType
		}

		fromLocationLonF, err := strconv.ParseFloat(consumerLocation.Longitude, 64)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongFromLonLatType
		}

		toLocationLatF, err := strconv.ParseFloat(restaurant.Location.Latitude, 64)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongToLocLatType
		}

		toLocationLonF, err := strconv.ParseFloat(restaurant.Location.Longitude, 64)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongToLonLatType
		}

		distanceKm, err := tools.VincentyDistance(domain.Coord{
			Lat: fromLocationLatF,
			Lon: fromLocationLonF,
		}, domain.Coord{
			Lat: toLocationLatF,
			Lon: toLocationLonF,
		})
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		// should be considered
		// 1.available couriers in this city
		// 2.number of orders which is in pending status
		// 3. Road and weather condition

		// but for now considered average delivery time/km and cost/km in the city multiplied by distance
		estimateDelivery.Time = getTimeByDistance(distanceKm).String()
		estimateDelivery.Cost = getCostByDistance(distanceKm)

		return &estimateDelivery, nil
	}

	// getting time by address
	if consumerLocation.City != "" || restaurant.Location.City != "" {

		fromAddress := consumerLocation.City + consumerLocation.Region + consumerLocation.Street + consumerLocation.HomeNumber
		fromAddressCoordinate, err := tools.GetCoordinates(fromAddress)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		toAddress := restaurant.Location.City + restaurant.Location.Region + restaurant.Location.Street + restaurant.Location.HomeNumber
		toAddressCoordinate, err := tools.GetCoordinates(toAddress)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		distnaceKm, err := tools.VincentyDistance(domain.Coord{
			Lat: fromAddressCoordinate.Latitude,
			Lon: fromAddressCoordinate.Longitude,
		}, domain.Coord{
			Lat: toAddressCoordinate.Latitude,
			Lon: toAddressCoordinate.Longitude,
		})
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		estimateDelivery.Time = getTimeByDistance(distnaceKm).String()
		estimateDelivery.Cost = getCostByDistance(distnaceKm)

		return &estimateDelivery, nil
	}

	return nil, errWrongLocData
}

func (c *deliveryService) AssignOrder(orderID string, order *domain.Order) (*domain.AssignOrder, error) {
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		return nil, errWrongOrderIDType
	}

	order.OrderID = orderIDInt

	// go to restaurant service get restaurant location
	restaurant, err := c.restaurantClient.GetRestaurant(order.FromRestaurantID)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	// find available courier near restaurant
	nearestCourier, err := c.courierClient.GetNearestCourier(restaurant.Location, 5)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	if nearestCourier == nil {
		return nil, errors.New("no courier available")
	}

	// assign order to available courier
	assignedOrder, err := c.deliveryStorage.AssignOrder(nearestCourier.ID, orderIDInt)
	if err != nil {
		return nil, systemErr
	}

	// update available courier to false
	_, err = c.courierClient.UpdateCourierAvailable(assignedOrder.CourierID, false)
	if err != nil {
		return nil, systemErr
	}

	return assignedOrder, nil
}

func getTimeByDistance(distance float64) (duration time.Duration) {
	// average delivery time in hour per 1 km in city by car
	averageDeliveryTimeHourPerKM := 0.025

	deliveryTime := distance * averageDeliveryTimeHourPerKM
	return time.Duration(deliveryTime)
}

func getCostByDistance(distance float64) float64 {
	// average delivery cost in $ per 1 km in city by car
	averageDeliveryCostPerKM := 0.11

	deliveryCost := distance * averageDeliveryCostPerKM
	return deliveryCost
}
