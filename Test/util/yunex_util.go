
package util

import (
	"fmt"
	"github.com/jie123108/glog"
	"math/rand"
	"time"
)

const (
	letters string ="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	secret string = "abcdefghijklmn"
	EXPIRE_RES_INFO time.Duration = 10 * time.Second
)

func RandomSample(letters string, n int) string {
	b := make([]byte, n)
	llen := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(llen)]
	}
	return string(b)
}


type sthotOrder struct {
	ReasonSt
	OrderId string `json:"order_id"`
	UserId int64 `json:"user_id"`
}


type YunexAccount struct {
	UserId int64 `json:"user_id,string"`
	Useable float64 `json:"useable,string"`
	Total float64 `json:"total,string"`	
	Freeze float64 `json:"freeze,string"`
}

type snapAccount struct {
	Snaps []YunexAccount `json:"snaps"`
	Total int64 `json:"total,string"`
	Symbol string `json:"symbol"`
}
type ReasonSt struct {
	Reason string `json:"reason"`
}
type addrYunex struct {
	ReasonSt
	UserId int64 `json:"user_id" `
}

type YunexWithDraw struct {
	Coin string `json:"coin"`
	Address string `json:"address"`
	OrderId string `json:"order_id,string"`	
	Amount float64 `json:"amount,string"`	
}

type yunexAccount struct {
	UserId int64 `json:"user_id,string"`
	Address string `json:"string"`
}
type yunexUserInfo struct {
	Plat string `json:"plat"`
	Zonenum string `json:"zone_num"`
	Mobile string `json:"mobile"`
	Symbol string `json:"symbol"`
}
// 绑定并获取yunex address
func GetYunexAddressByTel(zone, tel string)(addr string, err error) {
	uri := "/api/pub/user/deposit/bind/"
	var ret yunexAccount

	v := yunexUserInfo{Plat:"yunbay", Zonenum:zone, Mobile:tel, Symbol:"KT"}
	err = post_yunex(uri, "yunex", nil, v, "", &ret, false, EXPIRE_RES_INFO)
	if err != nil {
		glog.Error("IsYunexAddress fail! err=", err)
		return
	}

	addr = ret.Address
	return
}

// 判断是否为yunex地址接口
func IsYunexAddress(address string, ) (user_id int64, err error){
	uri := fmt.Sprintf("/api/yunex/user/address/query/?coin=KT&address=%v", address)
	var ret addrYunex
	err = get_yunex(uri, "yunex", "", &ret, false, EXPIRE_RES_INFO)
	if err != nil {
		if ret.Reason != "" {
			err = fmt.Errorf(ret.Reason)			
		}
		glog.Error("IsYunexAddress fail! err=", err)
		return
	}

	user_id = ret.UserId
	return
}

// 帐号提币接口
func YunexWithdrawWallet(v YunexWithDraw) (order_id string, err error){
	uri := "/api/platform/user/deposit/"
	var ret sthotOrder
	err = post_yunex(uri, "yunex", nil, v, "", &ret, false, EXPIRE_RES_INFO)
	if err != nil {
		if ret.Reason != "" {
			err = fmt.Errorf(ret.Reason)			
		}
		glog.Error("YunexWithdrawWallet fail! err=", err, " args:", v)		
		return
	}
	order_id = ret.OrderId
	return
}


// 获取云网持有ybt的用户
func SnapYunexYbtAccount(start, count int, day string) (total int64, vs []YunexAccount, err error) {
	uri := "http://a.yunex.io/api/coin/bonus/snap?" + fmt.Sprintf("start=%v&count=%v&day=%v", start, count, day)
	var ret snapAccount
	if err = get_yunex(uri, "yunex", "", &ret, false, EXPIRE_RES_INFO); err != nil {
		glog.Error("SnapYunexYbtAccount fail! err=", err)
		return
	}
	total = ret.Total
	vs = ret.Snaps
	return
}


// 获取yunex平台帐户余额信息
func GetYunbayAccountInYunex() (m map[string]float64, err error) {
	uri := "http://yunex.io/api/yunex/credit/query/"	
	if err = get_yunex(uri, "yunex", "", &m, false, EXPIRE_RES_INFO); err != nil {
		glog.Error("SnapYunexYbtAccount fail! err=", err)
		return
	}
	return
}

type YunexKtBonus struct {
	ToUid int64 `json:"to_uid"`
	Symbol string `json:"symbol"`
	Amount float64 `json:"amount,string"`
	Date string `json:"date"`
	OrderId int64 `json:"order_id,string"`	
}

type YunexKtBonusRet struct {
	OrderId int64 `json:"order_id,string"`
	Reason int `json:"reason"`
}

type bonusYunex struct {
	Bonus []YunexKtBonus `json:"bonus"`
}



// kt分红转帐给云网用户
func BonusYunexKt(vs []YunexKtBonus) (fail []YunexKtBonusRet, err error) {
	uri := "http://a.yunex.io/api/coin/bonus/transfer"
	v := bonusYunex{Bonus:vs}
	if err = post_yunex(uri, "yunex", nil, v, "fail", &fail, false, EXPIRE_RES_INFO); err != nil {
		glog.Error("BonusYunexKt fail! err=", err)
		return
	}
	return
}

type YunexYbtDeposit struct {
	Plat string `json:"plat"`
	Address string `json:"address"`
	Coin string `json:"coin"`
	Amount float64 `json:"amount,string"`
    OrderId int `json:"order_id,string"`
}
// 提币接口（yunbay提YBT币到yunex）
func RechargeYunExYBt(data YunexYbtDeposit) (dataStatus ReasonSt, err error) {
	uri := "http://a.yunex.io/api/pub/user/deposit/"
	err = RequestYunExApi(uri, "POST", nil, data, "fail", &dataStatus, false, EXPIRE_RES_INFO)
	if err != nil {
		glog.Error("/api/pub/user/deposit/ is fail, err:", err)
	}
    return
}
// 查询yunbay在yunex中的账号余额
func GetYunExBalance() (dataStatus ReasonSt, err error) {
	uri := "http://a.yunex.io/api/pub/account/balance/?plat=yunbay"
	err = RequestYunExApi(uri, "GET", nil, nil, "fail", &dataStatus, false, EXPIRE_RES_INFO)
	if err != nil {
		glog.Error("/api/pub/account/balance/ is fail, err:", err)
	}
    return
}