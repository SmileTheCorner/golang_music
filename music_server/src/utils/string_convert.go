package utils

import (
	"fmt"
	"strings"
)

func LowerCaseFirst(str string) string {
	if str == "" {
		return str
	}
	return fmt.Sprintf("%s%s", strings.ToLower(str[0:1]), str[1:])
}

func UpperCaseFirst(str string) string {
	if str == "" {
		return str
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(str[0:1]), str[1:])
}

// 去掉字符串中的下划线，并将下划线后的第一个字母大写
func RemoveUnderline(str string) string {
	split := strings.Split(str, "_")
	var builder strings.Builder
	for i := range split {
		builder.WriteString(fmt.Sprintf("%s%s", strings.ToUpper(split[i][0:1]), split[i][1:]))
	}
	return fmt.Sprintf("%s%s", strings.ToLower(builder.String()[0:1]), builder.String()[1:])
}

// 将mysql中的属性类型转成golang中对应的类型
func MysqlTypeToGoType(mysqlType string) string {
	switch mysqlType {
	case "varchar", "text", "char", "longtext", "mediumtext", "tinytext", "enum", "set":
		return "string"
	case "int", "tinyint", "smallint", "mediumint", "bigint":
		return "int"
	case "float", "double", "decimal":
		return "float64"
	case "date", "time", "datetime", "timestamp":
		return "time.Time"
	default:
		return "string"
	}
}

// 将mysql中的属性类型转成Ts中对应的类型
func MysqlTypeToTsType(mysqlType string) string {
	switch mysqlType {
	case "varchar", "text", "char", "longtext", "mediumtext", "tinytext", "enum", "set":
		return "string"
	case "int", "tinyint", "smallint", "mediumint", "bigint":
		return "number"
	case "float", "double", "decimal":
		return "number"
	case "date", "time", "datetime", "timestamp":
		return "Date"
	default:
		return "string"
	}
}
