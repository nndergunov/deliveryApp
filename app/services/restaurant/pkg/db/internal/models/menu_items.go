// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// MenuItem is an object representing the database table.
type MenuItem struct {
	ID     int     `boil:"id" json:"id" toml:"id" yaml:"id"`
	MenuID int     `boil:"menu_id" json:"menu_id" toml:"menu_id" yaml:"menu_id"`
	Name   string  `boil:"name" json:"name" toml:"name" yaml:"name"`
	Price  float64 `boil:"price" json:"price" toml:"price" yaml:"price"`
	Course string  `boil:"course" json:"course" toml:"course" yaml:"course"`

	R *menuItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L menuItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MenuItemColumns = struct {
	ID     string
	MenuID string
	Name   string
	Price  string
	Course string
}{
	ID:     "id",
	MenuID: "menu_id",
	Name:   "name",
	Price:  "price",
	Course: "course",
}

var MenuItemTableColumns = struct {
	ID     string
	MenuID string
	Name   string
	Price  string
	Course string
}{
	ID:     "menu_items.id",
	MenuID: "menu_items.menu_id",
	Name:   "menu_items.name",
	Price:  "menu_items.price",
	Course: "menu_items.course",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var MenuItemWhere = struct {
	ID     whereHelperint
	MenuID whereHelperint
	Name   whereHelperstring
	Price  whereHelperfloat64
	Course whereHelperstring
}{
	ID:     whereHelperint{field: "\"menu_items\".\"id\""},
	MenuID: whereHelperint{field: "\"menu_items\".\"menu_id\""},
	Name:   whereHelperstring{field: "\"menu_items\".\"name\""},
	Price:  whereHelperfloat64{field: "\"menu_items\".\"price\""},
	Course: whereHelperstring{field: "\"menu_items\".\"course\""},
}

// MenuItemRels is where relationship names are stored.
var MenuItemRels = struct {
	Menu string
}{
	Menu: "Menu",
}

// menuItemR is where relationships are stored.
type menuItemR struct {
	Menu *Menu `boil:"Menu" json:"Menu" toml:"Menu" yaml:"Menu"`
}

// NewStruct creates a new relationship struct
func (*menuItemR) NewStruct() *menuItemR {
	return &menuItemR{}
}

func (r *menuItemR) GetMenu() *Menu {
	if r == nil {
		return nil
	}
	return r.Menu
}

// menuItemL is where Load methods for each relationship are stored.
type menuItemL struct{}

var (
	menuItemAllColumns            = []string{"id", "menu_id", "name", "price", "course"}
	menuItemColumnsWithoutDefault = []string{"menu_id", "name", "price", "course"}
	menuItemColumnsWithDefault    = []string{"id"}
	menuItemPrimaryKeyColumns     = []string{"id"}
	menuItemGeneratedColumns      = []string{}
)

type (
	// MenuItemSlice is an alias for a slice of pointers to MenuItem.
	// This should almost always be used instead of []MenuItem.
	MenuItemSlice []*MenuItem
	// MenuItemHook is the signature for custom MenuItem hook methods
	MenuItemHook func(boil.Executor, *MenuItem) error

	menuItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	menuItemType                 = reflect.TypeOf(&MenuItem{})
	menuItemMapping              = queries.MakeStructMapping(menuItemType)
	menuItemPrimaryKeyMapping, _ = queries.BindMapping(menuItemType, menuItemMapping, menuItemPrimaryKeyColumns)
	menuItemInsertCacheMut       sync.RWMutex
	menuItemInsertCache          = make(map[string]insertCache)
	menuItemUpdateCacheMut       sync.RWMutex
	menuItemUpdateCache          = make(map[string]updateCache)
	menuItemUpsertCacheMut       sync.RWMutex
	menuItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var menuItemAfterSelectHooks []MenuItemHook

var menuItemBeforeInsertHooks []MenuItemHook
var menuItemAfterInsertHooks []MenuItemHook

var menuItemBeforeUpdateHooks []MenuItemHook
var menuItemAfterUpdateHooks []MenuItemHook

var menuItemBeforeDeleteHooks []MenuItemHook
var menuItemAfterDeleteHooks []MenuItemHook

var menuItemBeforeUpsertHooks []MenuItemHook
var menuItemAfterUpsertHooks []MenuItemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MenuItem) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MenuItem) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MenuItem) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MenuItem) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MenuItem) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MenuItem) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MenuItem) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MenuItem) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MenuItem) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range menuItemAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMenuItemHook registers your hook function for all future operations.
func AddMenuItemHook(hookPoint boil.HookPoint, menuItemHook MenuItemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		menuItemAfterSelectHooks = append(menuItemAfterSelectHooks, menuItemHook)
	case boil.BeforeInsertHook:
		menuItemBeforeInsertHooks = append(menuItemBeforeInsertHooks, menuItemHook)
	case boil.AfterInsertHook:
		menuItemAfterInsertHooks = append(menuItemAfterInsertHooks, menuItemHook)
	case boil.BeforeUpdateHook:
		menuItemBeforeUpdateHooks = append(menuItemBeforeUpdateHooks, menuItemHook)
	case boil.AfterUpdateHook:
		menuItemAfterUpdateHooks = append(menuItemAfterUpdateHooks, menuItemHook)
	case boil.BeforeDeleteHook:
		menuItemBeforeDeleteHooks = append(menuItemBeforeDeleteHooks, menuItemHook)
	case boil.AfterDeleteHook:
		menuItemAfterDeleteHooks = append(menuItemAfterDeleteHooks, menuItemHook)
	case boil.BeforeUpsertHook:
		menuItemBeforeUpsertHooks = append(menuItemBeforeUpsertHooks, menuItemHook)
	case boil.AfterUpsertHook:
		menuItemAfterUpsertHooks = append(menuItemAfterUpsertHooks, menuItemHook)
	}
}

// One returns a single menuItem record from the query.
func (q menuItemQuery) One(exec boil.Executor) (*MenuItem, error) {
	o := &MenuItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for menu_items")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MenuItem records from the query.
func (q menuItemQuery) All(exec boil.Executor) (MenuItemSlice, error) {
	var o []*MenuItem

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MenuItem slice")
	}

	if len(menuItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MenuItem records in the query.
func (q menuItemQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count menu_items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q menuItemQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if menu_items exists")
	}

	return count > 0, nil
}

// Menu pointed to by the foreign key.
func (o *MenuItem) Menu(mods ...qm.QueryMod) menuQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MenuID),
	}

	queryMods = append(queryMods, mods...)

	return Menus(queryMods...)
}

// LoadMenu allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (menuItemL) LoadMenu(e boil.Executor, singular bool, maybeMenuItem interface{}, mods queries.Applicator) error {
	var slice []*MenuItem
	var object *MenuItem

	if singular {
		object = maybeMenuItem.(*MenuItem)
	} else {
		slice = *maybeMenuItem.(*[]*MenuItem)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &menuItemR{}
		}
		args = append(args, object.MenuID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &menuItemR{}
			}

			for _, a := range args {
				if a == obj.MenuID {
					continue Outer
				}
			}

			args = append(args, obj.MenuID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`menus`),
		qm.WhereIn(`menus.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Menu")
	}

	var resultSlice []*Menu
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Menu")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for menus")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for menus")
	}

	if len(menuItemAfterSelectHooks) != 0 {
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
		object.R.Menu = foreign
		if foreign.R == nil {
			foreign.R = &menuR{}
		}
		foreign.R.MenuItems = append(foreign.R.MenuItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MenuID == foreign.ID {
				local.R.Menu = foreign
				if foreign.R == nil {
					foreign.R = &menuR{}
				}
				foreign.R.MenuItems = append(foreign.R.MenuItems, local)
				break
			}
		}
	}

	return nil
}

// SetMenu of the menuItem to the related item.
// Sets o.R.Menu to related.
// Adds o to related.R.MenuItems.
func (o *MenuItem) SetMenu(exec boil.Executor, insert bool, related *Menu) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"menu_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"menu_id"}),
		strmangle.WhereClause("\"", "\"", 2, menuItemPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MenuID = related.ID
	if o.R == nil {
		o.R = &menuItemR{
			Menu: related,
		}
	} else {
		o.R.Menu = related
	}

	if related.R == nil {
		related.R = &menuR{
			MenuItems: MenuItemSlice{o},
		}
	} else {
		related.R.MenuItems = append(related.R.MenuItems, o)
	}

	return nil
}

// MenuItems retrieves all the records using an executor.
func MenuItems(mods ...qm.QueryMod) menuItemQuery {
	mods = append(mods, qm.From("\"menu_items\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"menu_items\".*"})
	}

	return menuItemQuery{q}
}

// FindMenuItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMenuItem(exec boil.Executor, iD int, selectCols ...string) (*MenuItem, error) {
	menuItemObj := &MenuItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"menu_items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, menuItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from menu_items")
	}

	if err = menuItemObj.doAfterSelectHooks(exec); err != nil {
		return menuItemObj, err
	}

	return menuItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MenuItem) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no menu_items provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(menuItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	menuItemInsertCacheMut.RLock()
	cache, cached := menuItemInsertCache[key]
	menuItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			menuItemAllColumns,
			menuItemColumnsWithDefault,
			menuItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(menuItemType, menuItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(menuItemType, menuItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"menu_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"menu_items\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into menu_items")
	}

	if !cached {
		menuItemInsertCacheMut.Lock()
		menuItemInsertCache[key] = cache
		menuItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the MenuItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MenuItem) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	menuItemUpdateCacheMut.RLock()
	cache, cached := menuItemUpdateCache[key]
	menuItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			menuItemAllColumns,
			menuItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update menu_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"menu_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, menuItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(menuItemType, menuItemMapping, append(wl, menuItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update menu_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for menu_items")
	}

	if !cached {
		menuItemUpdateCacheMut.Lock()
		menuItemUpdateCache[key] = cache
		menuItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q menuItemQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for menu_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for menu_items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MenuItemSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"menu_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, menuItemPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in menuItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all menuItem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MenuItem) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no menu_items provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(menuItemColumnsWithDefault, o)

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

	menuItemUpsertCacheMut.RLock()
	cache, cached := menuItemUpsertCache[key]
	menuItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			menuItemAllColumns,
			menuItemColumnsWithDefault,
			menuItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			menuItemAllColumns,
			menuItemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert menu_items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(menuItemPrimaryKeyColumns))
			copy(conflict, menuItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"menu_items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(menuItemType, menuItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(menuItemType, menuItemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert menu_items")
	}

	if !cached {
		menuItemUpsertCacheMut.Lock()
		menuItemUpsertCache[key] = cache
		menuItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single MenuItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MenuItem) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MenuItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), menuItemPrimaryKeyMapping)
	sql := "DELETE FROM \"menu_items\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from menu_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for menu_items")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q menuItemQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no menuItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from menu_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for menu_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MenuItemSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(menuItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"menu_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, menuItemPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from menuItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for menu_items")
	}

	if len(menuItemAfterDeleteHooks) != 0 {
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
func (o *MenuItem) Reload(exec boil.Executor) error {
	ret, err := FindMenuItem(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MenuItemSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MenuItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"menu_items\".* FROM \"menu_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, menuItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MenuItemSlice")
	}

	*o = slice

	return nil
}

// MenuItemExists checks if the MenuItem row exists.
func MenuItemExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"menu_items\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if menu_items exists")
	}

	return exists, nil
}
