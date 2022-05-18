package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aiio/core/config"
	"github.com/aiio/core/token"
	"github.com/aiio/core/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// RenewToken 续签token
func (u *AuthService) RenewToken(tokenStr string) (string, error) {
	return token.RenewToken(tokenStr)
}

func (u *AuthService) Login(scenes, identify, certificate string) (tokenStr string, err error) {
	switch scenes {
	case "password":
		auth := UserAuth{}
		err = u.db.Where("scenes =? and identity =?", scenes, identify).
			First(&auth).Error
		if err != nil {
			return "", errors.New("用户名不存在")
		}
		if !strings.EqualFold(auth.Certificate, utils.MD5Sum(utils.MD5Sum(certificate))) {
			return "", errors.New("密码错误")
		}
		return token.GenerateToken(auth.UID, "user")
	case "email":
		auth := UserAuth{}
		err = u.db.Where("scenes =? and identity =?", scenes, identify).
			First(&auth).Error
		if err != nil {
			return "", errors.New("用户名不存在")
		}
		if !strings.EqualFold(auth.Certificate, utils.MD5Sum(utils.MD5Sum(certificate))) {
			return "", errors.New("密码错误")
		}
		return token.GenerateToken(auth.UID, "user")
	case "mini_program":
		appid := config.V.MiniParam.AppId
		secret := config.V.MiniParam.AppSecret
		url := "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
		httpGet, err := utils.HTTPGet(fmt.Sprintf(url, appid, secret, certificate))
		if err != nil {
			return "", err
		}
		var result map[string]string
		_ = json.Unmarshal(httpGet, &result)
		if openid, ok := result["openid"]; ok {
			auth := UserAuth{}
			err = u.db.Where("scenes =? and identity =? and certificate =?", scenes, appid, openid).
				First(&auth).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				auth.UID = strings.Join(strings.Split(uuid.New().String(), "-"), "")
				auth.Scenes = scenes
				auth.Identity = appid
				auth.Certificate = openid
				err := u.db.Save(&auth).Error
				if err != nil {
					return "", errors.New("写入数据失败")
				}
			} else if err != nil {
				return "", errors.New("用户名不存在")
			}
			return token.GenerateToken(auth.UID, "user")
		}
		return "", errors.New(result["errmsg"])
	default:
		return "", errors.New("scenes error")
	}
}

func (u *AuthService) Registry(scenes, identify, certificate string) (tokenStr string, err error) {
	switch scenes {
	case "password":
		auth := UserAuth{}
		err = u.db.Where("scenes =? and identity =?", scenes, identify).
			First(&auth).Error
		if err == nil {
			return "", errors.New("用户名存在")
		}
		auth.UID = strings.Join(strings.Split(uuid.New().String(), "-"), "")
		auth.Scenes = scenes
		auth.Identity = identify
		auth.Certificate = utils.MD5Sum(utils.MD5Sum(certificate))
		err := u.db.Create(&auth).Error
		if err != nil {
			return "", errors.New("保存用户数据失败")
		}
		return token.GenerateToken(auth.UID, "user")
	case "mini_program":
		// 感送
		appid := "wx42f6ab953554d9f4"
		secret := "7a5b92278c5720f39be86870924a3a6d"
		url := "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
		httpGet, err := utils.HTTPGet(fmt.Sprintf(url, appid, secret, identify))
		if err != nil {
			return "", err
		}
		var result map[string]interface{}
		_ = json.Unmarshal(httpGet, &result)
		auth := UserAuth{}
		err = u.db.Where("scenes =? and identity =?", scenes, result["openid"]).
			First(&auth).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			auth.UID = strings.Join(strings.Split(uuid.New().String(), "-"), "")
			auth.Scenes = scenes
			auth.Identity = result["openid"].(string)
			err := u.db.Save(&auth).Error
			if err != nil {
				return "", errors.New("写入数据失败")
			}
		} else if err != nil {
			return "", errors.New("用户名不存在")
		}
		return token.GenerateToken(auth.UID, "user")
	//case "verification_code":
	default:
		return "", errors.New("scenes error")
	}
}

//func (u *AuthService) Traefik(host, uri, method, authorization string) (uid string, err error) {
//	if index := strings.Index(uri, "?"); index != -1 {
//		uri = uri[0:index]
//	}
//	uri = strings.TrimPrefix(uri, "/")
//	rule := AuthRule{}
//	err = u.db.Where("host =? and uri =? and method =?", host, uri, method).First(&rule).Error
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		rule.Host = host
//		rule.Uri = uri
//		rule.Method = method
//		u.db.Create(&rule)
//	}
//	return "1000", nil
//}
