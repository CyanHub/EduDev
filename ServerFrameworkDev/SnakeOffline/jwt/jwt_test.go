package jwt

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken()
	fmt.Println(token)
}
func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiLniafniZvngavpk74iLCJzdWIiOiJKd3QgRm9yIEdvIFdlYiIsImV4cCI6MTczMzgxNDAwNiwibmJmIjoxNzMzODEzOTQ2fQ.riT2dfSmzm5FoZOOC5Ku-VUbLJ8O5ej5y8OWANGdvbM\n"
	fmt.Println(ParseToken(token))

}
