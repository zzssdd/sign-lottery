package main

import (
	"sign-lottery/pkg/constants"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const MysqlDSN = constants.MysqlDSN //mysql连接的dsn

func main() {
	//连接数据库
	db, err := gorm.Open(mysql.Open(MysqlDSN))
	if err != nil {
		panic(err)
	}

	//创建generator实例
	g := gen.NewGenerator(gen.Config{
		OutPath:        "dao/db",             //模型输出地址
		Mode:           gen.WithDefaultQuery, //WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		FieldNullable:  true,                 // 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldCoverable: false,
		/*表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.*/
		FieldSignable:     false, // 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldWithIndexTag: false, // 生成 gorm 标签的字段索引属性
		FieldWithTypeTag:  true,  // 生成 gorm 标签的字段类型属性
	})
	g.UseDB(db)

	// 自定义字段的数据类型
	dataMap := map[string]func(detailType string) (dataType string){
		"varchar": func(detailType string) (dataType string) {
			return "string"
		},
		"bigint": func(detailType string) (dataType string) {
			return "int64"
		},
		"int": func(detailType string) (dataType string) {
			return "int32"
		},
		"double": func(detailType string) (dataType string) {
			return "float64"
		},
		"text": func(detailType string) (dataType string) {
			return "string"
		},
		"tinytext": func(detailType string) (dataType string) {
			return "string"
		},
	}

	g.WithDataTypeMap(dataMap)

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `balance, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	autoUpdateTimeField := gen.FieldGORMTag("updated_time", "column:updated_at;type:int unsigned;autoUpdateTime")
	autoCreateTimeField := gen.FieldGORMTag("created_at", "column:created_at;type:int unsigned;autoCreateTime")
	softDeleteField := gen.FieldType("deleted_at", "soft_delete.DeletedAt")
	fieldOpts := []gen.ModelOpt{jsonField, autoCreateTimeField, autoUpdateTimeField, softDeleteField}
	// 创建全部模型文件, 并覆盖前面创建的同名模型
	g.GenerateAllTable(fieldOpts...)
	g.Execute()
}
