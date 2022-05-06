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

// Restaurant is an object representing the database table.
type Restaurant struct {
	ID      int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name    string `boil:"name" json:"name" toml:"name" yaml:"name"`
	City    string `boil:"city" json:"city" toml:"city" yaml:"city"`
	Address string `boil:"address" json:"address" toml:"address" yaml:"address"`

	R *restaurantR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L restaurantL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RestaurantColumns = struct {
	ID      string
	Name    string
	City    string
	Address string
}{
	ID:      "id",
	Name:    "name",
	City:    "city",
	Address: "address",
}

var RestaurantTableColumns = struct {
	ID      string
	Name    string
	City    string
	Address string
}{
	ID:      "restaurants.id",
	Name:    "restaurants.name",
	City:    "restaurants.city",
	Address: "restaurants.address",
}

// Generated where

var RestaurantWhere = struct {
	ID      whereHelperint
	Name    whereHelperstring
	City    whereHelperstring
	Address whereHelperstring
}{
	ID:      whereHelperint{field: "\"restaurants\".\"id\""},
	Name:    whereHelperstring{field: "\"restaurants\".\"name\""},
	City:    whereHelperstring{field: "\"restaurants\".\"city\""},
	Address: whereHelperstring{field: "\"restaurants\".\"address\""},
}

// RestaurantRels is where relationship names are stored.
var RestaurantRels = struct {
	Menu string
}{
	Menu: "Menu",
}

// restaurantR is where relationships are stored.
type restaurantR struct {
	Menu *Menu `boil:"Menu" json:"Menu" toml:"Menu" yaml:"Menu"`
}

// NewStruct creates a new relationship struct
func (*restaurantR) NewStruct() *restaurantR {
	return &restaurantR{}
}

// restaurantL is where Load methods for each relationship are stored.
type restaurantL struct{}

var (
	restaurantAllColumns            = []string{"id", "name", "city", "address"}
	restaurantColumnsWithoutDefault = []string{"name", "city", "address"}
	restaurantColumnsWithDefault    = []string{"id"}
	restaurantPrimaryKeyColumns     = []string{"id"}
	restaurantGeneratedColumns      = []string{}
)

type (
	// RestaurantSlice is an alias for a slice of pointers to Restaurant.
	// This should almost always be used instead of []Restaurant.
	RestaurantSlice []*Restaurant
	// RestaurantHook is the signature for custom Restaurant hook methods
	RestaurantHook func(boil.Executor, *Restaurant) error

	restaurantQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	restaurantType                 = reflect.TypeOf(&Restaurant{})
	restaurantMapping              = queries.MakeStructMapping(restaurantType)
	restaurantPrimaryKeyMapping, _ = queries.BindMapping(restaurantType, restaurantMapping, restaurantPrimaryKeyColumns)
	restaurantInsertCacheMut       sync.RWMutex
	restaurantInsertCache          = make(map[string]insertCache)
	restaurantUpdateCacheMut       sync.RWMutex
	restaurantUpdateCache          = make(map[string]updateCache)
	restaurantUpsertCacheMut       sync.RWMutex
	restaurantUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var restaurantAfterSelectHooks []RestaurantHook

var restaurantBeforeInsertHooks []RestaurantHook
var restaurantAfterInsertHooks []RestaurantHook

var restaurantBeforeUpdateHooks []RestaurantHook
var restaurantAfterUpdateHooks []RestaurantHook

var restaurantBeforeDeleteHooks []RestaurantHook
var restaurantAfterDeleteHooks []RestaurantHook

var restaurantBeforeUpsertHooks []RestaurantHook
var restaurantAfterUpsertHooks []RestaurantHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Restaurant) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Restaurant) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Restaurant) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Restaurant) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Restaurant) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Restaurant) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Restaurant) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Restaurant) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Restaurant) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range restaurantAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRestaurantHook registers your hook function for all future operations.
func AddRestaurantHook(hookPoint boil.HookPoint, restaurantHook RestaurantHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		restaurantAfterSelectHooks = append(restaurantAfterSelectHooks, restaurantHook)
	case boil.BeforeInsertHook:
		restaurantBeforeInsertHooks = append(restaurantBeforeInsertHooks, restaurantHook)
	case boil.AfterInsertHook:
		restaurantAfterInsertHooks = append(restaurantAfterInsertHooks, restaurantHook)
	case boil.BeforeUpdateHook:
		restaurantBeforeUpdateHooks = append(restaurantBeforeUpdateHooks, restaurantHook)
	case boil.AfterUpdateHook:
		restaurantAfterUpdateHooks = append(restaurantAfterUpdateHooks, restaurantHook)
	case boil.BeforeDeleteHook:
		restaurantBeforeDeleteHooks = append(restaurantBeforeDeleteHooks, restaurantHook)
	case boil.AfterDeleteHook:
		restaurantAfterDeleteHooks = append(restaurantAfterDeleteHooks, restaurantHook)
	case boil.BeforeUpsertHook:
		restaurantBeforeUpsertHooks = append(restaurantBeforeUpsertHooks, restaurantHook)
	case boil.AfterUpsertHook:
		restaurantAfterUpsertHooks = append(restaurantAfterUpsertHooks, restaurantHook)
	}
}

// One returns a single restaurant record from the query.
func (q restaurantQuery) One(exec boil.Executor) (*Restaurant, error) {
	o := &Restaurant{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for restaurants")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Restaurant records from the query.
func (q restaurantQuery) All(exec boil.Executor) (RestaurantSlice, error) {
	var o []*Restaurant

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Restaurant slice")
	}

	if len(restaurantAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Restaurant records in the query.
func (q restaurantQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count restaurants rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q restaurantQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if restaurants exists")
	}

	return count > 0, nil
}

// Menu pointed to by the foreign key.
func (o *Restaurant) Menu(mods ...qm.QueryMod) menuQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"restaurant_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	return Menus(queryMods...)
}

// LoadMenu allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (restaurantL) LoadMenu(e boil.Executor, singular bool, maybeRestaurant interface{}, mods queries.Applicator) error {
	var slice []*Restaurant
	var object *Restaurant

	if singular {
		object = maybeRestaurant.(*Restaurant)
	} else {
		slice = *maybeRestaurant.(*[]*Restaurant)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &restaurantR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &restaurantR{}
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
		qm.From(`menus`),
		qm.WhereIn(`menus.restaurant_id in ?`, args...),
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

	if len(restaurantAfterSelectHooks) != 0 {
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
		foreign.R.Restaurant = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ID == foreign.RestaurantID {
				local.R.Menu = foreign
				if foreign.R == nil {
					foreign.R = &menuR{}
				}
				foreign.R.Restaurant = local
				break
			}
		}
	}

	return nil
}

// SetMenu of the restaurant to the related item.
// Sets o.R.Menu to related.
// Adds o to related.R.Restaurant.
func (o *Restaurant) SetMenu(exec boil.Executor, insert bool, related *Menu) error {
	var err error

	if insert {
		related.RestaurantID = o.ID

		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"menus\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"restaurant_id"}),
			strmangle.WhereClause("\"", "\"", 2, menuPrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}
		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.RestaurantID = o.ID
	}

	if o.R == nil {
		o.R = &restaurantR{
			Menu: related,
		}
	} else {
		o.R.Menu = related
	}

	if related.R == nil {
		related.R = &menuR{
			Restaurant: o,
		}
	} else {
		related.R.Restaurant = o
	}
	return nil
}

// Restaurants retrieves all the records using an executor.
func Restaurants(mods ...qm.QueryMod) restaurantQuery {
	mods = append(mods, qm.From("\"restaurants\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"restaurants\".*"})
	}

	return restaurantQuery{q}
}

// FindRestaurant retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRestaurant(exec boil.Executor, iD int, selectCols ...string) (*Restaurant, error) {
	restaurantObj := &Restaurant{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"restaurants\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, restaurantObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from restaurants")
	}

	if err = restaurantObj.doAfterSelectHooks(exec); err != nil {
		return restaurantObj, err
	}

	return restaurantObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Restaurant) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no restaurants provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(restaurantColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	restaurantInsertCacheMut.RLock()
	cache, cached := restaurantInsertCache[key]
	restaurantInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			restaurantAllColumns,
			restaurantColumnsWithDefault,
			restaurantColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(restaurantType, restaurantMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(restaurantType, restaurantMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"restaurants\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"restaurants\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into restaurants")
	}

	if !cached {
		restaurantInsertCacheMut.Lock()
		restaurantInsertCache[key] = cache
		restaurantInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the Restaurant.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Restaurant) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	restaurantUpdateCacheMut.RLock()
	cache, cached := restaurantUpdateCache[key]
	restaurantUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			restaurantAllColumns,
			restaurantPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update restaurants, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"restaurants\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, restaurantPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(restaurantType, restaurantMapping, append(wl, restaurantPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update restaurants row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for restaurants")
	}

	if !cached {
		restaurantUpdateCacheMut.Lock()
		restaurantUpdateCache[key] = cache
		restaurantUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q restaurantQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for restaurants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for restaurants")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RestaurantSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restaurantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"restaurants\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, restaurantPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in restaurant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all restaurant")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Restaurant) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no restaurants provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(restaurantColumnsWithDefault, o)

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

	restaurantUpsertCacheMut.RLock()
	cache, cached := restaurantUpsertCache[key]
	restaurantUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			restaurantAllColumns,
			restaurantColumnsWithDefault,
			restaurantColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			restaurantAllColumns,
			restaurantPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert restaurants, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(restaurantPrimaryKeyColumns))
			copy(conflict, restaurantPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"restaurants\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(restaurantType, restaurantMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(restaurantType, restaurantMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert restaurants")
	}

	if !cached {
		restaurantUpsertCacheMut.Lock()
		restaurantUpsertCache[key] = cache
		restaurantUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single Restaurant record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Restaurant) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Restaurant provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), restaurantPrimaryKeyMapping)
	sql := "DELETE FROM \"restaurants\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from restaurants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for restaurants")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q restaurantQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no restaurantQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from restaurants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for restaurants")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RestaurantSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(restaurantBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restaurantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"restaurants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, restaurantPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from restaurant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for restaurants")
	}

	if len(restaurantAfterDeleteHooks) != 0 {
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
func (o *Restaurant) Reload(exec boil.Executor) error {
	ret, err := FindRestaurant(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RestaurantSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RestaurantSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restaurantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"restaurants\".* FROM \"restaurants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, restaurantPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RestaurantSlice")
	}

	*o = slice

	return nil
}

// RestaurantExists checks if the Restaurant row exists.
func RestaurantExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"restaurants\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if restaurants exists")
	}

	return exists, nil
}
