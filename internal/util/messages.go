package util

import "fmt"

// SuccessMessage in ra thông báo thành công với emoji
func SuccessMessage(msg string) {
	fmt.Printf("✅ %s\n", msg)
}

// ErrorMessage in ra thông báo lỗi với emoji
func ErrorMessage(msg string) {
	fmt.Printf("❌ %s\n", msg)
}

// InfoMessage in ra thông báo thông tin với emoji
func InfoMessage(msg string) {
	fmt.Printf("ℹ️  %s\n", msg)
}

// WarningMessage in ra thông báo cảnh báo với emoji
func WarningMessage(msg string) {
	fmt.Printf("⚠️  %s\n", msg)
}
