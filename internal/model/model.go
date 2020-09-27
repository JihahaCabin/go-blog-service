package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DataBaseSettings) (*gorm.DB, error) {
	/**
	替换字符
	bool: %t
	int, int8 etc.: %d
	uint, uint8 etc.: %d, %#x if printed with %#v
	float32, complex64, etc: %g
	string: %s
	chan: %p
	pointer: %p
	*/
	db, err := gorm.Open(databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.Username,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
		))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true) //建立的表的表名，不为复数
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallBack)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallBack)
	db.Callback().Delete().Replace("gorm:delete", deleteCallBack)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil

}

func updateTimeStampForCreateCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()

		if createTime, ok := scope.FieldByName("CreatedOn"); ok {
			if createTime.IsBlank {
				_ = createTime.Set(nowTime)
			}
		}

		if modifyTime, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTime.IsBlank {
				_ = modifyTime.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallBack(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnFiled, hasDeletedOnFiled := scope.FieldByName("DeletedOn")
		isDelFileld, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnFiled && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v Set %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnFiled.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelFileld.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
