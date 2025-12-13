package main

import "test/cmd"

func main() {
	cmd.Serve()
	// ======================== Base64 ========================
	// dummyStr := "Nadim"
	// fmt.Println(dummyStr)

	// dummyByteArr := []byte(dummyStr)
	// fmt.Println(dummyByteArr)

	// enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	// dummyBase64Str := enc.EncodeToString(dummyByteArr)
	// fmt.Println(dummyBase64Str)

	// decodedByteArr, err := enc.DecodeString(dummyBase64Str)
	// if err != nil {
	// 	println(err)
	// }
	// fmt.Println(decodedByteArr)

	// ======================== SHA-256 ========================
	// data := []byte("Hello")
	// hash := sha256.Sum256(data)

	// fmt.Println("Hash:", hash)

	// ======================== HMAC-SHA-256 (HS-256) ========================
	// message := []byte("Hi Nadim.")
	// secret := []byte("my-secret")

	// h := hmac.New(sha256.New, secret)
	// h.Write(message)

	// text := h.Sum(nil)
	// fmt.Println(text)

	// ======================== JWT ========================
	// jwt, err := utils.CreateJwt("my-secret", utils.Payload{
	// 	ID:          1,
	// 	FirstName:   "Nadim",
	// 	LastName:    "Chowdhury",
	// 	Email:       "nadim@gmail.com",
	// 	IsShopOwner: true,
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(jwt)
}
