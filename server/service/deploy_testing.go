package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strconv"
	"strings"
)

// @title    TestingList
// @description   get test order list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func TestingList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.DeployTesting{})
	var testingList []model.DeployTesting
	err = db.Count(&total).Error
	err = db.Order("id DESC ").Preload("DeployProject.Environment").Preload("DeployProject.Server").Limit(limit).Offset(offset).Find(&testingList).Error
	return err, testingList, total
}

// @title    TestingContrast
// @description   对比文件
// @auth                      （2020/07/15  09:50）
// @param     info             request.ContrastInfo
// @return    err              error
// @return    list             interface{}
// @return    path             string

func TestingContrast(testting request.ContrastInfo) (err error, list interface{}, path string) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", testting.DeployProjectId).Preload("Server").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err)), list, path
	}

	filepath := fmt.Sprintf("%d-%s-%.1f", project.ID, project.Name, project.ReleaseVersion+0.1)
	path, err = utils.Gitpull(testting.Tag, project.GitUrl, filepath)
	if err != nil {
		return errors.New(fmt.Sprint("Git拉取项目报错, 报错信息: %s", err)), list, path
	}

	exclude := strings.Fields(project.IgnoreFiles)

	var result []map[string]string
	var errMsg string
	for _, value := range project.Server {
		err, ret := utils.FileContrast(path, value.User, value.Host, value.Port, project.Directory, exclude)
		if err != nil {
			errMsg += fmt.Sprint("对比文件报错, 报错信息: %s", err)
		}
		result = append(result, ret...)
	}

	return err, result, path
}

// @title    TestingRelease
// @description   提交并同步要发布的文件，并清理过期备份
// @auth                      （2020/07/17  19:45）
// @param     info             request.TestingReleaseInfo
// @return    err              error

func TestingRelease(testting request.TestingReleaseInfo, username *request.CustomClaims) (err error) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", testting.DeployProjectId).Preload("Server").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err))
	}
	go func() {
		version, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", project.ReleaseVersion+0.1), 64)
		testOrder := &model.DeployTesting{
			Applicant:       username.NickName,
			Tag:             testting.Tag,
			DeployProjectId: testting.DeployProjectId,
			Describe:        testting.Describe,
			Path:            testting.Path,
			Version:         version,
			Isdelete:        1,
		}

		err = global.GVA_DB.Create(testOrder).Error
		result := ""
		exclude := strings.Fields(project.IgnoreFiles)

		for _, value := range project.Server {
			result += fmt.Sprintf("==========================主机开始同步文件: %s=========================\n", value.Host)
			err, ret := utils.FileSync(testting.Path, value.User, value.Host, value.Port, project.Directory, exclude)
			if err != nil {
				result += fmt.Sprintf("同步报错: %s", err)
			}
			result += ret
		}

		if strings.HasPrefix(result, "同步报错") {
			err = TestingUpdate(testOrder.ID, 2, result)
		} else {
			err = TestingUpdate(testOrder.ID, 1, result)
			project.ReleaseVersion = version
			err = ProjectStatusUpdate(project)
		}

		// 删除过期备份
		_ = TestingVersionDelete(float64(project.ID))
	}()

	return err
}

// @title    TestingUpdate
// @description    更新上线工单
// @auth                     （2020/04/05  20:22）
// @param     env             model.DeployTesting
// @return                    error

func TestingUpdate(id uint, status int, result string) (err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&model.DeployTesting{}).Updates(&model.DeployTesting{Status: status, Result: result}).Error
	return err
}

// @title   ProjectRversion
// @description    可回滚版本
// @auth                     （2020/07/10  15:11）
// @param     id
// @return    err             error

func TestingRversion(id float64) (err error, list interface{}) {
	var testing []model.DeployTesting
	err = global.GVA_DB.Where("deploy_project_id = ? and isdelete = 1", id).Group("version").Order("created_at desc").Find(&testing).Error
	return err, testing
}
