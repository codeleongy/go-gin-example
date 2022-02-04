package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 获取所有Tag
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		// SELECT * FROM `blog_tag`  LIMIT pageSize OFFSET pageNum
		// 意为从pageNum往后开始找，找pageSize条
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// 获取Tag总数量
func GetTagToTal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Tag{}).Where(maps).Where("deleted_on = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 通过名字判断Tag是否已经存在
/* 	SQL语句:
 		SELECT id FROM `blog_tag`  WHERE (name = 'test')
		ORDER BY `blog_tag`.`id` ASC LIMIT 1
*/
func ExisitTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

// 通过ID判断是否存在Tag
func ExisitTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

// 添加Tag值
/* 	SQL语句：
 	   	INSERT INTO `blog_tag` (`created_on`,`modified_on`,`name`,`created_by`,`modified_by`,`state`)
		VALUES (1643176202,0,'test','leong','',0)
*/
func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

/*
	这属于gorm的Callbacks，可以将回调方法定义为模型结构的指针，
	在创建、更新、查询、删除时将被调用，
	如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。
	gorm所支持的回调方法：
		创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
		更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
		删除：BeforeDelete、AfterDelete
		查询：AfterFind
*/
// 添加数据后给CreatedOn赋值
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// 修改数据后给ModifiedOn赋值
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

// 修改Tag值
/*
	UPDATE `blog_tag` SET `modified_by` = 'leong', `modified_on` = 1643188366, `name` = 'test02'
	WHERE (id = 3)
*/
func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// 删除Tag值
func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}
