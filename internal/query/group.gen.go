// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"jim/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newGroup(db *gorm.DB) group {
	_group := group{}

	_group.groupDo.UseDB(db)
	_group.groupDo.UseModel(&model.Group{})

	tableName := _group.groupDo.TableName()
	_group.ALL = field.NewField(tableName, "*")
	_group.ID = field.NewInt64(tableName, "id")
	_group.Name = field.NewString(tableName, "name")
	_group.Introduction = field.NewString(tableName, "introduction")
	_group.Extra = field.NewString(tableName, "extra")
	_group.CreatedAt = field.NewTime(tableName, "created_at")
	_group.UpdatedAt = field.NewTime(tableName, "updated_at")
	_group.DeletedAt = field.NewField(tableName, "deleted_at")

	_group.fillFieldMap()

	return _group
}

type group struct {
	groupDo groupDo

	ALL          field.Field
	ID           field.Int64
	Name         field.String
	Introduction field.String
	Extra        field.String
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field

	fieldMap map[string]field.Expr
}

func (g group) As(alias string) *group {
	g.groupDo.DO = *(g.groupDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.ID = field.NewInt64(alias, "id")
	g.Name = field.NewString(alias, "name")
	g.Introduction = field.NewString(alias, "introduction")
	g.Extra = field.NewString(alias, "extra")
	g.CreatedAt = field.NewTime(alias, "created_at")
	g.UpdatedAt = field.NewTime(alias, "updated_at")
	g.DeletedAt = field.NewField(alias, "deleted_at")

	g.fillFieldMap()

	return &g
}

func (g *group) WithContext(ctx context.Context) *groupDo { return g.groupDo.WithContext(ctx) }

func (g group) TableName() string { return g.groupDo.TableName() }

func (g *group) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *group) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["name"] = g.Name
	g.fieldMap["introduction"] = g.Introduction
	g.fieldMap["extra"] = g.Extra
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["deleted_at"] = g.DeletedAt
}

func (g group) clone(db *gorm.DB) group {
	g.groupDo.ReplaceDB(db)
	return g
}

type groupDo struct{ gen.DO }

func (g groupDo) Debug() *groupDo {
	return g.withDO(g.DO.Debug())
}

func (g groupDo) WithContext(ctx context.Context) *groupDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g groupDo) Clauses(conds ...clause.Expression) *groupDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g groupDo) Not(conds ...gen.Condition) *groupDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g groupDo) Or(conds ...gen.Condition) *groupDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g groupDo) Select(conds ...field.Expr) *groupDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g groupDo) Where(conds ...gen.Condition) *groupDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g groupDo) Order(conds ...field.Expr) *groupDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g groupDo) Distinct(cols ...field.Expr) *groupDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g groupDo) Omit(cols ...field.Expr) *groupDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g groupDo) Join(table schema.Tabler, on ...field.Expr) *groupDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g groupDo) LeftJoin(table schema.Tabler, on ...field.Expr) *groupDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g groupDo) RightJoin(table schema.Tabler, on ...field.Expr) *groupDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g groupDo) Group(cols ...field.Expr) *groupDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g groupDo) Having(conds ...gen.Condition) *groupDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g groupDo) Limit(limit int) *groupDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g groupDo) Offset(offset int) *groupDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g groupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *groupDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g groupDo) Unscoped() *groupDo {
	return g.withDO(g.DO.Unscoped())
}

func (g groupDo) Create(values ...*model.Group) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g groupDo) CreateInBatches(values []*model.Group, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g groupDo) Save(values ...*model.Group) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g groupDo) First() (*model.Group, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Group), nil
	}
}

func (g groupDo) Take() (*model.Group, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Group), nil
	}
}

func (g groupDo) Last() (*model.Group, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Group), nil
	}
}

func (g groupDo) Find() ([]*model.Group, error) {
	result, err := g.DO.Find()
	return result.([]*model.Group), err
}

func (g groupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Group, err error) {
	buf := make([]*model.Group, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g groupDo) FindInBatches(result *[]*model.Group, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g groupDo) Attrs(attrs ...field.AssignExpr) *groupDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g groupDo) Assign(attrs ...field.AssignExpr) *groupDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g groupDo) Joins(field field.RelationField) *groupDo {
	return g.withDO(g.DO.Joins(field))
}

func (g groupDo) Preload(field field.RelationField) *groupDo {
	return g.withDO(g.DO.Preload(field))
}

func (g groupDo) FirstOrInit() (*model.Group, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Group), nil
	}
}

func (g groupDo) FirstOrCreate() (*model.Group, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Group), nil
	}
}

func (g groupDo) FindByPage(offset int, limit int) (result []*model.Group, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g groupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *groupDo) withDO(do gen.Dao) *groupDo {
	g.DO = *do.(*gen.DO)
	return g
}
