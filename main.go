/**
 * @Author : NewtSun
 * @Date : 2023/8/12 13:41
 * @Description :
 **/

package main

import (
	"deploy-trace/cron"
	"fmt"
	"time"
)

func init() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Local = loc
}

func main() {
	cron.Cron()
}
