package deliveryservice

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/nndergunov/deliveryApp/app/pkg/logger"

	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/domain"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/service"
	"github.com/nndergunov/deliveryApp/app/services/delivery/pkg/service/deliveryservice/tools"
)

// DeliveryService is the interface for the user service.
type DeliveryService interface {
	GetEstimateDelivery(consumerID, restaurantID string) (*domain.EstimateDeliveryResponse, error)
	AssignOrder(orderID string, order *domain.Order) (*domain.AssignOrder, error)
}

// Params is the input parameter struct for the module that contains its dependencies
type Params struct {
	DeliveryStorage  DeliveryStorage
	Logger           *logger.Logger
	RestaurantClient service.RestaurantClient
	CourierClient    service.CourierClient
	ConsumerClient   service.ConsumerClient
}

type deliveryService struct {
	deliveryStorage  DeliveryStorage
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
func (c *deliveryService) GetEstimateDelivery(consumerID, restaurantID string) (*domain.EstimateDeliveryResponse, error) {
	consumerIDInt, err := strconv.Atoi(consumerID)
	if err != nil {
		return nil, fmt.Errorf("wrong consumer_id err: %q", err)
	}

	restaurantIDInt, err := strconv.Atoi(restaurantID)
	if err != nil {
		return nil, fmt.Errorf("wrong restaurant_id err: %q", err)
	}

	consumerLocation, err := c.consumerClient.GetLocation(consumerIDInt)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	restaurant, err := c.restaurantClient.GetRestaurant(restaurantIDInt)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	var estimateDelivery domain.EstimateDeliveryResponse
	// check latitude and longitude
	if consumerLocation == nil || restaurant == nil {
		return nil, errWrongLocData
	}

	if consumerLocation.Latitude != "" && consumerLocation.Longitude != "" &&
		restaurant.Latitude != 0 && restaurant.Longitude != 0 {

		fromLocationLatF, err := strconv.ParseFloat(consumerLocation.Latitude, 10)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		fromLocationLatF = c.convertToDecimalAfterDot(fromLocationLatF)

		fromLocationLonF, err := strconv.ParseFloat(consumerLocation.Longitude, 64)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		fromLocationLonF = c.convertToDecimalAfterDot(fromLocationLonF)

		toLocationLatF := c.convertToDecimalAfterDot(restaurant.Latitude)

		toLocationLonF := c.convertToDecimalAfterDot(restaurant.Longitude)

		distanceKm, err := tools.VincentyDistance(domain.Coordinate{
			Lat: fromLocationLatF,
			Lon: fromLocationLonF,
		}, domain.Coordinate{
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
		estimateDelivery.Cost, err = getCostByDistance(distanceKm)
		if err != nil {
			c.logger.Println(err)
		}

		return &estimateDelivery, nil
	}

	// getting time by address
	if consumerLocation.City != "" || restaurant.City != "" {

		fromAddress := consumerLocation.City + consumerLocation.Region + consumerLocation.Street + consumerLocation.HomeNumber
		fromAddressCoordinate, err := tools.GetCoordinates(fromAddress)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		toAddress := restaurant.City + restaurant.Address
		toAddressCoordinate, err := tools.GetCoordinates(toAddress)
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		distanceKm, err := tools.VincentyDistance(domain.Coordinate{
			Lat: fromAddressCoordinate.Latitude,
			Lon: fromAddressCoordinate.Longitude,
		}, domain.Coordinate{
			Lat: toAddressCoordinate.Latitude,
			Lon: toAddressCoordinate.Longitude,
		})
		if err != nil {
			c.logger.Println(err)
			return nil, systemErr
		}

		estimateDelivery.Time = getTimeByDistance(distanceKm).String()
		estimateDelivery.Cost, err = getCostByDistance(distanceKm)
		if err != nil {
			c.logger.Println(err)
		}

		return &estimateDelivery, nil
	}

	return nil, errWrongLocData
}

func (c *deliveryService) convertToDecimalAfterDot(f float64) float64 {
	fStr := fmt.Sprintf("%.6f", f)
	f, err := strconv.ParseFloat(fStr, 64)
	if err != nil {
		c.logger.Println(err)
	}
	return f
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
	courierLocationList, err := c.courierClient.GetLocation(restaurant.City)
	if err != nil {
		c.logger.Println(err)
		return nil, systemErr
	}

	if courierLocationList == nil {
		return nil, errors.New("no courier available")
	}
	courierLocation := courierLocationList.LocationResponseList[0]
	// assign order to available courier

	assignOrder := domain.AssignOrder{
		OrderID:   courierLocation.UserID,
		CourierID: orderIDInt,
	}

	assignedOrder, err := c.deliveryStorage.AssignOrder(assignOrder)
	if err != nil {
		return nil, systemErr
	}

	// update available courier to false
	_, err = c.courierClient.UpdateCourierAvailable(assignedOrder.CourierID, "false")
	if err != nil {
		return nil, systemErr
	}

	return assignedOrder, nil
}

func getTimeByDistance(distance float64) (duration time.Duration) {
	// average delivery time in hour per 1 km in city by car
	averageCarSpeedInCityM := 60.0

	deliveryTime := distance * (1 / averageCarSpeedInCityM) * float64(time.Hour)
	round := time.Duration(time.Second)

	return time.Duration(deliveryTime).Round(round)
}

func getCostByDistance(distance float64) (float64, error) {
	// average delivery cost in $ per 1 km in city by car
	averageDeliveryCostPerKM := 0.11

	deliveryCost := distance * averageDeliveryCostPerKM
	deliveryCostStr := fmt.Sprintf("%.2f", deliveryCost)

	return strconv.ParseFloat(deliveryCostStr, 64)
}
