package main

import (
	"github.com/gin-gonic/gin"
	sms_verification "github.com/lll-lll-lll-lll/sms-verification"
)

func main() {
	router := gin.Default()

	app := sms_verification.Config{Router: router}

	app.Routes()

	router.Run(":80")
}
