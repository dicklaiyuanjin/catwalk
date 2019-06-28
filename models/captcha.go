package models

import (
  "github.com/mojocn/base64Captcha"
)

func CaptchaCreate() (string, string){
  var config = base64Captcha.ConfigCharacter{
    Height:             60,
    Width:              240,
    Mode:               base64Captcha.CaptchaModeNumber,
    ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
    ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
    IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    true,
		IsShowSlimeLine:    true,
		IsShowSineLine:     true,
		CaptchaLen:         6,
  }

  idKey, captcha := base64Captcha.GenerateCaptcha("", config)
  base64string := base64Captcha.CaptchaWriteToBase64Encoding(captcha)
  return idKey, base64string
}

func VerifyCaptcha(idkey, verifyValue string) bool {
  return base64Captcha.VerifyCaptcha(idkey, verifyValue)
}

