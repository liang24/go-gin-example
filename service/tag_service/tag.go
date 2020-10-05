package tag_service

import (
	"encoding/json"
	"io"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/liang24/go-gin-example/models"
	"github.com/liang24/go-gin-example/pkg/export"
	"github.com/liang24/go-gin-example/pkg/gredis"
	"github.com/liang24/go-gin-example/pkg/logging"
	"github.com/liang24/go-gin-example/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)

	cache := cache_service.Tag{
		State: t.State,

		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}

	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}

	file := excelize.NewFile()
	sheetName := "标签信息"
	file.SetSheetName("Sheet1", sheetName)

	cols := []string{"A", "B", "C", "D", "E", "F"}
	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}

	for i, title := range titles {
		file.SetCellValue(sheetName, cols[i]+"1", title)
	}

	for i, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(v.CreatedOn),
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}
		row := strconv.Itoa(i + 2)
		for ii, value := range values {
			file.SetCellValue(sheetName, cols[ii]+row, value)
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + ".xlsx"

	// 根据指定路径保存文件
	fullPath := export.GetExcelFullPath() + filename
	err = file.SaveAs(fullPath)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (t *Tag) Import(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}

			models.AddTag(data[1], 1, data[2])
		}
	}

	return nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
