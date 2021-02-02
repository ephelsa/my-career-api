package domain

// SurveyWithQuestions
type SurveyWithQuestions struct {
	Id          int                   `json:"survey_id"`
	Name        string                `json:"survey_name"`
	Description interface{}           `json:"description,omitempty"`
	Questions   []QuestionWithAnswers `json:"questions"`
}

// QuestionWithAnswers
type QuestionWithAnswers struct {
	Id       int      `json:"question_id"`
	Question string   `json:"question"`
	Type     string   `json:"type"`
	Answers  []Answer `json:"answers,omitempty"`
}

// Answer
type Answer struct {
	Id              interface{} `json:"id"`
	Value           interface{} `json:"value"`
	ClassifierValue interface{} `json:"classifier_value,omitempty"`
}

// Survey
type Survey struct {
	Id                int         `json:"survey_id"`
	Name              string      `json:"survey_name"`
	Description       interface{} `json:"description,omitempty"`
	IsActive          bool        `json:"is_active"`
	ResolveAttempt    interface{} `json:"resolve_attempt,omitempty"`
	QuestionsAnswered interface{} `json:"questions_answered,omitempty"`
	TotalQuestions    int         `json:"total_questions"`
}

// UserAnswer is used with the table user_answer
type UserAnswer struct {
	Email            string      `json:"email,omitempty"`
	DocumentTypeCode string      `json:"document_type_code,omitempty"`
	Document         string      `json:"document,omitempty"`
	Survey           int         `json:"survey,omitempty"`
	Question         int         `json:"question,omitempty"`
	Answer           interface{} `json:"answer,omitempty"`
	ClassifierValue  int         `json:"classifier_value,omitempty"`
	ResolveAttempt   int         `json:"resolve_attempt,omitempty"`
}

type ClassifyAnswersRequest struct {
	LikeChallenges               int `json:"like_challenges"`
	AffinityMath                 int `json:"affinity_math"`
	AffinityPhysics              int `json:"affinity_physics"`
	AffinityChemistry            int `json:"affinity_chemistry"`
	AffinityHumanSci             int `json:"affinity_human_sci"`
	AffinityBiologySci           int `json:"affinity_biology_sci"`
	AffinityEng                  int `json:"affinity_eng"`
	InterestNewResults           int `json:"interest_new_results"`
	LikeMentalGames              int `json:"like_mental_games"`
	LikeImproveOpportunities     int `json:"like_improve_opportunities"`
	InterestHowThingWorks        int `json:"interest_how_thing_works"`
	LikeIndividualWork           int `json:"like_individual_work"`
	HowManyCreative              int `json:"how_many_creative"`
	AreYouLeader                 int `json:"are_you_leader"`
	HowManyConcentration         int `json:"how_many_concentration"`
	AreYouAutodidact             int `json:"are_you_autodidact"`
	VocacionalOrientationInside  int `json:"vocacional_orientation_inside"`
	VocacionalOrientationOutside int `json:"vocacional_orientation_outside"`
	LikePracticalUsages          int `json:"like_practical_usages"`
	EnjoyOlderPeople             int `json:"enjoy_older_people"`
	HowManyOrganized             int `json:"how_many_organized"`
	FeelParentalSupport          int `json:"feel_parental_support"`
	HowManyParentalSupport       int `json:"how_many_parental_support"`
	LikeTeach                    int `json:"like_teach"`
	LikeDissarm                  int `json:"like_dissarm"`
	ExperienceTeamWorks          int `json:"experience_team_works"`
	AssimilateText               int `json:"assimilate_text"`
	LikePaint                    int `json:"like_paint"`
	LikeBuild                    int `json:"like_build"`
	LikeRead                     int `json:"like_read"`
}

type ClassifierAnswerResponse struct {
	Career  string  `json:"name"`
	Percent float64 `json:"percent"`
}
