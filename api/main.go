package main

import (
	"fmt"

	webpush "github.com/SherClockHolmes/webpush-go"
)

func main() {
	// Decode subscription
	s := &webpush.Subscription{
		Endpoint: "https://fcm.googleapis.com/fcm/send/cgs5qeDN0UI:APA91bFP9RP9S39LxcFkgAmN-PihsN6dWIFRxITUrAagJmR7vQCZx_YGh3-F7C8OqTt7I0Z_iiaCXgo_Sab788GeezoutR6XBLwIwKxv5uDrqCPHP57FOCmsvTP5UiGYE8egB6au_ua0",
		Keys: webpush.Keys{
			Auth:   "tMURIHkVIwuu-eLQYpIIQA",
			P256dh: "BIy6ncYs7c2ZDe1eVOf5ky9H7ZULrLK3FpPjuuoPDjV43Dzq0pgw6YmAgdHzcBjKi8dpyaTqdpxq9Bt9G0Pujgs",
		},
	}

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Teste mensagem PWA"), s, &webpush.Options{
		VAPIDPublicKey:  "BE4Qe5Od5gOom__MA_DQPPjniZRfXsG1Pd0WZmGM8GbxTTUlXbcUkR4DyaheBjErju_sdi_kNuLZkPyEcPdh2ek",
		VAPIDPrivateKey: "nl1jxaxZRlN0yiItmicjHqK7k1j4JTiCnPWyqfcG6WU",
		TTL:             30,
	})

	if err != nil {
		// TODO: Handle error
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
}
