/**
 * @Author : NewtSun
 * @Date : 2023/10/4 17:41
 * @Description :
 **/

package cron

import (
	"deploy-trace/global"
	"deploy-trace/model"
	"deploy-trace/plugins"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	AlignmentMinutes = 5
	SchedulerMinutes = 5
	FrequencyMinutes = 10
)

func initTimeArgs() {
	if os.Getenv("ALIGNMENTMINUTES") != "" {
		AlignmentMinutes, _ = strconv.Atoi(os.Getenv("ALIGNMENTMINUTES"))
	}
	if os.Getenv("SCHEDULERMINUTES") != "" {
		SchedulerMinutes, _ = strconv.Atoi(os.Getenv("SCHEDULERMINUTES"))
	}
	if os.Getenv("FREQUENCYMINUTES") != "" {
		FrequencyMinutes, _ = strconv.Atoi(os.Getenv("FREQUENCYMINUTES"))
	}
	fmt.Println("[AlignmentMinutes] args is:", AlignmentMinutes)
	fmt.Println("[SchedulerMinutes] args is:", SchedulerMinutes)
	fmt.Println("[FrequencyMinutes] args is:", FrequencyMinutes)
}

func Cron() {
	initTimeArgs()
	alignmentTime()

	userList := getUserList()

	schedulerTicker := time.NewTicker(time.Duration(SchedulerMinutes) * time.Minute)
	frequencyTicker := time.NewTicker(time.Duration(FrequencyMinutes) * time.Minute)

	for {
		select {
		case <-schedulerTicker.C:
			fmt.Println("[schedulerTicker] Now time is:", time.Now().Format("2006-01-02 15:04:05"))
			for _, user := range userList {
				runSchedule(global.UserName(user.UserName), global.RealName(user.RealName))
			}
		case <-frequencyTicker.C:
			fmt.Println("[frequencyTicker] Now time is:", time.Now().Format("2006-01-02 15:04:05"))
			for _, user := range userList {
				submissionFrequency(time.Now(), user.UserName, user.RealName, user.NickName)
			}
		}
	}

	fmt.Println()
}

func alignmentTime() {
	timestamp := time.Now()

	for timestamp.Minute()%AlignmentMinutes != 0 || timestamp.Second()%60 > 30 {
		fmt.Println("[alignmentTime] Waiting for time alignment.", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(30 * time.Second)
		timestamp = time.Now()
	}

	fmt.Println("[alignmentTime] Now time is OK.", time.Now().Format("2006-01-02 15:04:05"))
}

func submissionFrequency(timeNow time.Time, userName, realName, nickName string) {
	db := global.Connect()

	nowTime := timeNow.Unix()
	lastTime := timeNow.Unix() - int64(FrequencyMinutes)*60

	var submissions []*model.Submission

	db.Where("submit_time_stamp BETWEEN ? AND ? AND user_name = ?", lastTime, nowTime, userName).Find(&submissions)

	value := len(submissions)
	fre := new(model.Frequency)

	fre.Value = int64(value)
	fre.UserName = userName
	fre.RealName = realName
	fre.NickName = nickName

	result := db.Create(&fre)
	if result.Error != nil {
		fmt.Println(result)
	}
}

func getUserList() []*model.User {
	var userList []*model.User

	db := global.Connect()
	result := db.Find(&userList)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return userList
}

func runSchedule(args ...interface{}) {
	var (
		userName global.UserName
		realName global.RealName
		nickName global.NickName
	)

	for _, arg := range args {
		if _, ok := arg.(global.UserName); ok {
			userName = arg.(global.UserName)
		}
		if _, ok := arg.(global.RealName); ok {
			realName = arg.(global.RealName)
		}
		if _, ok := arg.(global.NickName); ok {
			nickName = arg.(global.NickName)
		}
	}

	questionList := new(plugins.QuestionList)

	plugins.Schedule(questionList, userName)

	res := conversion(questionList, string(userName), string(realName), string(nickName))

	checkAndInsert(res)
}

func conversion(questionList *plugins.QuestionList, userName, realName, nickName string) []*model.Submission {
	length := len(questionList.Data.RecentACSubmissions)

	var res []*model.Submission

	for i := 0; i < length; i++ {
		sub := new(model.Submission)
		tmp := questionList.Data.RecentACSubmissions[i]

		sub.Question.Title = tmp.Title
		sub.Question.TitleSlug = tmp.TitleSlug
		sub.Question.TranslatedTitle = tmp.TranslatedTitle
		sub.Question.QuestionFrontendId = tmp.QuestionFrontendId

		sub.SubmitTimeStamp = tmp.SubmitTime
		sub.SubmitTime = time.Unix(tmp.SubmitTime, 0)
		sub.SubmissionId = tmp.SubmissionId

		sub.UserName = userName
		sub.RealName = realName
		sub.NickName = nickName

		res = append(res, sub)
	}

	return res
}

func checkAndInsert(submissions []*model.Submission) {
	length := len(submissions)

	db := global.Connect()

	for i := 0; i < length; i++ {
		var checkRes model.Submission
		db.Where("submission_id = ?", submissions[i].SubmissionId).First(&checkRes)

		if checkRes.SubmissionId != 0 {
			continue
		}

		result := db.Create(submissions[i])
		if result.Error != nil {
			fmt.Println(result)
		}
	}
}
