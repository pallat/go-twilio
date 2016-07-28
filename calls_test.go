package twilio

import (
	"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	SetAccount("AC3cab34bf1d6c63472a41a40189feb3da", "949989be14ba4480ba245d34f4318c7d")
	resp, err := CreateCall("+15005550006", "client:jenny", "http://demo.twilio.com/docs/voice.xml")
	fmt.Println(err)
	fmt.Println(resp)
}
