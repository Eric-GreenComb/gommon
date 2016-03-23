// cr6868.com post sms api
package sms

import (
	"crypto/tls"
	"log"
	"net"
	"net/url"
	"strings"
)

type SmsApiBean struct {
	Name    string
	Pwd     string
	Content string
	Mobile  string
	Sign    string
	Extno   string
}

func (self *SmsApiBean) Server(name, pwd, content, mobile, sign, extno string) bool {
	if len(name) == 0 || len(pwd) == 0 || len(content) == 0 || len(mobile) == 0 || len(sign) == 0 || len(extno) == 0 {
		return false
	}
	self.Name = name
	self.Pwd = pwd
	self.Content = content
	self.Mobile = mobile
	self.Sign = sign
	self.Extno = extno
	return true
}

//sms_api_url = "http://web.cr6868.com/asmx/smsservice.aspx"
func (self *SmsApiBean) SendSms(sms_api_url string) (bool, error) {
	_response, _err := self.PostSms(sms_api_url, form)

	if _err != nil {
		return false, _err
	}
	_s := strings.Split(_response, ",")

	if len(_s) != 6 {
		return false, nil
	}
	code = _s[0]
	sendid = _s[1]
	invalidcount = _s[2]
	successcount = _s[3]
	blackcount = _s[4]
	msg = _s[5]

	if code == "0" {
		return true, nil
	}

	return false, nil
}

func (self *SmsApiBean) PostSms() (string, error) {
	form := make(url.Values)
	form.Set("name", self.Name)
	form.Set("pwd", self.Pwd)
	form.Set("content", self.Content)
	form.Set("mobile", self.Mobile)
	form.Set("sign", self.Sign)
	form.Set("type", "pt")
	form.Set("extno", self.Extno)
	_response, _err := PostForm(urlStr, form)

	if _err != nil {
		return "", _err
	}

	return _response, nil
}
