/**
 * @Author : NewtSun
 * @Date : 2023/10/4 17:22
 * @Description :
 **/

package model

import (
	"gorm.io/gorm"
	"time"
)

type Question struct {
	QuestionFrontendId string
	Title              string
	TitleSlug          string
	TranslatedTitle    string
}

type Submission struct {
	gorm.Model
	Question
	UserName        string
	RealName        string
	NickName        string
	SubmissionId    int64
	SubmitTimeStamp int64
	SubmitTime      time.Time
}
