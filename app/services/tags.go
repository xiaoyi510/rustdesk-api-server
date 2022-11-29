package services

import (
	"github.com/beego/beego/v2/client/orm"
	"rustdesk-api-server/app/models"
)

var Tags = new(TagsService)

type TagsService struct {
}

// 批量插入
func (this *TagsService) BatchAdd(uid int32, tags []string) bool {
	if len(tags) == 0 {
		return false
	}

	tagList := []models.Tags{}
	for _, t := range tags {
		tagList = append(tagList, models.Tags{
			Uid: uid,
			Tag: t,
		})
	}

	_, err := orm.NewOrm().InsertMulti(3, tagList)
	if err != nil {
		return false
	}
	return true
}

func (this *TagsService) DeleteAll(uid int32) bool {
	_, err := orm.NewOrm().Raw("delete from rustdesk_tags where uid = ?", uid).Exec()
	if err != nil {
		return false
	}
	return true
}

// 查询用户名下Tag
func (this *TagsService) FindTags(uid int32) []models.Tags {
	ret := []models.Tags{}
	_, err := orm.NewOrm().QueryTable(new(models.Tags)).Filter("uid", uid).All(&ret, "tag")
	if err != nil {
		return nil
	}
	return ret
}
