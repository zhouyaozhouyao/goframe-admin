// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package users

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table users.
type Entity struct {
    Id           int64  `orm:"id,primary"      json:"id"`             // 主键                                                                
    Uuid         string `orm:"uuid"            json:"uuid"`           // UUID                                                                
    Username     string `orm:"username,unique" json:"username"`       // 登录名/11111                                                        
    Password     string `orm:"password"        json:"password"`       // 密码                                                                
    Salt         string `orm:"salt"            json:"salt"`           // 密码盐                                                              
    RealName     string `orm:"real_name"       json:"real_name"`      // 真实姓名                                                            
    DepartId     int    `orm:"depart_id"       json:"depart_id"`      // 部门/11111/dict                                                     
    UserType     int    `orm:"user_type"       json:"user_type"`      // 类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户  
    Status       int    `orm:"status"          json:"status"`         // 状态                                                                
    Thirdid      string `orm:"thirdid"         json:"thirdid"`        // 第三方ID                                                            
    Endtime      string `orm:"endtime"         json:"endtime"`        // 结束时间                                                            
    Email        string `orm:"email"           json:"email"`          // email                                                               
    Tel          string `orm:"tel"             json:"tel"`            // 手机号                                                              
    Address      string `orm:"address"         json:"address"`        // 地址                                                                
    TitleUrl     string `orm:"title_url"       json:"title_url"`      // 头像地址                                                            
    Remark       string `orm:"remark"          json:"remark"`         // 说明                                                                
    Theme        string `orm:"theme"           json:"theme"`          // 主题                                                                
    BackSiteId   int    `orm:"back_site_id"    json:"back_site_id"`   // 后台选择站点ID                                                      
    CreateSiteId int    `orm:"create_site_id"  json:"create_site_id"` // 创建站点ID                                                          
    ProjectId    int64  `orm:"project_id"      json:"project_id"`     // 项目ID                                                              
    ProjectName  string `orm:"project_name"    json:"project_name"`   // 项目名称                                                            
    Enable       int    `orm:"enable"          json:"enable"`         // 是否启用//radio/1,启用,2,禁用                                       
    UpdateTime   string `orm:"update_time"     json:"update_time"`    // 更新时间                                                            
    UpdateId     int64  `orm:"update_id"       json:"update_id"`      // 更新人                                                              
    CreateTime   string `orm:"create_time"     json:"create_time"`    // 创建时间                                                            
    CreateId     int64  `orm:"create_id"       json:"create_id"`      // 创建者                                                              
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}