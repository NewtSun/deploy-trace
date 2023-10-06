/*
@Time : 2023/8/10 下午2:47
@Author : newt
@DESC : TODO
*/

package plugins

import (
	"deploy-trace/global"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func Schedule(conversionOutPut interface{}, args ...interface{}) {
	var (
		userName global.UserName
		//realName global.RealName
	)

	for _, arg := range args {
		if _, ok := arg.(global.UserName); ok {
			userName = arg.(global.UserName)
		}
		if _, ok := arg.(global.RealName); ok {
			//realName = arg.(global.RealName)
		}
	}

	//cmd := exec.Command("python", "static/plugin/test.py", "--name", string(userName))
	cmd := exec.Command("python3", "linux_test.py", "--name", string(userName))

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd run failed with %s", err)
	}

	err = json.Unmarshal(output, &conversionOutPut)
	if err != nil {
		log.Fatalf("json unmarshal failed with %s", err)
	}

	//unmarshalPrint(conversionOutPut)
}

func unmarshalPrint(needPrint interface{}) {
	if _, ok := needPrint.(*QuestionList); !ok {
		return
	} else {
		//marshal, err := json.Marshal(needPrint)
		marshal, err := json.MarshalIndent(needPrint, "", "    ")
		if err != nil {
			log.Fatalf("json marshal failed with %s", err)
			return
		}
		fmt.Println(string(marshal))
	}
}
