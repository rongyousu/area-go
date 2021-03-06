package main

import (
	"fmt"
	"github.com/foxiswho/area-go/consts"
	"github.com/foxiswho/area-go/mca.gov"
	"github.com/foxiswho/area-go/stats.gov"
	"os"
)

////// mca stats
var SELECT_TYPE = "stats"

//数据 获取
func GetAreaData() {
	if SELECT_TYPE == "mca" {
		mca_gov.GetAreaData()
	} else {
		//FIX ME
		stats_gov.GetAreaData()
	}
}

//格式化数据
func FormatArea() {
	if SELECT_TYPE == "mca" {
		mca_gov.FormatArea()
	} else {
		//FIX ME
		stats_gov.FormatArea()
	}
}

func saveFile() {
	if SELECT_TYPE == "mca" {
		mca_gov.SaveFile()
	} else {
		//FIX ME
		stats_gov.SaveFile()
	}
}

//扩展数据填充
func FormatExtData() {
	if SELECT_TYPE == "mca" {
		mca_gov.FormatExtData()
	} else {
		//FIX ME
		stats_gov.FormatExtData()
	}
}

func saveSqlFile() {
	if SELECT_TYPE == "mca" {
		mca_gov.SaveSqlFile()
	} else {
		//FIX ME
		stats_gov.SaveSqlFile()
	}
}

func saveCsvFile() {
	if SELECT_TYPE == "mca" {
		mca_gov.SaveCsvFile()
	} else {
		//FIX ME
		stats_gov.SaveCsvFile()
	}
}

func main() {
	consts.APP_PATH, _ = os.Getwd()

	fmt.Println("=======获取数据======")
	GetAreaData()
	fmt.Println("=======获取成功=====")
	fmt.Println("=======格式化数据=====")
	//
	FormatArea()
	fmt.Println("=======格式化 扩展数据=====")
	//
	FormatExtData()
	fmt.Println("=======处理成功=====")
	//
	fmt.Println("===================")
	//
	fmt.Println("=======  写入到文件 " + consts.JSON_FILE + "    ============")
	fmt.Println("=======  路径: $GOPATH/src/github.com/foxiswho/area-go/" + consts.JSON_FILE + "    ============")
	fmt.Println("=======  正在写入中…… ")
	//
	saveFile()
	//
	fmt.Println("=======  写入到文件 " + consts.SQL_FILE + "    ============")
	fmt.Println("=======  路径: $GOPATH/src/github.com/foxiswho/area-go/" + consts.SQL_FILE + "    ============")
	fmt.Println("=======  正在写入中…… ")
	saveSqlFile()

	fmt.Println("=======  写入到文件 " + consts.CSV_FILE + "    ============")
	fmt.Println("=======  路径: $GOPATH/src/github.com/foxiswho/area-go/" + consts.CSV_FILE + "    ============")
	fmt.Println("=======  正在写入中…… ")
	saveCsvFile()

	fmt.Println("===================")
	fmt.Println("=======数据保存成功=======")
}
