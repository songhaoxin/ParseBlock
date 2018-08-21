package blocklist

import (
	"clmwallet/database/mysqltools"
	"errors"
	"fmt"
)

type BlockNodeInfo struct {
	BlockNumber string		`gorm:"primary_key"`
	BlockHash   string
	ParentHash  string
}

func (BlockNodeInfo)  TableName() string	{
	return "blockNodeInfo"
}

func init() {
	db := mysqltools.GetInstance().GetMysqlDB()
	if !db.HasTable(&BlockNodeInfo{}) {
		db.CreateTable(&BlockNodeInfo{})
		fmt.Printf("不存在blockNodeInfo表")
		fmt.Println("创建了blockNodeInfo表")
	} else {
		fmt.Println("已经存在blockNodeInfo表")
	}
}

/// 持久化到数据库
func (info *BlockNodeInfo) Store()  error {
	db := mysqltools.GetInstance().GetMysqlDB()
	if err := db.Create(info).Error;err != nil {
		return err
	}
	return nil
}

/// 更新到数据库
func (info *BlockNodeInfo) Save() error  {
	db := mysqltools.GetInstance().GetMysqlDB()
	if err := db.Save(info).Error;err != nil {
		return err
	}
	return nil
}

func (info *BlockNodeInfo) Delete() error  {
	if "" == info.BlockNumber {
		return errors.New("Primary key don't allow Empty.")
	}

	db := mysqltools.GetInstance().GetMysqlDB()
	if err := db.Delete(info).Error;err != nil {
		return err
	}
	return nil
}




