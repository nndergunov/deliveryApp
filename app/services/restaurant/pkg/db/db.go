package db

import (
	"database/sql"
	"errors"
	"fmt"

	// Postgres drivers.
	_ "github.com/lib/pq"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/db/internal/models"
	"github.com/nndergunov/deliveryApp/app/services/restaurant/pkg/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

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

func (d Database) getRestaurantID(restaurant domain.Restaurant) (int, error) {
	rest, err := models.Restaurants(qm.Where("name=? and city=? and address=?",
		restaurant.Name, restaurant.City, restaurant.Address)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("getRestaurantID: %w", err)
	}

	return rest.ID, nil
}

func (d Database) getRestaurantByID(restaurantID int) (*models.Restaurant, error) {
	rest, err := models.Restaurants(qm.Where("id=?", restaurantID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("getRestaurantByID: %w", err)
	}

	return rest, nil
}

func (d Database) GetAllRestaurants() ([]domain.Restaurant, error) {
	dbRestaurants, err := models.Restaurants().All(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetAllRestaurants: %w", err)
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

func (d Database) InsertRestaurant(restaurant domain.Restaurant) (int, error) {
	restaurantID, err := d.getRestaurantID(restaurant)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("InsertRestaurant: %w", err)
	}

	if restaurantID != 0 {
		return 0, fmt.Errorf("InsertRestaurant: %w", errExistsInDatabase)
	}

	var dbRestaurant models.Restaurant

	dbRestaurant.Name = restaurant.Name
	dbRestaurant.City = restaurant.City
	dbRestaurant.Address = restaurant.Address

	err = dbRestaurant.Insert(d.db, boil.Infer())
	if err != nil {
		return 0, fmt.Errorf("InsertRestaurant: %w", err)
	}

	restaurantID, err = d.getRestaurantID(restaurant)
	if err != nil {
		return 0, fmt.Errorf("InsertRestaurant: %w", err)
	}

	return restaurantID, nil
}

func (d Database) GetRestaurant(restaurantID int) (*domain.Restaurant, error) {
	dbRestaurant, err := models.Restaurants(qm.Where("id=?", restaurantID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetRestaurant: %w", err)
	}

	rest := &domain.Restaurant{
		ID:      dbRestaurant.ID,
		Name:    dbRestaurant.Name,
		City:    dbRestaurant.City,
		Address: dbRestaurant.Address,
	}

	return rest, nil
}

func (d Database) UpdateRestaurant(restaurant domain.Restaurant) error {
	dbRestaurant, err := d.getRestaurantByID(restaurant.ID)
	if err != nil {
		return fmt.Errorf("UpdateRestaurant: %w", err)
	}

	dbRestaurant.Name = restaurant.Name
	dbRestaurant.City = restaurant.City
	dbRestaurant.Address = restaurant.Address

	_, err = dbRestaurant.Update(d.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("UpdateRestaurant: %w", err)
	}

	return nil
}

func (d Database) DeleteRestaurant(restaurantID int) error {
	dbRestaurant, err := models.Restaurants(qm.Where("id=?", restaurantID)).One(d.db)
	if err != nil {
		return fmt.Errorf("DeleteRestaurant: %w", err)
	}

	_, err = dbRestaurant.Delete(d.db)
	if err != nil {
		return fmt.Errorf("DeleteRestaurant: %w", err)
	}

	return nil
}

func (d Database) getMenuID(restaurantID int) (int, error) {
	menu, err := models.Menus(qm.Where("restaurant_id=?", restaurantID)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("getMenuID: %w", err)
	}

	return menu.ID, nil
}

func (d Database) getMenuByID(menuID int) (*models.Menu, error) {
	menu, err := models.Menus(qm.Where("id=?", menuID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("getRestaurantByID: %w", err)
	}

	return menu, nil
}

func (d Database) InsertMenu(menu domain.Menu) (int, []domain.MenuItem, error) {
	menuID, err := d.getMenuID(menu.RestaurantID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, nil, fmt.Errorf("InsertMenu: %w", err)
	}

	if menuID != 0 {
		return 0, nil, fmt.Errorf("InsertMenu: %w", errExistsInDatabase)
	}

	var dbMenu models.Menu

	dbMenu.RestaurantID = menu.RestaurantID

	err = dbMenu.Insert(d.db, boil.Infer())
	if err != nil {
		return 0, nil, fmt.Errorf("InsertMenu: %w", err)
	}

	for _, item := range menu.Items {
		itemID, err := d.InsertMenuItem(menu.RestaurantID, item)
		if err != nil {
			return 0, nil, fmt.Errorf("InsertMenu: %w", err)
		}

		item.ID = itemID
	}

	menuID, err = d.getMenuID(menu.RestaurantID)
	if err != nil {
		return 0, nil, fmt.Errorf("InsertMenu: %w", err)
	}

	return menuID, menu.Items, nil
}

func (d Database) GetMenu(restaurantID int) (*domain.Menu, error) {
	dbMenu, err := models.Menus(qm.Where("restaurant_id=?", restaurantID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetMenu: %w", err)
	}

	menuItems, err := d.getItemsFromMenu(dbMenu.ID)
	if err != nil {
		return nil, fmt.Errorf("GetMenu: %w", err)
	}

	return &domain.Menu{
		RestaurantID: restaurantID,
		Items:        menuItems,
	}, nil
}

func (d Database) UpdateMenu(menu domain.Menu) error {
	dbMenu, err := models.Menus(qm.Where("restaurant_id=?", menu.RestaurantID)).One(d.db)
	if err != nil {
		return fmt.Errorf("UpdateMenu: %w", err)
	}

	if dbMenu == nil {
		return fmt.Errorf("UpdateMenu: %w", errNotExistsInDatabase)
	}

	for _, item := range menu.Items {
		dbID, err := d.getMenuItemID(item)
		if err != nil {
			return fmt.Errorf("UpdateMenu: %w", err)
		}

		if dbID != 0 {
			err = d.UpdateMenuItem(item)
			if err != nil {
				return fmt.Errorf("UpdateMenu: %w", err)
			}
		} else {
			itemID, err := d.InsertMenuItem(menu.RestaurantID, item)
			if err != nil {
				return fmt.Errorf("UpdateMenu: %w", err)
			}

			item.ID = itemID
		}
	}

	dbMenuItems, err := models.MenuItems(qm.Where("menu_id=?", dbMenu.ID)).All(d.db)
	if err != nil {
		return fmt.Errorf("UpdateMenu: %w", err)
	}

	for _, dbMenuItem := range dbMenuItems {
		var found bool

		for _, menuItem := range menu.Items {
			if dbMenuItem.ID == menuItem.ID {
				found = true

				break
			}
		}

		if !found {
			_, err = dbMenuItem.Delete(d.db)
			if err != nil {
				return fmt.Errorf("UpdateMenu: %w", err)
			}
		}
	}

	return nil
}

func (d Database) DeleteMenu(restaurantID int) error {
	dbMenu, err := models.Menus(qm.Where("restaurant_id=?", restaurantID)).One(d.db)
	if err != nil {
		return fmt.Errorf("DeleteMenu: %w", err)
	}

	dbMenuItems, err := models.MenuItems(qm.Where("menu_id=?", dbMenu.ID)).All(d.db)
	if err != nil {
		return fmt.Errorf("DeleteMenu: %w", err)
	}

	for _, dbMenuItem := range dbMenuItems {
		err = d.DeleteMenuItem(dbMenuItem.ID)
		if err != nil {
			return fmt.Errorf("DeleteMenu: %w", err)
		}
	}

	_, err = dbMenu.Delete(d.db)
	if err != nil {
		return fmt.Errorf("DeleteMenu: %w", err)
	}

	return nil
}

func (d Database) getMenuItemID(menuItem domain.MenuItem) (int, error) {
	dbMenuItem, err := models.MenuItems(qm.Where("menu_id=? and name=? and course=?",
		menuItem.MenuID, menuItem.Name, menuItem.Course)).One(d.db)
	if err != nil {
		return 0, fmt.Errorf("getMenuItemID: %w", err)
	}

	return dbMenuItem.ID, nil
}

func (d Database) getMenuItemByID(menuItemID int) (*models.MenuItem, error) {
	menuItem, err := models.MenuItems(qm.Where("id=?", menuItemID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("getRestaurantByID: %w", err)
	}

	return menuItem, nil
}

func (d Database) InsertMenuItem(restaurantID int, menuItem domain.MenuItem) (int, error) {
	menuID, err := d.getMenuID(restaurantID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("InsertMenuItem: %w", err)
	}

	if menuID == 0 {
		return 0, fmt.Errorf("InsertMenuItem: %w for table models", errNotExistsInDatabase)
	}

	menuItemID, err := d.getMenuItemID(menuItem)
	if err != nil {
		return 0, fmt.Errorf("InsertMenuItem: %w", err)
	}

	if menuItemID != 0 {
		return 0, fmt.Errorf("InsertMenuItem: %w", errExistsInDatabase)
	}

	var dbMenuItem models.MenuItem

	dbMenuItem.MenuID = menuID
	dbMenuItem.Name = menuItem.Name
	dbMenuItem.Course = menuItem.Course

	err = dbMenuItem.Insert(d.db, boil.Infer())
	if err != nil {
		return 0, fmt.Errorf("InsertMenuItem: %w", err)
	}

	menuItemID, err = d.getMenuItemID(menuItem)
	if err != nil {
		return 0, fmt.Errorf("InsertMenuItem: %w", err)
	}

	return menuItemID, nil
}

func (d Database) GetMenuItem(menuItemID int) (*domain.MenuItem, error) {
	dbMenuItem, err := models.MenuItems(qm.Where("id=?", menuItemID)).One(d.db)
	if err != nil {
		return nil, fmt.Errorf("GetMenuItem: %w", err)
	}

	menuItem := &domain.MenuItem{
		ID:     dbMenuItem.ID,
		MenuID: dbMenuItem.MenuID,
		Name:   dbMenuItem.Name,
		Course: dbMenuItem.Course,
	}

	return menuItem, nil
}

func (d Database) getItemsFromMenu(menuID int) ([]domain.MenuItem, error) {
	dbMenuItems, err := models.MenuItems(qm.Where("menu_id=?", menuID)).All(d.db)
	if err != nil {
		return nil, fmt.Errorf("getItemsFromMenu: %w", err)
	}

	menuItems := make([]domain.MenuItem, 0, len(dbMenuItems))

	for _, dbMenuItem := range dbMenuItems {
		menuItem := domain.MenuItem{
			ID:     dbMenuItem.ID,
			MenuID: menuID,
			Name:   dbMenuItem.Name,
			Course: dbMenuItem.Course,
		}

		menuItems = append(menuItems, menuItem)
	}

	return menuItems, nil
}

func (d Database) UpdateMenuItem(menuItem domain.MenuItem) error {
	dbMenuItem, err := d.getMenuItemByID(menuItem.ID)
	if err != nil {
		return fmt.Errorf("UpdateMenuItem: %w", err)
	}

	dbMenuItem.Name = menuItem.Name
	dbMenuItem.Course = menuItem.Course

	_, err = dbMenuItem.Update(d.db, boil.Infer())
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

	_, err = dbMenuItem.Delete(d.db)
	if err != nil {
		return fmt.Errorf("DeleteMenuItem: %w", err)
	}

	return nil
}
