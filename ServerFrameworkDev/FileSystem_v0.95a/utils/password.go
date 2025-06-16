package utils

import (
    "regexp"
)

// CheckPasswordStrength 检查密码强度
func CheckPasswordStrength(password string) bool {
    if len(password) < 8 {
        return false
    }
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
    hasLower := regexp.MustCompile(`[a-z]`).MatchString
    hasDigit := regexp.MustCompile(`\d`).MatchString
    hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString
    
    return hasUpper(password) && 
           hasLower(password) && 
           hasDigit(password) && 
           hasSpecial(password)
}
