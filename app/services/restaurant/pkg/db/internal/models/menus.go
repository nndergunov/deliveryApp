// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Menu is an object representing the database table.
type Menu struct {
	ID           int `boil:"id" json:"id" toml:"id" yaml:"id"`
	RestaurantID int `boil:"restaurant_id" json:"restaurant_id" toml:"restaurant_id" yaml:"restaurant_id"`

	R *menuR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L menuL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MenuColumns = struct {
	ID           string
	RestaurantID string
}{
	ID:           "id",
	RestaurantID: "restaurant_id",
}

var MenuTableColumns = struct {
	ID           string
	RestaurantID string
}{
	ID:           "menus.id",
	RestaurantID: "menus.restaurant_id",
}

// Generated where

var MenuWhere = struct {
	ID           whereHelperint
	RestaurantID whereHelperint
}{
	ID:           whereHelperint{field: "\"menus\".\"id\""},
	RestaurantID: whereHelperint{field: "\"menus\".\"restaurant_id\""},
}

// MenuRels is where relationship names are stored.
var MenuRels = struct {
	Restaurant string
	MenuItems  string
}{
	Restaurant: "Restaurant",
	MenuItems:  "MenuItems",
}

// menuR is where relationships are stored.
type menuR struct {
	Restaurant *Restaurant   `boil:"Restaurant" json:"Restaurant" toml:"Restaurant" yaml:"Restaurant"`
	MenuItems  MenuItemSlice `boil:"MenuItems" json:"MenuItems" toml:"MenuItems" yaml:"MenuItems"`
}

// NewStruct creates a new relationship struct
func (*menuR) NewStruct() *menuR {
	return &menuR{}
}

// menuL is where Load methods for each relationship are stored.
type menuL struct{}

var (
	menuAllColumns            = []string{"id", "restaurant_id"}
	menuColumnsWithoutDefault = []string{"restaurant_id"}
	menuColumnsWithDefault    = []string{"id"}
	menuPrimaryKeyColumns     = []string{"id"}
	menuGeneratedColumns      = []string{}
)

type (
	// MenuSlice is an alias for a slice of pointers to Menu.
	// This should almost always be used instead of []Menu.
	MenuSlice []*Menu
	// MenuHook is the signature for custom Menu hook methods
	MenuHook func(boil.Executor, *Menu) error

	menuQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	menuType                 = reflect.TypeOf(&Menu{})
	menuMapping              = queries.MakeStructMapping(menuType)
	menuPrimaryKeyMapping, _ = queries.BindMapping(menuType, menuMapping, menuPrimaryKeyColumns)
	menuInsertCacheMut       sync.RWMutex
	menuInsertCache          = make(map[string]insertCache)
	menuUpdateCacheMut       sync.RWMutex
	menuUpdateCache          = make(map[string]updateCache)
	menuUpsertCacheMut       sync.RWMutex
	menuUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var menuAfterSelectHooks []MenuHook

var menuBeforeInsertHooks []MenuHook
var menuAfterInsertHooks []MenuHook

var menuBeforeUpdateHooks []MenuHook
var menuAfterUpdateHooks []MenuHook

var menuBeforeDeleteHooks []MenuHook
var menuAfterDeleteHooks []MenuHook

var menuBeforeUpsertHooks []MenuHook
var menuAfterUpsertHooks []MenuHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Menu) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range menuAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Menu) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range menuBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Menu) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range menuAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Menu) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range menuBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Menu) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range menuAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Menu) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range menuBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Menu) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range menuAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Menu) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range menuBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Menu) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range menuAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMenuHook registers your hook function for all future operations.
func AddMenuHook(hookPoint boil.HookPoint, menuHook MenuHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		menuAfterSelectHooks = append(menuAfterSelectHooks, menuHook)
	case boil.BeforeInsertHook:
		menuBeforeInsertHooks = append(menuBeforeInsertHooks, menuHook)
	case boil.AfterInsertHook:
		menuAfterInsertHooks = append(menuAfterInsertHooks, menuHook)
	case boil.BeforeUpdateHook:
		menuBeforeUpdateHooks = append(menuBeforeUpdateHooks, menuHook)
	case boil.AfterUpdateHook:
		menuAfterUpdateHooks = append(menuAfterUpdateHooks, menuHook)
	case boil.BeforeDeleteHook:
		menuBeforeDeleteHooks = append(menuBeforeDeleteHooks, menuHook)
	case boil.AfterDeleteHook:
		menuAfterDeleteHooks = append(menuAfterDeleteHooks, menuHook)
	case boil.BeforeUpsertHook:
		menuBeforeUpsertHooks = append(menuBeforeUpsertHooks, menuHook)
	case boil.AfterUpsertHook:
		menuAfterUpsertHooks = append(menuAfterUpsertHooks, menuHook)
	}
}

// One returns a single menu record from the query.
func (q menuQuery) One(exec boil.Executor) (*Menu, error) {
	o := &Menu{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for menus")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Menu records from the query.
func (q menuQuery) All(exec boil.Executor) (MenuSlice, error) {
	var o []*Menu

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Menu slice")
	}

	if len(menuAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Menu records in the query.
func (q menuQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count menus rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q menuQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if menus exists")
	}

	return count > 0, nil
}

// Restaurant pointed to by the foreign key.
func (o *Menu) Restaurant(mods ...qm.QueryMod) restaurantQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RestaurantID),
	}

	queryMods = append(queryMods, mods...)

	return Restaurants(queryMods...)
}

// MenuItems retrieves all the menu_item's MenuItems with an executor.
func (o *Menu) MenuItems(mods ...qm.QueryMod) menuItemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"menu_items\".\"menu_id\"=?", o.ID),
	)

	return MenuItems(queryMods...)
}

// LoadRestaurant allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (menuL) LoadRestaurant(e boil.Executor, singular bool, maybeMenu interface{}, mods queries.Applicator) error {
	var slice []*Menu
	var object *Menu

	if singular {
		object = maybeMenu.(*Menu)
	} else {
		slice = *maybeMenu.(*[]*Menu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &menuR{}
		}
		args = append(args, object.RestaurantID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &menuR{}
			}

			for _, a := range args {
				if a == obj.RestaurantID {
					continue Outer
				}
			}

			args = append(args, obj.RestaurantID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`restaurants`),
		qm.WhereIn(`restaurants.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Restaurant")
	}

	var resultSlice []*Restaurant
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Restaurant")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for restaurants")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for restaurants")
	}

	if len(menuAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Restaurant = foreign
		if foreign.R == nil {
			foreign.R = &restaurantR{}
		}
		foreign.R.Menu = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RestaurantID == foreign.ID {
				local.R.Restaurant = foreign
				if foreign.R == nil {
					foreign.R = &restaurantR{}
				}
				foreign.R.Menu = local
				break
			}
		}
	}

	return nil
}

// LoadMenuItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (menuL) LoadMenuItems(e boil.Executor, singular bool, maybeMenu interface{}, mods queries.Applicator) error {
	var slice []*Menu
	var object *Menu

	if singular {
		object = maybeMenu.(*Menu)
	} else {
		slice = *maybeMenu.(*[]*Menu)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &menuR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &menuR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`menu_items`),
		qm.WhereIn(`menu_items.menu_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load menu_items")
	}

	var resultSlice []*MenuItem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice menu_items")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on menu_items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for menu_items")
	}

	if len(menuItemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.MenuItems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &menuItemR{}
			}
			foreign.R.Menu = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.MenuID {
				local.R.MenuItems = append(local.R.MenuItems, foreign)
				if foreign.R == nil {
					foreign.R = &menuItemR{}
				}
				foreign.R.Menu = local
				break
			}
		}
	}

	return nil
}

// SetRestaurant of the menu to the related item.
// Sets o.R.Restaurant to related.
// Adds o to related.R.Menu.
func (o *Menu) SetRestaurant(exec boil.Executor, insert bool, related *Restaurant) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"menus\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"restaurant_id"}),
		strmangle.WhereClause("\"", "\"", 2, menuPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RestaurantID = related.ID
	if o.R == nil {
		o.R = &menuR{
			Restaurant: related,
		}
	} else {
		o.R.Restaurant = related
	}

	if related.R == nil {
		related.R = &restaurantR{
			Menu: o,
		}
	} else {
		related.R.Menu = o
	}

	return nil
}

// AddMenuItems adds the given related objects to the existing relationships
// of the menu, optionally inserting them as new records.
// Appends related to o.R.MenuItems.
// Sets related.R.Menu appropriately.
func (o *Menu) AddMenuItems(exec boil.Executor, insert bool, related ...*MenuItem) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.MenuID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"menu_items\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"menu_id"}),
				strmangle.WhereClause("\"", "\"", 2, menuItemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.MenuID = o.ID
		}
	}

	if o.R == nil {
		o.R = &menuR{
			MenuItems: related,
		}
	} else {
		o.R.MenuItems = append(o.R.MenuItems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &menuItemR{
				Menu: o,
			}
		} else {
			rel.R.Menu = o
		}
	}
	return nil
}

// Menus retrieves all the records using an executor.
func Menus(mods ...qm.QueryMod) menuQuery {
	mods = append(mods, qm.From("\"menus\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"menus\".*"})
	}

	return menuQuery{q}
}

// FindMenu retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMenu(exec boil.Executor, iD int, selectCols ...string) (*Menu, error) {
	menuObj := &Menu{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"menus\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, menuObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from menus")
	}

	if err = menuObj.doAfterSelectHooks(exec); err != nil {
		return menuObj, err
	}

	return menuObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Menu) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no menus provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(menuColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	menuInsertCacheMut.RLock()
	cache, cached := menuInsertCache[key]
	menuInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			menuAllColumns,
			menuColumnsWithDefault,
			menuColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(menuType, menuMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"menus\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"menus\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into menus")
	}

	if !cached {
		menuInsertCacheMut.Lock()
		menuInsertCache[key] = cache
		menuInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the Menu.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Menu) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	menuUpdateCacheMut.RLock()
	cache, cached := menuUpdateCache[key]
	menuUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			menuAllColumns,
			menuPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update menus, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"menus\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, menuPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, append(wl, menuPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update menus row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for menus")
	}

	if !cached {
		menuUpdateCacheMut.Lock()
		menuUpdateCache[key] = cache
		menuUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q menuQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for menus")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MenuSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"menus\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, menuPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in menu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all menu")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Menu) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no menus provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(menuColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	menuUpsertCacheMut.RLock()
	cache, cached := menuUpsertCache[key]
	menuUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			menuAllColumns,
			menuColumnsWithDefault,
			menuColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			menuAllColumns,
			menuPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert menus, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(menuPrimaryKeyColumns))
			copy(conflict, menuPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"menus\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(menuType, menuMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert menus")
	}

	if !cached {
		menuUpsertCacheMut.Lock()
		menuUpsertCache[key] = cache
		menuUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single Menu record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Menu) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Menu provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), menuPrimaryKeyMapping)
	sql := "DELETE FROM \"menus\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for menus")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q menuQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no menuQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from menus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for menus")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MenuSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(menuBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"menus\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, menuPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from menu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for menus")
	}

	if len(menuAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Menu) Reload(exec boil.Executor) error {
	ret, err := FindMenu(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MenuSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MenuSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"menus\".* FROM \"menus\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, menuPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MenuSlice")
	}

	*o = slice

	return nil
}

// MenuExists checks if the Menu row exists.
func MenuExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"menus\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if menus exists")
	}

	return exists, nil
}
