/*
@Time : 2023/8/11 下午4:34
@Author : newt
@DESC : TODO
*/

package plugins

type QuestionList struct {
	Data data `json:"data"`
}

type data struct {
	RecentACSubmissions recentACSubmissions `json:"recentACSubmissions"`
}

type recentACSubmissions []*submissions

type submissions struct {
	//Question     *question `json:"question"`
	question     `json:"question"`
	SubmissionId int64 `json:"submissionId"`
	SubmitTime   int64 `json:"submitTime"`
}

type question struct {
	QuestionFrontendId string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TitleSlug          string `json:"titleSlug"`
	TranslatedTitle    string `json:"translatedTitle"`
}
