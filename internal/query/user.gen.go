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

func newUser(db *gorm.DB) user {
	_user := user{}

	_user.userDo.UseDB(db)
	_user.userDo.UseModel(&model.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewField(tableName, "*")
	_user.ID = field.NewInt64(tableName, "id")
	_user.Nickname = field.NewString(tableName, "nickname")
	_user.Password = field.NewString(tableName, "password")
	_user.Key = field.NewString(tableName, "key")
	_user.Gender = field.NewBool(tableName, "gender")
	_user.AvatarURL = field.NewString(tableName, "avatar_url")
	_user.Extra = field.NewString(tableName, "extra")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")
	_user.Device = userHasManyDevice{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Device", "model.Device"),
	}

	_user.Message = userHasManyMessage{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Message", "model.Device"),
	}

	_user.GroupUser = userManyToManyGroupUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("GroupUser", "model.GroupUser"),
	}

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo userDo

	ALL       field.Field
	ID        field.Int64
	Nickname  field.String
	Password  field.String
	Key       field.String
	Gender    field.Bool
	AvatarURL field.String
	Extra     field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Device    userHasManyDevice

	Message userHasManyMessage

	GroupUser userManyToManyGroupUser

	fieldMap map[string]field.Expr
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))

	u.ALL = field.NewField(alias, "*")
	u.ID = field.NewInt64(alias, "id")
	u.Nickname = field.NewString(alias, "nickname")
	u.Password = field.NewString(alias, "password")
	u.Key = field.NewString(alias, "key")
	u.Gender = field.NewBool(alias, "gender")
	u.AvatarURL = field.NewString(alias, "avatar_url")
	u.Extra = field.NewString(alias, "extra")
	u.CreatedAt = field.NewTime(alias, "created_at")
	u.UpdatedAt = field.NewTime(alias, "updated_at")
	u.DeletedAt = field.NewField(alias, "deleted_at")

	u.fillFieldMap()

	return &u
}

func (u *user) WithContext(ctx context.Context) *userDo { return u.userDo.WithContext(ctx) }

func (u user) TableName() string { return u.userDo.TableName() }

func (u *user) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := u.fieldMap[fieldName]
	return field, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 13)
	u.fieldMap["id"] = u.ID
	u.fieldMap["nickname"] = u.Nickname
	u.fieldMap["password"] = u.Password
	u.fieldMap["key"] = u.Key
	u.fieldMap["gender"] = u.Gender
	u.fieldMap["avatar_url"] = u.AvatarURL
	u.fieldMap["extra"] = u.Extra
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt

}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userHasManyDevice struct {
	db *gorm.DB

	field.RelationField
}

func (a userHasManyDevice) Where(conds ...field.Expr) *userHasManyDevice {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userHasManyDevice) WithContext(ctx context.Context) *userHasManyDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userHasManyDevice) Model(m *model.User) *userHasManyDeviceTx {
	return &userHasManyDeviceTx{a.db.Model(m).Association(a.Name())}
}

type userHasManyDeviceTx struct{ tx *gorm.Association }

func (a userHasManyDeviceTx) Find() (result []*model.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a userHasManyDeviceTx) Append(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userHasManyDeviceTx) Replace(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userHasManyDeviceTx) Delete(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userHasManyDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a userHasManyDeviceTx) Count() int64 {
	return a.tx.Count()
}

type userHasManyMessage struct {
	db *gorm.DB

	field.RelationField
}

func (a userHasManyMessage) Where(conds ...field.Expr) *userHasManyMessage {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userHasManyMessage) WithContext(ctx context.Context) *userHasManyMessage {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userHasManyMessage) Model(m *model.User) *userHasManyMessageTx {
	return &userHasManyMessageTx{a.db.Model(m).Association(a.Name())}
}

type userHasManyMessageTx struct{ tx *gorm.Association }

func (a userHasManyMessageTx) Find() (result []*model.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a userHasManyMessageTx) Append(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userHasManyMessageTx) Replace(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userHasManyMessageTx) Delete(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userHasManyMessageTx) Clear() error {
	return a.tx.Clear()
}

func (a userHasManyMessageTx) Count() int64 {
	return a.tx.Count()
}

type userManyToManyGroupUser struct {
	db *gorm.DB

	field.RelationField
}

func (a userManyToManyGroupUser) Where(conds ...field.Expr) *userManyToManyGroupUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userManyToManyGroupUser) WithContext(ctx context.Context) *userManyToManyGroupUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userManyToManyGroupUser) Model(m *model.User) *userManyToManyGroupUserTx {
	return &userManyToManyGroupUserTx{a.db.Model(m).Association(a.Name())}
}

type userManyToManyGroupUserTx struct{ tx *gorm.Association }

func (a userManyToManyGroupUserTx) Find() (result *model.GroupUser, err error) {
	return result, a.tx.Find(&result)
}

func (a userManyToManyGroupUserTx) Append(values ...*model.GroupUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userManyToManyGroupUserTx) Replace(values ...*model.GroupUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userManyToManyGroupUserTx) Delete(values ...*model.GroupUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userManyToManyGroupUserTx) Clear() error {
	return a.tx.Clear()
}

func (a userManyToManyGroupUserTx) Count() int64 {
	return a.tx.Count()
}

type userDo struct{ gen.DO }

func (u userDo) Debug() *userDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) *userDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) Clauses(conds ...clause.Expression) *userDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Not(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) *userDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) *userDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() *userDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*model.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*model.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Take() (*model.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Last() (*model.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Find() ([]*model.User, error) {
	result, err := u.DO.Find()
	return result.([]*model.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error) {
	buf := make([]*model.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(field field.RelationField) *userDo {
	return u.withDO(u.DO.Joins(field))
}

func (u userDo) Preload(field field.RelationField) *userDo {
	return u.withDO(u.DO.Preload(field))
}

func (u userDo) FirstOrInit() (*model.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FirstOrCreate() (*model.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*model.User, count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	result, err = u.Offset(offset).Limit(limit).Find()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
