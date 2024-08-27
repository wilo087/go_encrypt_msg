package main

import (
	"fmt"

	"github.com/wilo087/go_encrypt_msg/packages/encrypt_msg"
	"github.com/wilo087/go_encrypt_msg/packages/sumtozero"
)

func main() {
	encryptMsg := encrypt_msg.EncryptMsg{
		Key:     "dcj",
		Message: "I love prOgrAmming!",
	}

	// Test for the encryption function
	encryptedMessage := encryptMsg.EncryptMessage()
	fmt.Println("Encrypted Message:", encryptedMessage)

	// Test for the sum to zero function
	sumToZero := sumtozero.SumToZero{
		NumberList: []int{3, 4, -7, 5, -6, 2, 5, -1, 8},
	}

	result := sumToZero.RemoveZeroSumSublists()
	fmt.Println("Sum to zero:", result)
}
