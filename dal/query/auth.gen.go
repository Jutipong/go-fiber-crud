// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"fiber-crud/dal/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newAuth(db *gorm.DB) auth {
	_auth := auth{}

	_auth.authDo.UseDB(db)
	_auth.authDo.UseModel(&model.Auth{})

	tableName := _auth.authDo.TableName()
	_auth.ALL = field.NewField(tableName, "*")
	_auth.ID = field.NewString(tableName, "Id")
	_auth.UserName = field.NewString(tableName, "UserName")
	_auth.Password = field.NewString(tableName, "Password")
	_auth.CreateDate = field.NewTime(tableName, "CreateDate")
	_auth.CreateBy = field.NewString(tableName, "CreateBy")
	_auth.UpdateDate = field.NewTime(tableName, "UpdateDate")
	_auth.UpdateBy = field.NewString(tableName, "UpdateBy")
	_auth.IsActive = field.NewBool(tableName, "IsActive")
	_auth.UserID = field.NewString(tableName, "UserId")
	_auth.Amount = field.NewField(tableName, "Amount")

	_auth.fillFieldMap()

	return _auth
}

type auth struct {
	authDo

	ALL        field.Field
	ID         field.String
	UserName   field.String
	Password   field.String
	CreateDate field.Time
	CreateBy   field.String
	UpdateDate field.Time
	UpdateBy   field.String
	IsActive   field.Bool
	UserID     field.String
	Amount     field.Field

	fieldMap map[string]field.Expr
}

func (a auth) Table(newTableName string) *auth {
	a.authDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a auth) As(alias string) *auth {
	a.authDo.DO = *(a.authDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *auth) updateTableName(table string) *auth {
	a.ALL = field.NewField(table, "*")
	a.ID = field.NewString(table, "Id")
	a.UserName = field.NewString(table, "UserName")
	a.Password = field.NewString(table, "Password")
	a.CreateDate = field.NewTime(table, "CreateDate")
	a.CreateBy = field.NewString(table, "CreateBy")
	a.UpdateDate = field.NewTime(table, "UpdateDate")
	a.UpdateBy = field.NewString(table, "UpdateBy")
	a.IsActive = field.NewBool(table, "IsActive")
	a.UserID = field.NewString(table, "UserId")
	a.Amount = field.NewField(table, "Amount")

	a.fillFieldMap()

	return a
}

func (a *auth) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *auth) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["Id"] = a.ID
	a.fieldMap["UserName"] = a.UserName
	a.fieldMap["Password"] = a.Password
	a.fieldMap["CreateDate"] = a.CreateDate
	a.fieldMap["CreateBy"] = a.CreateBy
	a.fieldMap["UpdateDate"] = a.UpdateDate
	a.fieldMap["UpdateBy"] = a.UpdateBy
	a.fieldMap["IsActive"] = a.IsActive
	a.fieldMap["UserId"] = a.UserID
	a.fieldMap["Amount"] = a.Amount
}

func (a auth) clone(db *gorm.DB) auth {
	a.authDo.ReplaceDB(db)
	return a
}

type authDo struct{ gen.DO }

func (a authDo) Debug() *authDo {
	return a.withDO(a.DO.Debug())
}

func (a authDo) WithContext(ctx context.Context) *authDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authDo) Clauses(conds ...clause.Expression) *authDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authDo) Returning(value interface{}, columns ...string) *authDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authDo) Not(conds ...gen.Condition) *authDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authDo) Or(conds ...gen.Condition) *authDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authDo) Select(conds ...field.Expr) *authDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authDo) Where(conds ...gen.Condition) *authDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *authDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a authDo) Order(conds ...field.Expr) *authDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authDo) Distinct(cols ...field.Expr) *authDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authDo) Omit(cols ...field.Expr) *authDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authDo) Join(table schema.Tabler, on ...field.Expr) *authDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authDo) LeftJoin(table schema.Tabler, on ...field.Expr) *authDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authDo) RightJoin(table schema.Tabler, on ...field.Expr) *authDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authDo) Group(cols ...field.Expr) *authDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authDo) Having(conds ...gen.Condition) *authDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authDo) Limit(limit int) *authDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authDo) Offset(offset int) *authDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *authDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authDo) Unscoped() *authDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authDo) Create(values ...*model.Auth) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authDo) CreateInBatches(values []*model.Auth, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authDo) Save(values ...*model.Auth) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authDo) First() (*model.Auth, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) Take() (*model.Auth, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) Last() (*model.Auth, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) Find() ([]*model.Auth, error) {
	result, err := a.DO.Find()
	return result.([]*model.Auth), err
}

func (a authDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Auth, err error) {
	buf := make([]*model.Auth, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authDo) FindInBatches(result *[]*model.Auth, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authDo) Attrs(attrs ...field.AssignExpr) *authDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authDo) Assign(attrs ...field.AssignExpr) *authDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authDo) Joins(field field.RelationField) *authDo {
	return a.withDO(a.DO.Joins(field))
}

func (a authDo) Preload(field field.RelationField) *authDo {
	return a.withDO(a.DO.Preload(field))
}

func (a authDo) FirstOrInit() (*model.Auth, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) FirstOrCreate() (*model.Auth, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Auth), nil
	}
}

func (a authDo) FindByPage(offset int, limit int) (result []*model.Auth, count int64, err error) {
	if limit <= 0 {
		count, err = a.Count()
		return
	}

	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a authDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a *authDo) withDO(do gen.Dao) *authDo {
	a.DO = *do.(*gen.DO)
	return a
}
