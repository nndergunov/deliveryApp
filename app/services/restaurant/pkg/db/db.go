package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/db/internal/models"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var defaultContext = context.TODO()

type Database struct {
	db *sql.DB
}

func NewDatabase(dbURL string) (*Database, error) {
	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("NewDatabase: %w", err)
	}

	return &Database{
		db: database,
	}, nil
}

func (d Database) ReturnAllRestaurants() ([]domain.Restaurant, error) {
	dbRestaurants, err := models.Restaurants().All(defaultContext, d.db)
	if err != nil {
		return nil, fmt.Errorf("ReturnAllRestaurants: %w", err)
	}

	restaurants := make([]domain.Restaurant, 0, len(dbRestaurants))

	for _, restaurant := range dbRestaurants {
		restaurants = append(restaurants, domain.Restaurant{
			ID:      restaurant.ID,
			Name:    restaurant.Name,
			City:    restaurant.City,
			Address: restaurant.Address,
		})
	}

	return restaurants, nil
}

func (d Database) getRestaurantID(restaurant domain.Restaurant) (int, error) {
	rest, err := models.Restaurants(qm.Where("name=? and city=? and address=?",
		restaurant.Name, restaurant.City, restaurant.Address)).One(defaultContext, d.db)
	if err != nil {
		return 0, fmt.Errorf("getRestaurantID: %w", err)
	}

	return rest.ID, nil
}

func (d Database) getRestaurantByID(restaurantID int) (*models.Restaurant, error) {
	rest, err := models.Restaurants(qm.Where("id=?", restaurantID)).One(defaultContext, d.db)
	if err != nil {
		return nil, fmt.Errorf("getRestaurantByID: %w", err)
	}

	return rest, nil
}

func (d Database) CreateNewRestaurant(restaurant domain.Restaurant) error {
	restaurantID, err := d.getRestaurantID(restaurant)
	if err != nil {
		return fmt.Errorf("CreateNewRestaurant: %w - error: %v", ErrCouldNotCheckExistence, err)
	}

	if restaurantID != 0 {
		return fmt.Errorf("CreateNewRestaurant: %w under id: %d", ErrExistsInDatabase, restaurantID)
	}

	var dbRestaurant models.Restaurant

	dbRestaurant.Name = restaurant.Name
	dbRestaurant.City = restaurant.City
	dbRestaurant.Address = restaurant.Address

	err = dbRestaurant.Insert(defaultContext, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("AddRestaurant: %w", err)
	}

	return nil
}

func (d Database) UpdateRestaurant(restaurant domain.Restaurant) error {
	dbRestaurant, err := d.getRestaurantByID(restaurant.ID)
	if err != nil {
		return fmt.Errorf("UpdateRestaurant: %w", err)
	}

	dbRestaurant.Name = restaurant.Name
	dbRestaurant.City = restaurant.City
	dbRestaurant.Address = restaurant.Address

	_, err = dbRestaurant.Update(defaultContext, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("UpdateRestaurant: %w", err)
	}

	return nil
}

func (d Database) ReturnMenu(restaurantID int) (*domain.Menu, error) {
	dbMenu, err := models.Menus(qm.Where("restaurant_id=?", restaurantID)).One(defaultContext, d.db)
	if err != nil {
		return nil, fmt.Errorf("ReturnMenu: %w", err)
	}

	dbMenuItems, err := models.MenuItems(qm.Where("menu_id=?", dbMenu.ID)).All(defaultContext, d.db)
	if err != nil {
		return nil, fmt.Errorf("ReturnMenu: %w", err)
	}

	menuItems := make([]domain.MenuItem, 0, len(dbMenuItems))

	for _, dbMenuItem := range dbMenuItems {
		menuItem := domain.MenuItem{
			ID:     dbMenuItem.ID,
			MenuID: dbMenu.ID,
			Name:   dbMenuItem.Name,
			Course: dbMenuItem.Course,
		}

		menuItems = append(menuItems, menuItem)
	}

	return &domain.Menu{
		RestaurantID: restaurantID,
		Items:        menuItems,
	}, nil
}

func (d Database) getMenuID(restaurantID int) (int, error) {
	menu, err := models.Menus(qm.Where("restaurant_id=?", restaurantID)).One(defaultContext, d.db)
	if err != nil {
		return 0, fmt.Errorf("getMenuID: %w", err)
	}

	return menu.ID, nil
}

func (d Database) getMenuByID(menuID int) (*models.Menu, error) {
	menu, err := models.Menus(qm.Where("id=?", menuID)).One(defaultContext, d.db)
	if err != nil {
		return nil, fmt.Errorf("getRestaurantByID: %w", err)
	}

	return menu, nil
}

func (d Database) CreateMenu(menu domain.Menu) error {
	menuID, err := d.getMenuID(menu.RestaurantID)
	if err != nil {
		return fmt.Errorf("CreateMenu: %w - error: %v", ErrCouldNotCheckExistence, err)
	}

	if menuID != 0 {
		return fmt.Errorf("CreateMenu: %w under id: %d", ErrExistsInDatabase, menuID)
	}

	var dbMenu models.Menu

	dbMenu.RestaurantID = menu.RestaurantID

	err = dbMenu.Insert(defaultContext, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("CreateMenu: %w", err)
	}

	for _, item := range menu.Items {
		err = d.AddMenuItem(menu.RestaurantID, item)
		if err != nil {
			return fmt.Errorf("CreateMenu: %w", err)
		}
	}

	return nil
}

func (d Database) getMenuItemID(menuItem domain.MenuItem) (int, error) {
	dbMenuItem, err := models.MenuItems(qm.Where("menu_id=? and name=? and course=?",
		menuItem.MenuID, menuItem.Name, menuItem.Course)).One(defaultContext, d.db)
	if err != nil {
		return 0, fmt.Errorf("getMenuItemID: %w", err)
	}

	return dbMenuItem.ID, nil
}

func (d Database) getMenuItemByID(menuItemID int) (*models.MenuItem, error) {
	menuItem, err := models.MenuItems(qm.Where("id=?", menuItemID)).One(defaultContext, d.db)
	if err != nil {
		return nil, fmt.Errorf("getRestaurantByID: %w", err)
	}

	return menuItem, nil
}

func (d Database) AddMenuItem(restaurantID int, menuItem domain.MenuItem) error {
	menuID, err := d.getMenuID(restaurantID)
	if err != nil {
		return fmt.Errorf("AddMenuItem: %w", err)
	}

	if menuID == 0 {
		return fmt.Errorf("AddMenuItem: %w for table models", ErrNotExistsInDatabase)
	}

	menuItemID, err := d.getMenuItemID(menuItem)
	if err != nil {
		return fmt.Errorf("AddMenuItem: %w - error: %v", ErrCouldNotCheckExistence, err)
	}

	if menuItemID != 0 {
		return fmt.Errorf("AddMenuItem: %w under id: %d", ErrExistsInDatabase, menuItemID)
	}

	var dbMenuItem models.MenuItem

	dbMenuItem.MenuID = menuID
	dbMenuItem.Name = menuItem.Name
	dbMenuItem.Course = menuItem.Course

	err = dbMenuItem.Insert(defaultContext, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("AddMenuItem: %w", err)
	}

	return nil
}

func (d Database) UpdateMenuItem(menuItem domain.MenuItem) error {
	dbMenuItem, err := d.getMenuItemByID(menuItem.ID)
	if err != nil {
		return fmt.Errorf("UpdateMenuItem: %w", err)
	}

	dbMenuItem.Name = menuItem.Name
	dbMenuItem.Course = menuItem.Course

	_, err = dbMenuItem.Update(defaultContext, d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("UpdateMenuItem: %w", err)
	}

	return nil
}

func (d Database) DeleteMenuItem(menuItemID int) error {
	dbMenuItem, err := d.getMenuItemByID(menuItemID)
	if err != nil {
		return fmt.Errorf("DeleteMenuItem: %w", err)
	}

	_, err = dbMenuItem.Delete(defaultContext, d.db)
	if err != nil {
		return fmt.Errorf("DeleteMenuItem: %w", err)
	}

	return nil
}
