package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
)

func openFileAppendly(filename string) (io.Writer, func()) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return f, func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}
}

func randomNumBetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

type void struct{}

func randomTimerBetween(a, b int) chan void {
	ch := make(chan void)
	go func() {
		var loop func()
		loop = func() {
			n := randomNumBetween(a, b)
			time.Sleep(time.Duration(n) * time.Second)
			ch <- void{}
			loop()
		}
		loop()
	}()
	return ch
}

type FakeUser struct {
	UUID string `json:"uuid" faker:"uuid_digit"`
	Name string `json:"name" faker:"name"`

	PhoneNumber string `json:"phone_number" faker:"phone_number"`
	Email       string `faker:"email" json:"email"`

	Latitude  float32 `faker:"lat" json:"latitude"`
	Longitude float32 `faker:"long" json:"longitude"`

	DomainName string `faker:"domain_name" json:"domain_name"`
	IPV4       string `faker:"ipv4" json:"ipv4"`
	IPV6       string `faker:"ipv6" json:"ipv6"`

	CreditCardNumber string `faker:"cc_number" json:"cc_number"`
	CreditCardType   string `faker:"cc_type" json:"cc_type"`
	PaymentMethod    string `faker:"oneof: cc, paypal, check, money order" json:"payment_method"` // oneof will randomly pick one of the comma-separated values supplied in the tag

	Comment string `faker:"sentence" json:"comment"`
}

func FakeData() string {
	a := FakeUser{}
	err := faker.FakeData(&a)
	if err != nil {
		panic(err)
	}
	d, _ := json.Marshal(a)
	return string(d)
}
