package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dim13/otpauth/migration"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"time"
)

// 从二维码导入otp令牌
func getMfaCode(key string) string {
	file, err := os.Open("/Users/admin/Downloads/export_mfa.jpeg")
	if err != nil {
		log.Fatalln(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		log.Fatalln(err)
	}
	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)

	if err != nil {
		log.Fatalln(err)
	}

	urlStr, _ := url.QueryUnescape(result.String())
	obj, _ := migration.UnmarshalURL(urlStr)
	for _, o := range obj.OtpParameters {
		if o.GetName() == key {
			return o.EvaluateString()
		}
	}
	return ""
}

func display(key *otp.Key, data []byte) {
	fmt.Printf("Issuer:       %s\n", key.Issuer())
	fmt.Printf("Account Name: %s\n", key.AccountName())
	fmt.Printf("Secret:       %s\n", key.Secret())
	fmt.Println("Writing PNG to qr-code.png....")
	ioutil.WriteFile("qr-code.png", data, 0644)
	fmt.Println("")
	fmt.Println("Please add your TOTP to your OTP Application now!")
	fmt.Println("")
}

// Demo function, not used in main
// Generates Passcode using a UTF-8 (not base32) secret and custom paramters
func GeneratePassCode(secret string) string {
	//secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}
func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

// 创建otp令牌
func createOtp() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Boncloud.com",
		AccountName: "yaoshicheng@bonc.com.cn",
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})

	log.Printf("%#v", key)
	if err != nil {
		panic(err)
	}
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	png.Encode(&buf, img)

	// display the QR code to the user.
	display(key, buf.Bytes())

	// Now Validate that the user's successfully added the passcode.
	fmt.Println("Validating TOTP...")
	fmt.Println(GeneratePassCode(key.Secret()))
	passcode := promptForPasscode()
	valid := totp.Validate(passcode, key.Secret())
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passcode!")
		os.Exit(1)
	}

}
func main() {

	// google export
	str := "otpauth-migration://offline?data=CjkKCmlCfkkKiuOPRvYSHeawtOa7tDpxaXJvbmdqaWUxQGJvbmMuY29tLmNuGgbmsLTmu7QgASgBMAIQARgBIAAoj5vE2vj%2F%2F%2F%2F%2FAQ%3D%3D"
	urlStr, _ := url.QueryUnescape(str)
	obj, _ := migration.UnmarshalURL(urlStr)
	for _, o := range obj.OtpParameters {

		fmt.Printf("key: %s ,code: %s \n", o.GetName(), o.EvaluateString())
		fmt.Printf("otp: %s\n", o.URL().String())
		//o.WriteFile(fmt.Sprintf("/tmp/%s.png", o.GetName()))

	}

	// from mfa get code
	mfaUrl := "otpauth://totp/yaoshicheng@bonc.com.cn?algorithm=SHA1&digits=6&issuer=%E6%B0%B4%E6%BB%B4&period=30&secret=ZLPQVYDC7NC26W3A"

	key, _ := otp.NewKeyFromURL(mfaUrl)
	fmt.Printf("code: %s\n", GeneratePassCode(key.Secret()))

}
