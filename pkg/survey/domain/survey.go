package domain

type SurveyWithQuestions struct {
	Id          int         `json:"survey_id"`
	Name        string      `json:"survey_name"`
	Description interface{} `json:"description,omitempty"`
	Questions   []Question  `json:"questions"`
}

type Question struct {
	Id       int      `json:"question_id"`
	Question string   `json:"question"`
	Type     string   `json:"type"`
	Answers  []Answer `json:"answers,omitempty"`
}

type Answer struct {
	Id    interface{} `json:"id"`
	Value interface{} `json:"value"`
}

type Survey struct {
	Id          int         `json:"survey_id"`
	Name        string      `json:"survey_name"`
	Description interface{} `json:"description,omitempty"`
	IsActive    bool        `json:"is_active"`
}
