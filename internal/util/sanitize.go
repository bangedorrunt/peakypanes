package util

import "strings"

// sanitizeSessionName chuẩn hóa tên session tmux
func SanitizeSessionName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return "session"
	}
	var b strings.Builder
	lastDash := false
	for _, r := range name {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
			lastDash = false
		case r >= '0' && r <= '9':
			b.WriteRune(r)
			lastDash = false
		case r == '-' || r == '_' || r == ' ':
			if !lastDash {
				b.WriteRune('-')
				lastDash = true
			}
		}
	}
	result := strings.Trim(b.String(), "-")
	if result == "" {
		return "session"
	}
	return result
}

// GetStringOrEmpty an toàn chuyển đổi interface{} thành string, trả về empty string nếu nil
func GetStringOrEmpty(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
