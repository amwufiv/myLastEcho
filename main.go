package main

import (
	"github.com/amwufiv/myLastEcho/util"
	"log"
	"os"
)

const (
	privateKey  = "PRIVATE_KEY"
	updateTime  = "UPDATE_TIME"
	encryptFile = "ENCRYPT_FILE"
)

func main() {
	privateKeyValue := os.Getenv(privateKey)
	updateTimeValue := os.Getenv(updateTime)
	encryptFileValue := os.Getenv(encryptFile)

	isExpiredJustNow := !util.CheckAlive(updateTimeValue, 30)
	needSendMail := !util.CheckAlive(updateTimeValue, 33)
	if isExpiredJustNow && !needSendMail {
		// 刚过期，发送邮件提醒
	} else if needSendMail {
		// 发送邮件
		plaintext := util.Decrypt(encryptFileValue, privateKeyValue)
		if len(plaintext) > 0 {
			log.Println(plaintext)
		}
	} else {
		log.Println("you are still alive")
	}

}
