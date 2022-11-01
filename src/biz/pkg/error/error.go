package error

import (
	"fmt"
	"time"
)

type TokenTimeoutError struct {
	UserName        	string
	CreateTime			time.Time
	Role				string
}

func (e *TokenTimeoutError) Error() string {
   return fmt.Sprintf("Token out of time: %s %s %s",e.UserName, e.CreateTime.Format("2006-01-02 15:04:05"), e.Role)
}

// func test() error {
//    return &MyError{ "Something happened", "server.go",  42 }
// }