package middleware_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/lwinmgmg/linux-http/middleware"
)

func TestGetToken(t *testing.T) {
	tkn := middleware.GetToken("github", "letmein", time.Hour)
	fmt.Println(tkn)
}
