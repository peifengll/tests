// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/peifengll/tests/testgorm/testgen/gendemo/dal/model"
)

func newFloder(db *gorm.DB, opts ...gen.DOOption) floder {
	_floder := floder{}

	_floder.floderDo.UseDB(db, opts...)
	_floder.floderDo.UseModel(&model.Floder{})

	tableName := _floder.floderDo.TableName()
	_floder.ALL = field.NewAsterisk(tableName)
	_floder.ID = field.NewInt64(tableName, "id")
	_floder.CreatedAt = field.NewTime(tableName, "created_at")
	_floder.UpdatedAt = field.NewTime(tableName, "updated_at")
	_floder.DeletedAt = field.NewField(tableName, "deleted_at")
	_floder.Name = field.NewString(tableName, "name")
	_floder.UserID = field.NewInt64(tableName, "user_id")
	_floder.Decknum = field.NewInt64(tableName, "decknum")

	_floder.fillFieldMap()

	return _floder
}

type floder struct {
	floderDo floderDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	UserID    field.Int64
	Decknum   field.Int64

	fieldMap map[string]field.Expr
}

func (f floder) Table(newTableName string) *floder {
	f.floderDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f floder) As(alias string) *floder {
	f.floderDo.DO = *(f.floderDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *floder) updateTableName(table string) *floder {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")
	f.Name = field.NewString(table, "name")
	f.UserID = field.NewInt64(table, "user_id")
	f.Decknum = field.NewInt64(table, "decknum")

	f.fillFieldMap()

	return f
}

func (f *floder) WithContext(ctx context.Context) *floderDo { return f.floderDo.WithContext(ctx) }

func (f floder) TableName() string { return f.floderDo.TableName() }

func (f floder) Alias() string { return f.floderDo.Alias() }

func (f floder) Columns(cols ...field.Expr) gen.Columns { return f.floderDo.Columns(cols...) }

func (f *floder) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *floder) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 7)
	f.fieldMap["id"] = f.ID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
	f.fieldMap["name"] = f.Name
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["decknum"] = f.Decknum
}

func (f floder) clone(db *gorm.DB) floder {
	f.floderDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f floder) replaceDB(db *gorm.DB) floder {
	f.floderDo.ReplaceDB(db)
	return f
}

type floderDo struct{ gen.DO }

// sql(select * from @@table where user_id = @UserID)
func (f floderDo) FindByUserId(UserID int64) (result []model.Floder, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, UserID)
	generateSQL.WriteString("select * from floder where user_id = ? ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(delete from @@table where id = @id)
func (f floderDo) DelOneByID(id int64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("delete from floder where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (f floderDo) Debug() *floderDo {
	return f.withDO(f.DO.Debug())
}

func (f floderDo) WithContext(ctx context.Context) *floderDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f floderDo) ReadDB() *floderDo {
	return f.Clauses(dbresolver.Read)
}

func (f floderDo) WriteDB() *floderDo {
	return f.Clauses(dbresolver.Write)
}

func (f floderDo) Session(config *gorm.Session) *floderDo {
	return f.withDO(f.DO.Session(config))
}

func (f floderDo) Clauses(conds ...clause.Expression) *floderDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f floderDo) Returning(value interface{}, columns ...string) *floderDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f floderDo) Not(conds ...gen.Condition) *floderDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f floderDo) Or(conds ...gen.Condition) *floderDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f floderDo) Select(conds ...field.Expr) *floderDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f floderDo) Where(conds ...gen.Condition) *floderDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f floderDo) Order(conds ...field.Expr) *floderDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f floderDo) Distinct(cols ...field.Expr) *floderDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f floderDo) Omit(cols ...field.Expr) *floderDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f floderDo) Join(table schema.Tabler, on ...field.Expr) *floderDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f floderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *floderDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f floderDo) RightJoin(table schema.Tabler, on ...field.Expr) *floderDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f floderDo) Group(cols ...field.Expr) *floderDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f floderDo) Having(conds ...gen.Condition) *floderDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f floderDo) Limit(limit int) *floderDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f floderDo) Offset(offset int) *floderDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f floderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *floderDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f floderDo) Unscoped() *floderDo {
	return f.withDO(f.DO.Unscoped())
}

func (f floderDo) Create(values ...*model.Floder) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f floderDo) CreateInBatches(values []*model.Floder, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f floderDo) Save(values ...*model.Floder) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f floderDo) First() (*model.Floder, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Floder), nil
	}
}

func (f floderDo) Take() (*model.Floder, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Floder), nil
	}
}

func (f floderDo) Last() (*model.Floder, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Floder), nil
	}
}

func (f floderDo) Find() ([]*model.Floder, error) {
	result, err := f.DO.Find()
	return result.([]*model.Floder), err
}

func (f floderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Floder, err error) {
	buf := make([]*model.Floder, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f floderDo) FindInBatches(result *[]*model.Floder, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f floderDo) Attrs(attrs ...field.AssignExpr) *floderDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f floderDo) Assign(attrs ...field.AssignExpr) *floderDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f floderDo) Joins(fields ...field.RelationField) *floderDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f floderDo) Preload(fields ...field.RelationField) *floderDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f floderDo) FirstOrInit() (*model.Floder, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Floder), nil
	}
}

func (f floderDo) FirstOrCreate() (*model.Floder, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Floder), nil
	}
}

func (f floderDo) FindByPage(offset int, limit int) (result []*model.Floder, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f floderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f floderDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f floderDo) Delete(models ...*model.Floder) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *floderDo) withDO(do gen.Dao) *floderDo {
	f.DO = *do.(*gen.DO)
	return f
}