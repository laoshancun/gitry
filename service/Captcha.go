package service

import (
	"errors"

	"github.com/mojocn/base64Captcha"
)

//Captcha service
type Captcha struct {
}

var config = base64Captcha.ConfigCharacter{
	Height: 60,
	Width:  240,
	//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
	Mode:               base64Captcha.CaptchaModeNumberAlphabet,
	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
	IsUseSimpleFont:    true,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     true,
	IsShowNoiseText:    true,
	IsShowSlimeLine:    true,
	IsShowSineLine:     false,
	CaptchaLen:         6,
}

//GenerateCaptcha for generate a new captcha
func (captcha *Captcha) GenerateCaptcha(id string) (string, string) {
	idkey, cap := base64Captcha.GenerateCaptcha(id, config)
	//以 base64 编码
	base64string := base64Captcha.CaptchaWriteToBase64Encoding(cap)
	return idkey, base64string
}

//VerfiyCaptcha method
func (captcha *Captcha) VerfiyCaptcha(id string, verifyValue string) (bool, error) {
	verifyResult := base64Captcha.VerifyCaptcha(id, verifyValue)
	if verifyResult {
		return true, nil
	} else {
		return false, errors.New("captcha verfiy faild")
	}
}
