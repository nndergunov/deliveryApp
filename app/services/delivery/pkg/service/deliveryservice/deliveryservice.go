package deliveryservice

import (
	"delivery/pkg/domain"
	"delivery/pkg/service"
	"delivery/pkg/service/deliveryservice/tools"
	"errors"
	"github.com/nndergunov/deliveryApp/app/pkg/logger"
	"strconv"
	"time"
)

// DeliveryService is the interface for the user service.
type DeliveryService interface {
	GetDeliveryTime(distance *domain.DeliveryDistanceLocation) (*domain.DeliveryTime, error)
	GetDeliveryCost(distance *domain.DeliveryDistanceLocation) (*domain.DeliveryCost, error)
	AssignCourierToOrder(orderID string, order *domain.Order) (*domain.AssignedCourier, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DeliveryStorage  service.DeliveryStorage
	Logger           *logger.Logger
	RestaurantClient service.RestaurantClient
	CourierClient    service.CourierClient
}

type deliveryService struct {
	deliveryStorage  service.DeliveryStorage
	logger           *logger.Logger
	restaurantClient service.RestaurantClient
	courierClient    service.CourierClient
}

// NewDeliveryService constructs a new NewDeliveryService.
func NewDeliveryService(p Params) DeliveryService {
	deliveryServiceItem := &deliveryService{
		deliveryStorage:  p.DeliveryStorage,
		logger:           p.Logger,
		restaurantClient: p.RestaurantClient,
		courierClient:    p.CourierClient,
	}

	return deliveryServiceItem
}

//GetDeliveryTime get delivery time based on location from and to. Calculating based on average delivery time in the city by car
func (c *deliveryService) GetDeliveryTime(distance *domain.DeliveryDistanceLocation) (*domain.DeliveryTime, error) {
	var deliveryTime domain.DeliveryTime
	//check latitude and longitude
	if distance.FromLocation == nil || distance.ToLocation == nil {
		return nil, errWrongLocData
	}
	//getting time by coordinates
	if distance.FromLocation.Latitude != "" || distance.FromLocation.Longitude != "" ||
		distance.ToLocation.Latitude != "" || distance.ToLocation.Longitude != "" {

		fromLocationLatF, err := strconv.ParseFloat(distance.FromLocation.Latitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongFromLocLatType
		}

		fromLocationLonF, err := strconv.ParseFloat(distance.FromLocation.Longitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongFromLonLatType
		}

		toLocationLatF, err := strconv.ParseFloat(distance.ToLocation.Latitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongToLocLatType
		}

		toLocationLonF, err := strconv.ParseFloat(distance.FromLocation.Longitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongToLonLatType
		}

		_, kilometers, err := tools.VincentyDistance(domain.Coord{
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

		//should be considered
		//1.available couriers in this city
		//2.number of orders which is in pending status
		//3. Road and weather condition

		//but for now only deliveryTime
		deliveryTime.Time = getTimeByDistance(kilometers).String()

		return &deliveryTime, nil
	}

	//getting time by address
	if distance.FromLocation.City != "" || distance.ToLocation.City != "" {
		fromCityCoordinate, err := tools.GetCoordinates(distance.FromLocation.City)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		toCityCoordinate, err := tools.GetCoordinates(distance.ToLocation.City)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		_, kilometers, err := tools.VincentyDistance(domain.Coord{
			Lat: fromCityCoordinate.Latitude,
			Lon: fromCityCoordinate.Longitude,
		}, domain.Coord{
			Lat: toCityCoordinate.Latitude,
			Lon: toCityCoordinate.Longitude,
		})
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		deliveryTime.Time = getTimeByDistance(kilometers).String()
		return &deliveryTime, nil
	}

	return nil, errWrongLocData
}

//GetDeliveryCost get delivery time based on location from and to. Calculating based on average delivery time in the city by car
func (c *deliveryService) GetDeliveryCost(distance *domain.DeliveryDistanceLocation) (*domain.DeliveryCost, error) {
	var deliveryCost domain.DeliveryCost
	//check latitude and longitude
	if distance.FromLocation == nil || distance.ToLocation == nil {
		return nil, errWrongLocData
	}
	//getting time by coordinates
	if distance.FromLocation.Latitude != "" || distance.FromLocation.Longitude != "" ||
		distance.ToLocation.Latitude != "" || distance.ToLocation.Longitude != "" {

		fromLocationLatF, err := strconv.ParseFloat(distance.FromLocation.Latitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongFromLocLatType
		}

		fromLocationLonF, err := strconv.ParseFloat(distance.FromLocation.Longitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongFromLonLatType
		}

		toLocationLatF, err := strconv.ParseFloat(distance.ToLocation.Latitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongToLocLatType
		}

		toLocationLonF, err := strconv.ParseFloat(distance.FromLocation.Longitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, errWrongToLonLatType
		}

		_, kilometers, err := tools.VincentyDistance(domain.Coord{
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

		//should be considered
		//1.available couriers in this city
		//2.number of orders which is in pending status
		//3. Road and weather condition

		//but for now only deliveryCost
		deliveryCost.Cost = getCostByDistance(kilometers)
		return &deliveryCost, nil
	}

	//getting time by address
	if distance.FromLocation.City != "" || distance.ToLocation.City != "" {
		fromCityCoordinate, err := tools.GetCoordinates(distance.FromLocation.City)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		toCityCoordinate, err := tools.GetCoordinates(distance.ToLocation.City)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		_, kilometers, err := tools.VincentyDistance(domain.Coord{
			Lat: fromCityCoordinate.Latitude,
			Lon: fromCityCoordinate.Longitude,
		}, domain.Coord{
			Lat: toCityCoordinate.Latitude,
			Lon: toCityCoordinate.Longitude,
		})
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		//but for now only deliveryCost
		deliveryCost.Cost = getCostByDistance(kilometers)
		return &deliveryCost, nil
	}

	return nil, errWrongLocData
}
func (c *deliveryService) AssignCourierToOrder(orderID string, order *domain.Order) (*domain.AssignedCourier, error) {
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		return nil, errWrongOrderIDType
	}

	order.OrderID = orderIDInt

	//go to restaurantLocation service get restaurantLocation location
	restaurantLocation, err := c.restaurantClient.GetRestaurantLocation(order.FromRestaurantID)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	//find available courier near restaurantLocation
	nearestCourier, err := c.courierClient.GetNearestCourier(restaurantLocation, 5)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	if nearestCourier == nil {
		return nil, errors.New("no courier available")
	}

	//assign order to available courier
	assignedCourier, err := c.deliveryStorage.AssignCourier(nearestCourier.ID, orderIDInt)
	if err != nil {
		return nil, systemErr
	}

	//update available courier to false
	_, err = c.courierClient.UpdateCourierAvailable(assignedCourier.CourierID, false)
	if err != nil {
		return nil, systemErr
	}

	return assignedCourier, nil
}

func getTimeByDistance(distance float64) (duration time.Duration) {
	//average delivery time in hour per 1 km in city by car
	averageDeliveryTimeHourPerKM := 0.025

	deliveryTime := distance * averageDeliveryTimeHourPerKM
	return time.Duration(deliveryTime)
}

func getCostByDistance(distance float64) float64 {
	//average delivery cost in $ per 1 km in city by car
	averageDeliveryCostPerKM := 0.11

	deliveryCost := distance * averageDeliveryCostPerKM
	return deliveryCost
}
