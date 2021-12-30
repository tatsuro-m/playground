// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
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

// TownArea is an object representing the database table.
type TownArea struct {
	ID             int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name           string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	NameRoma       string    `boil:"name_roma" json:"name_roma" toml:"name_roma" yaml:"name_roma"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	MunicipalityID int       `boil:"municipality_id" json:"municipality_id" toml:"municipality_id" yaml:"municipality_id"`

	R *townAreaR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L townAreaL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TownAreaColumns = struct {
	ID             string
	Name           string
	NameRoma       string
	CreatedAt      string
	UpdatedAt      string
	MunicipalityID string
}{
	ID:             "id",
	Name:           "name",
	NameRoma:       "name_roma",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	MunicipalityID: "municipality_id",
}

var TownAreaTableColumns = struct {
	ID             string
	Name           string
	NameRoma       string
	CreatedAt      string
	UpdatedAt      string
	MunicipalityID string
}{
	ID:             "town_areas.id",
	Name:           "town_areas.name",
	NameRoma:       "town_areas.name_roma",
	CreatedAt:      "town_areas.created_at",
	UpdatedAt:      "town_areas.updated_at",
	MunicipalityID: "town_areas.municipality_id",
}

// Generated where

var TownAreaWhere = struct {
	ID             whereHelperint
	Name           whereHelperstring
	NameRoma       whereHelperstring
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
	MunicipalityID whereHelperint
}{
	ID:             whereHelperint{field: "`town_areas`.`id`"},
	Name:           whereHelperstring{field: "`town_areas`.`name`"},
	NameRoma:       whereHelperstring{field: "`town_areas`.`name_roma`"},
	CreatedAt:      whereHelpertime_Time{field: "`town_areas`.`created_at`"},
	UpdatedAt:      whereHelpertime_Time{field: "`town_areas`.`updated_at`"},
	MunicipalityID: whereHelperint{field: "`town_areas`.`municipality_id`"},
}

// TownAreaRels is where relationship names are stored.
var TownAreaRels = struct {
	Municipality string
	PostalCodes  string
}{
	Municipality: "Municipality",
	PostalCodes:  "PostalCodes",
}

// townAreaR is where relationships are stored.
type townAreaR struct {
	Municipality *Municipality   `boil:"Municipality" json:"Municipality" toml:"Municipality" yaml:"Municipality"`
	PostalCodes  PostalCodeSlice `boil:"PostalCodes" json:"PostalCodes" toml:"PostalCodes" yaml:"PostalCodes"`
}

// NewStruct creates a new relationship struct
func (*townAreaR) NewStruct() *townAreaR {
	return &townAreaR{}
}

// townAreaL is where Load methods for each relationship are stored.
type townAreaL struct{}

var (
	townAreaAllColumns            = []string{"id", "name", "name_roma", "created_at", "updated_at", "municipality_id"}
	townAreaColumnsWithoutDefault = []string{"name", "name_roma", "municipality_id"}
	townAreaColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	townAreaPrimaryKeyColumns     = []string{"id"}
)

type (
	// TownAreaSlice is an alias for a slice of pointers to TownArea.
	// This should almost always be used instead of []TownArea.
	TownAreaSlice []*TownArea
	// TownAreaHook is the signature for custom TownArea hook methods
	TownAreaHook func(context.Context, boil.ContextExecutor, *TownArea) error

	townAreaQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	townAreaType                 = reflect.TypeOf(&TownArea{})
	townAreaMapping              = queries.MakeStructMapping(townAreaType)
	townAreaPrimaryKeyMapping, _ = queries.BindMapping(townAreaType, townAreaMapping, townAreaPrimaryKeyColumns)
	townAreaInsertCacheMut       sync.RWMutex
	townAreaInsertCache          = make(map[string]insertCache)
	townAreaUpdateCacheMut       sync.RWMutex
	townAreaUpdateCache          = make(map[string]updateCache)
	townAreaUpsertCacheMut       sync.RWMutex
	townAreaUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var townAreaBeforeInsertHooks []TownAreaHook
var townAreaBeforeUpdateHooks []TownAreaHook
var townAreaBeforeDeleteHooks []TownAreaHook
var townAreaBeforeUpsertHooks []TownAreaHook

var townAreaAfterInsertHooks []TownAreaHook
var townAreaAfterSelectHooks []TownAreaHook
var townAreaAfterUpdateHooks []TownAreaHook
var townAreaAfterDeleteHooks []TownAreaHook
var townAreaAfterUpsertHooks []TownAreaHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TownArea) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TownArea) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TownArea) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TownArea) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TownArea) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TownArea) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TownArea) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TownArea) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TownArea) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range townAreaAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTownAreaHook registers your hook function for all future operations.
func AddTownAreaHook(hookPoint boil.HookPoint, townAreaHook TownAreaHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		townAreaBeforeInsertHooks = append(townAreaBeforeInsertHooks, townAreaHook)
	case boil.BeforeUpdateHook:
		townAreaBeforeUpdateHooks = append(townAreaBeforeUpdateHooks, townAreaHook)
	case boil.BeforeDeleteHook:
		townAreaBeforeDeleteHooks = append(townAreaBeforeDeleteHooks, townAreaHook)
	case boil.BeforeUpsertHook:
		townAreaBeforeUpsertHooks = append(townAreaBeforeUpsertHooks, townAreaHook)
	case boil.AfterInsertHook:
		townAreaAfterInsertHooks = append(townAreaAfterInsertHooks, townAreaHook)
	case boil.AfterSelectHook:
		townAreaAfterSelectHooks = append(townAreaAfterSelectHooks, townAreaHook)
	case boil.AfterUpdateHook:
		townAreaAfterUpdateHooks = append(townAreaAfterUpdateHooks, townAreaHook)
	case boil.AfterDeleteHook:
		townAreaAfterDeleteHooks = append(townAreaAfterDeleteHooks, townAreaHook)
	case boil.AfterUpsertHook:
		townAreaAfterUpsertHooks = append(townAreaAfterUpsertHooks, townAreaHook)
	}
}

// One returns a single townArea record from the query.
func (q townAreaQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TownArea, error) {
	o := &TownArea{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for town_areas")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TownArea records from the query.
func (q townAreaQuery) All(ctx context.Context, exec boil.ContextExecutor) (TownAreaSlice, error) {
	var o []*TownArea

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TownArea slice")
	}

	if len(townAreaAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TownArea records in the query.
func (q townAreaQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count town_areas rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q townAreaQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if town_areas exists")
	}

	return count > 0, nil
}

// Municipality pointed to by the foreign key.
func (o *TownArea) Municipality(mods ...qm.QueryMod) municipalityQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.MunicipalityID),
	}

	queryMods = append(queryMods, mods...)

	query := Municipalities(queryMods...)
	queries.SetFrom(query.Query, "`municipalities`")

	return query
}

// PostalCodes retrieves all the postal_code's PostalCodes with an executor.
func (o *TownArea) PostalCodes(mods ...qm.QueryMod) postalCodeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`postal_codes`.`town_area_id`=?", o.ID),
	)

	query := PostalCodes(queryMods...)
	queries.SetFrom(query.Query, "`postal_codes`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`postal_codes`.*"})
	}

	return query
}

// LoadMunicipality allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (townAreaL) LoadMunicipality(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTownArea interface{}, mods queries.Applicator) error {
	var slice []*TownArea
	var object *TownArea

	if singular {
		object = maybeTownArea.(*TownArea)
	} else {
		slice = *maybeTownArea.(*[]*TownArea)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &townAreaR{}
		}
		args = append(args, object.MunicipalityID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &townAreaR{}
			}

			for _, a := range args {
				if a == obj.MunicipalityID {
					continue Outer
				}
			}

			args = append(args, obj.MunicipalityID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`municipalities`),
		qm.WhereIn(`municipalities.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Municipality")
	}

	var resultSlice []*Municipality
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Municipality")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for municipalities")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for municipalities")
	}

	if len(townAreaAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Municipality = foreign
		if foreign.R == nil {
			foreign.R = &municipalityR{}
		}
		foreign.R.TownAreas = append(foreign.R.TownAreas, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MunicipalityID == foreign.ID {
				local.R.Municipality = foreign
				if foreign.R == nil {
					foreign.R = &municipalityR{}
				}
				foreign.R.TownAreas = append(foreign.R.TownAreas, local)
				break
			}
		}
	}

	return nil
}

// LoadPostalCodes allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (townAreaL) LoadPostalCodes(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTownArea interface{}, mods queries.Applicator) error {
	var slice []*TownArea
	var object *TownArea

	if singular {
		object = maybeTownArea.(*TownArea)
	} else {
		slice = *maybeTownArea.(*[]*TownArea)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &townAreaR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &townAreaR{}
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
		qm.From(`postal_codes`),
		qm.WhereIn(`postal_codes.town_area_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load postal_codes")
	}

	var resultSlice []*PostalCode
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice postal_codes")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on postal_codes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for postal_codes")
	}

	if len(postalCodeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PostalCodes = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &postalCodeR{}
			}
			foreign.R.TownArea = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TownAreaID {
				local.R.PostalCodes = append(local.R.PostalCodes, foreign)
				if foreign.R == nil {
					foreign.R = &postalCodeR{}
				}
				foreign.R.TownArea = local
				break
			}
		}
	}

	return nil
}

// SetMunicipality of the townArea to the related item.
// Sets o.R.Municipality to related.
// Adds o to related.R.TownAreas.
func (o *TownArea) SetMunicipality(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Municipality) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `town_areas` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"municipality_id"}),
		strmangle.WhereClause("`", "`", 0, townAreaPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MunicipalityID = related.ID
	if o.R == nil {
		o.R = &townAreaR{
			Municipality: related,
		}
	} else {
		o.R.Municipality = related
	}

	if related.R == nil {
		related.R = &municipalityR{
			TownAreas: TownAreaSlice{o},
		}
	} else {
		related.R.TownAreas = append(related.R.TownAreas, o)
	}

	return nil
}

// AddPostalCodes adds the given related objects to the existing relationships
// of the town_area, optionally inserting them as new records.
// Appends related to o.R.PostalCodes.
// Sets related.R.TownArea appropriately.
func (o *TownArea) AddPostalCodes(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PostalCode) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TownAreaID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `postal_codes` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"town_area_id"}),
				strmangle.WhereClause("`", "`", 0, postalCodePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TownAreaID = o.ID
		}
	}

	if o.R == nil {
		o.R = &townAreaR{
			PostalCodes: related,
		}
	} else {
		o.R.PostalCodes = append(o.R.PostalCodes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &postalCodeR{
				TownArea: o,
			}
		} else {
			rel.R.TownArea = o
		}
	}
	return nil
}

// TownAreas retrieves all the records using an executor.
func TownAreas(mods ...qm.QueryMod) townAreaQuery {
	mods = append(mods, qm.From("`town_areas`"))
	return townAreaQuery{NewQuery(mods...)}
}

// FindTownArea retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTownArea(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TownArea, error) {
	townAreaObj := &TownArea{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `town_areas` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, townAreaObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from town_areas")
	}

	if err = townAreaObj.doAfterSelectHooks(ctx, exec); err != nil {
		return townAreaObj, err
	}

	return townAreaObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TownArea) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no town_areas provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(townAreaColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	townAreaInsertCacheMut.RLock()
	cache, cached := townAreaInsertCache[key]
	townAreaInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			townAreaAllColumns,
			townAreaColumnsWithDefault,
			townAreaColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(townAreaType, townAreaMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(townAreaType, townAreaMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `town_areas` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `town_areas` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `town_areas` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, townAreaPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into town_areas")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == townAreaMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for town_areas")
	}

CacheNoHooks:
	if !cached {
		townAreaInsertCacheMut.Lock()
		townAreaInsertCache[key] = cache
		townAreaInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TownArea.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TownArea) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	townAreaUpdateCacheMut.RLock()
	cache, cached := townAreaUpdateCache[key]
	townAreaUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			townAreaAllColumns,
			townAreaPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update town_areas, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `town_areas` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, townAreaPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(townAreaType, townAreaMapping, append(wl, townAreaPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update town_areas row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for town_areas")
	}

	if !cached {
		townAreaUpdateCacheMut.Lock()
		townAreaUpdateCache[key] = cache
		townAreaUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q townAreaQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for town_areas")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for town_areas")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TownAreaSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), townAreaPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `town_areas` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, townAreaPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in townArea slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all townArea")
	}
	return rowsAff, nil
}

var mySQLTownAreaUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TownArea) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no town_areas provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(townAreaColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTownAreaUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	townAreaUpsertCacheMut.RLock()
	cache, cached := townAreaUpsertCache[key]
	townAreaUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			townAreaAllColumns,
			townAreaColumnsWithDefault,
			townAreaColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			townAreaAllColumns,
			townAreaPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert town_areas, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`town_areas`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `town_areas` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(townAreaType, townAreaMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(townAreaType, townAreaMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for town_areas")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == townAreaMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(townAreaType, townAreaMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for town_areas")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for town_areas")
	}

CacheNoHooks:
	if !cached {
		townAreaUpsertCacheMut.Lock()
		townAreaUpsertCache[key] = cache
		townAreaUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TownArea record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TownArea) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TownArea provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), townAreaPrimaryKeyMapping)
	sql := "DELETE FROM `town_areas` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from town_areas")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for town_areas")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q townAreaQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no townAreaQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from town_areas")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for town_areas")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TownAreaSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(townAreaBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), townAreaPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `town_areas` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, townAreaPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from townArea slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for town_areas")
	}

	if len(townAreaAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TownArea) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTownArea(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TownAreaSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TownAreaSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), townAreaPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `town_areas`.* FROM `town_areas` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, townAreaPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TownAreaSlice")
	}

	*o = slice

	return nil
}

// TownAreaExists checks if the TownArea row exists.
func TownAreaExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `town_areas` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if town_areas exists")
	}

	return exists, nil
}
