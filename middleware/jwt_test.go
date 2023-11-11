package middleware_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/lwinmgmg/linux-http/utils"
)

func TestGetToken(t *testing.T) {
	tkn := utils.GetToken("github", "letmein", time.Hour)
	fmt.Println(tkn)
}
