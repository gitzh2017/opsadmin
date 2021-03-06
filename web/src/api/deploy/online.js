import service from '@/utils/request'
// @Tags Deploy_Online
// @Summary 分页获取提测列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取提测列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/online/onlineList [post]
// {
//  page     int
//	pageSize int
// }
export const onlineList = (data) => {
    return service({
        url: "/deploy/online/onlineList",
        method: 'post',
        data
    })
}

// @Tags Deploy_Online
// @Summary 文件对比
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "文件对比"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"文件对比成功"}"
// @Router /deploy/online/onlineContrast [post]
export const onlineContrast = (data) => {
    return service({
        url: "/deploy/online/onlineContrast",
        method: 'post',
        data
    })
}

// @Tags Deploy_Online
// @Summary 项目上线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "项目提侧"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"项目提侧成功"}"
// @Router /deploy/online/onlineCreate [post]
export const onlineCreate = (data) => {
    return service({
        url: "/deploy/online/onlineCreate",
        method: 'post',
        data
    })
}

// @Tags Deploy_Online
// @Summary 项目开发审核
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "项目开发审核"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"项目开发审核成功"}"
// @Router /deploy/online/devAudit [post]
export const devAudit = (data) => {
    return service({
        url: "/deploy/online/devAudit",
        method: 'post',
        data
    })
}

// @Tags Deploy_Online
// @Summary 项目测试审核
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "项目测试审核"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"项目测试审核成功"}"
// @Router /deploy/online/testAudit [post]
export const testAudit = (data) => {
    return service({
        url: "/deploy/online/testAudit",
        method: 'post',
        data
    })
}

// @Tags Deploy_Online
// @Summary 项目运维审核
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "项目运维审核"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"项目运维审核成功"}"
// @Router /deploy/online/opsAudit [post]
export const opsAudit = (data) => {
    return service({
        url: "/deploy/online/opsAudit",
        method: 'post',
        data
    })
}

// @Tags Deploy_Online
// @Summary 项目可回滚版本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "项目可回滚版本"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"项目可回滚版本"}"
// @Router /deploy/online/onlineRversion [post]
export const onlineRversion = (data) => {
    return service({
        url: "/deploy/online/onlineRversion",
        method: 'post',
        data
    })
}

// @Tags Deploy_Online
// @Summary 上线工单关闭
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "上线工单关闭"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"上线工单关闭成功"}"
// @Router /deploy/online/onlineClose [post]
export const onlineClose = (data) => {
    return service({
        url: "/deploy/online/onlineClose",
        method: 'post',
        data
    })
}
