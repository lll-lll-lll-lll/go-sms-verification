# go-sms-verification

# how to use
```sh
git clone https://github.com/lll-lll-lll-lll/go-sms-verification.git \
&& cd go-sms-verification 

```
## set up env
ref: [get env](https://dev.to/hackmamba/how-to-build-a-one-time-passwordotp-verification-api-with-go-and-twilio-3363)
```txt

TWILIO_ACCOUNT_SID=
TWILIO_AUTHTOKEN=
TWILIO_SERVICES_ID=

```


```sh
go run cmd/main.go
```

# Module 
- gin
- twilio-go
- godotenv
- go-playground/validator/v10-> struct, fieldsを検証


# Example Request
![](./example.png)


# Gif

![](./RPReplay_Final1667120122_MP4_AdobeExpress.gif)


# Ref
- https://dev.to/hackmamba/how-to-build-a-one-time-passwordotp-verification-api-with-go-and-twilio-3363
- https://www.twilio.com/ja/
