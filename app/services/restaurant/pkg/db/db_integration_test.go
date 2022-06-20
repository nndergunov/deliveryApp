package db_test

import (
	"fmt"
	"testing"

	"github.com/nndergunov/deliveryApp/app/pkg/configreader"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/db"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
)

func TestGetAllRestaurants(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
	}{
		{
			name: "Simple fetch restaurants from db test",
			restaurantData: domain.Restaurant{
				ID:              0,
				Name:            "GetAllName",
				AcceptingOrders: false,
				City:            "GetAllCity",
				Address:         "GetAllAddress",
				Longitude:       10,
				Latitude:        20,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			id, _ := database.InsertRestaurant(test.restaurantData)

			restaurants, err := database.GetAllRestaurants()
			if err != nil {
				t.Fatal(err)
			}

			var (
				found      bool
				restaurant domain.Restaurant
			)

			for _, rest := range restaurants {
				if rest.ID == id {
					restaurant = rest
					found = true

					break
				}
			}

			if !found {
				t.Fatalf("Could not find inserted restaurant")
			}

			if test.restaurantData.Name != restaurant.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.restaurantData.Name, restaurant.Name)
			}

			if test.restaurantData.AcceptingOrders != restaurant.AcceptingOrders {
				t.Errorf("AcceptingOrders: Expected: %v, Got: %v",
					test.restaurantData.AcceptingOrders, restaurant.AcceptingOrders)
			}

			if test.restaurantData.City != restaurant.City {
				t.Errorf("City: Expected: %s, Got: %s", test.restaurantData.City, restaurant.City)
			}

			if test.restaurantData.Address != restaurant.Address {
				t.Errorf("Address: Expected: %s, Got: %s", test.restaurantData.Address, restaurant.Address)
			}

			if test.restaurantData.Longitude != restaurant.Longitude {
				t.Errorf("Longitude: Expected: %f, Got: %f", test.restaurantData.Longitude, restaurant.Longitude)
			}

			if test.restaurantData.Latitude != restaurant.Latitude {
				t.Errorf("Latitude: Expected: %f, Got: %f", test.restaurantData.Latitude, restaurant.Latitude)
			}

			_ = database.DeleteRestaurant(id)
		})
	}
}

func TestInsertRestaurant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
	}{
		{
			name: "Simple insert restaurant into db test",
			restaurantData: domain.Restaurant{
				ID:              0,
				Name:            "InsertName",
				AcceptingOrders: true,
				City:            "InsertCity",
				Address:         "InsertAddress",
				Longitude:       1.2,
				Latitude:        3.4,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			id, err := database.InsertRestaurant(test.restaurantData)
			if err != nil {
				t.Fatal(err)
			}

			err = database.DeleteRestaurant(id)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestGetRestaurant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
	}{
		{
			name: "Simple fetch single restaurant from db test",
			restaurantData: domain.Restaurant{
				ID:              0,
				Name:            "GetName",
				AcceptingOrders: true,
				City:            "GetCity",
				Address:         "GetAddress",
				Longitude:       1.2,
				Latitude:        3.4,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			id, _ := database.InsertRestaurant(test.restaurantData)

			restaurant, err := database.GetRestaurant(id)
			if err != nil {
				t.Fatal(err)
			}

			if test.restaurantData.Name != restaurant.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.restaurantData.Name, restaurant.Name)
			}

			if test.restaurantData.AcceptingOrders != restaurant.AcceptingOrders {
				t.Errorf("AcceptingOrders: Expected: %v, Got: %v",
					test.restaurantData.AcceptingOrders, restaurant.AcceptingOrders)
			}

			if test.restaurantData.City != restaurant.City {
				t.Errorf("City: Expected: %s, Got: %s", test.restaurantData.City, restaurant.City)
			}

			if test.restaurantData.Address != restaurant.Address {
				t.Errorf("Address: Expected: %s, Got: %s", test.restaurantData.Address, restaurant.Address)
			}

			if test.restaurantData.Longitude != restaurant.Longitude {
				t.Errorf("Longitude: Expected: %f, Got: %f", test.restaurantData.Longitude, restaurant.Longitude)
			}

			if test.restaurantData.Latitude != restaurant.Latitude {
				t.Errorf("Latitude: Expected: %f, Got: %f", test.restaurantData.Latitude, restaurant.Latitude)
			}

			_ = database.DeleteRestaurant(id)
		})
	}
}

func TestUpdateRestaurant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		initialRestaurantData domain.Restaurant
		updatedRestaurantData domain.Restaurant
	}{
		{
			name: "Simple update restaurant in db test",
			initialRestaurantData: domain.Restaurant{
				ID:              0,
				Name:            "InitialName",
				AcceptingOrders: false,
				City:            "InitialCity",
				Address:         "InitialAddress",
				Longitude:       1.0,
				Latitude:        2.0,
			},
			updatedRestaurantData: domain.Restaurant{
				ID:              0,
				Name:            "UpdatedName",
				AcceptingOrders: true,
				City:            "UpdatedCity",
				Address:         "UpdatedAddress",
				Longitude:       1.1,
				Latitude:        2.2,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			id, err := database.InsertRestaurant(test.initialRestaurantData)

			test.updatedRestaurantData.ID = id

			err = database.UpdateRestaurant(test.updatedRestaurantData)
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteRestaurant(id)
		})
	}
}

func TestDeleteRestaurant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
	}{
		{
			name: "Simple delete restaurant from db test",
			restaurantData: domain.Restaurant{
				ID:              0,
				Name:            "DeleteName",
				AcceptingOrders: false,
				City:            "DeleteCity",
				Address:         "DeleteAddress",
				Longitude:       1.2,
				Latitude:        3.4,
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			id, _ := database.InsertRestaurant(test.restaurantData)

			err = database.DeleteRestaurant(id)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestInsertMenu(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
		menuItemData   domain.MenuItem
	}{
		{
			name: "Simple insert menu to db test",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "InsertMenuName",
				City:    "InsertMenuCity",
				Address: "InsertMenuAddress",
			},
			menuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "InsertName",
				Course: "InsertCity",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID, _ := database.InsertRestaurant(test.restaurantData)

			_, _, err = database.InsertMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{test.menuItemData},
			})
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteRestaurant(restaurantID)
		})
	}
}

func TestGetMenu(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
		menuItemData   domain.MenuItem
	}{
		{
			name: "Simple fetch menu from db test",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "GetMenuName",
				City:    "GetMenuCity",
				Address: "GetMenuAddress",
			},
			menuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "GetName",
				Course: "GetCity",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID, _ := database.InsertRestaurant(test.restaurantData)

			_, _, _ = database.InsertMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{test.menuItemData},
			})

			menu, err := database.GetMenu(restaurantID)
			if err != nil {
				t.Fatal(err)
			}

			menuItem := menu.Items[0]

			if test.menuItemData.Name != menuItem.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.menuItemData.Name, menuItem.Name)
			}

			if test.menuItemData.Course != menuItem.Course {
				t.Errorf("Course: Expected: %s, Got: %s", test.menuItemData.Course, menuItem.Course)
			}

			_ = database.DeleteRestaurant(restaurantID)
		})
	}
}

func TestUpdateMenu(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		restaurantData      domain.Restaurant
		initialMenuItemData domain.MenuItem
		updatedMenuItemData domain.MenuItem
	}{
		{
			name: "Simple update menu in db test",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "UpdateMenuName",
				City:    "UpdateMenuCity",
				Address: "UpdateMenuAddress",
			},
			initialMenuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "InitialName",
				Course: "InitialCity",
			},
			updatedMenuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "UpdatedName",
				Course: "UpdatedCity",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID, _ := database.InsertRestaurant(test.restaurantData)

			menuID, itemData, err := database.InsertMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{test.initialMenuItemData},
			})

			test.updatedMenuItemData.MenuID = menuID
			test.updatedMenuItemData.ID = itemData[0].ID

			err = database.UpdateMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{test.updatedMenuItemData},
			})
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteRestaurant(restaurantID)
		})
	}
}

func TestDeleteMenu(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
		menuItemData   domain.MenuItem
	}{
		{
			name: "Simple delete menu from db test",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "DeleteMenuName",
				City:    "DeleteMenuCity",
				Address: "DeleteMenuAddress",
			},
			menuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "DeleteName",
				Course: "DeleteCity",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID, _ := database.InsertRestaurant(test.restaurantData)

			_, _, err = database.InsertMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{test.menuItemData},
			})

			_ = database.DeleteMenu(restaurantID)
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteRestaurant(restaurantID)
		})
	}
}

func TestAddMenuItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
		menuItemData   domain.MenuItem
	}{
		{
			name: "Simple adding item to menu in db test",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "AddItemName",
				City:    "AddItemCity",
				Address: "AddItemAddress",
			},
			menuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "AddName",
				Course: "AddCity",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID, _ := database.InsertRestaurant(test.restaurantData)

			menuID, _, _ := database.InsertMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{},
			})

			test.menuItemData.MenuID = menuID

			_, err = database.InsertMenuItem(restaurantID, test.menuItemData)
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteRestaurant(restaurantID)
		})
	}
}

func TestGetMenuItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
		menuItemData   domain.MenuItem
	}{
		{
			name: "Simple getting item from a menu in db test",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "GetItemName",
				City:    "GetItemCity",
				Address: "GetItemAddress",
			},
			menuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "GetName",
				Course: "GetCity",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID, _ := database.InsertRestaurant(test.restaurantData)

			menuID, _, _ := database.InsertMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{},
			})

			test.menuItemData.MenuID = menuID

			itemID, _ := database.InsertMenuItem(restaurantID, test.menuItemData)
			if err != nil {
				t.Fatal(err)
			}

			menuItem, err := database.GetMenuItem(itemID)
			if err != nil {
				t.Fatal(err)
			}

			if test.menuItemData.Name != menuItem.Name {
				t.Errorf("Name: Expected: %s, Got: %s", test.menuItemData.Name, menuItem.Name)
			}

			if test.menuItemData.Course != menuItem.Course {
				t.Errorf("Course: Expected: %s, Got: %s", test.menuItemData.Course, menuItem.Course)
			}

			_ = database.DeleteRestaurant(restaurantID)
		})
	}
}

func TestUpdateMenuItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		restaurantData      domain.Restaurant
		initialMenuItemData domain.MenuItem
		updatedMenuItemData domain.MenuItem
	}{
		{
			name: "Simple updating item in menu in db test",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "UpdateItemName",
				City:    "UpdateItemCity",
				Address: "UpdateItemAddress",
			},
			initialMenuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "InitialName",
				Course: "InitialCity",
			},
			updatedMenuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "UpdatedName",
				Course: "UpdatedCity",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID, _ := database.InsertRestaurant(test.restaurantData)

			menuID, _, _ := database.InsertMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{},
			})

			test.initialMenuItemData.MenuID = menuID

			itemID, _ := database.InsertMenuItem(restaurantID, test.initialMenuItemData)

			test.updatedMenuItemData.MenuID = menuID
			test.updatedMenuItemData.ID = itemID

			err = database.UpdateMenuItem(test.updatedMenuItemData)
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteRestaurant(restaurantID)
		})
	}
}

func TestDeleteMenuItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		restaurantData domain.Restaurant
		menuItemData   domain.MenuItem
	}{
		{
			name: "Simple deleting item from a menu in db test",
			restaurantData: domain.Restaurant{
				ID:      0,
				Name:    "DeleteItemName",
				City:    "DeleteItemCity",
				Address: "DeleteItemAddress",
			},
			menuItemData: domain.MenuItem{
				ID:     0,
				MenuID: 0,
				Name:   "DeleteName",
				Course: "DeleteCity",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := configreader.SetConfigFile("config.yaml")
			if err != nil {
				t.Fatal(err)
			}

			dbURL := fmt.Sprintf("host=" + configreader.GetString("database.host") +
				" port=" + configreader.GetString("database.port") +
				" user=" + configreader.GetString("database.user") +
				" password=" + configreader.GetString("database.password") +
				" dbname=" + configreader.GetString("database.dbName") +
				" sslmode=" + configreader.GetString("database.sslmode"))

			database, err := db.NewDatabase(dbURL)
			if err != nil {
				t.Fatal(err)
			}

			restaurantID, _ := database.InsertRestaurant(test.restaurantData)

			menuID, _, _ := database.InsertMenu(domain.Menu{
				RestaurantID: restaurantID,
				Items:        []domain.MenuItem{},
			})

			test.menuItemData.MenuID = menuID

			itemID, _ := database.InsertMenuItem(restaurantID, test.menuItemData)

			err = database.DeleteMenuItem(itemID)
			if err != nil {
				t.Fatal(err)
			}

			_ = database.DeleteRestaurant(restaurantID)
		})
	}
}
