package middleware

/*
编写 gtoken 中间件的目的：
全局校验用户的登录状态
登录后的用户将用户名、id 这类用户信息写入到 Context 上下中，方便全局调用
在中间件中统一进行账号判断，比如：是否被拉黑等判断操作
*/

import (
	"shop/library/response"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	CtxAccountId      = "account_id"       //token获取
	CtxAccountName    = "account_name"     //token获取
	CtxAccountAvatar  = "account_avatar"   //token获取
	CtxAccountSex     = "account_sex"      //token获取
	CtxAccountStatus  = "account_status"   //token获取
	CtxAccountSign    = "account_sign"     //token获取
	CtxAccountIsAdmin = "account_is_admin" //token获取
	CtxAccountRoleIds = "account_role_ids" //token获取
)

type TokenInfo struct {
	Id      int
	Name    string
	Avatar  string
	Sex     int
	Status  int
	Sign    string
	RoleIds string
	IsAdmin int
}

var GToken *gtoken.GfToken

var MiddlewareGToken = tokenMiddleware{}

type tokenMiddleware struct{}

func (s *tokenMiddleware) GetToken(r *ghttp.Request) {
	var tokenInfo TokenInfo
	token := GToken.GetTokenData(r)
	err := gconv.Struct(token.GetString("data"), &tokenInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//账号被冻结拉黑
	if tokenInfo.Status == 2 {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(CtxAccountId, tokenInfo.Id)
	r.SetCtxVar(CtxAccountName, tokenInfo.Name)
	r.SetCtxVar(CtxAccountAvatar, tokenInfo.Avatar)
	r.SetCtxVar(CtxAccountSex, tokenInfo.Sex)
	r.SetCtxVar(CtxAccountStatus, tokenInfo.Status)
	r.SetCtxVar(CtxAccountSign, tokenInfo.Sign)
	r.SetCtxVar(CtxAccountRoleIds, tokenInfo.RoleIds)
	r.SetCtxVar(CtxAccountIsAdmin, tokenInfo.Sign)
	r.Middleware.Next()
}
