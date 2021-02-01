package database

import (
	"context"
	"database/sql"
	"ephelsa/my-career/pkg/shared/infrastructure/database"
	"ephelsa/my-career/pkg/survey/data"
	"ephelsa/my-career/pkg/survey/domain"
	"github.com/sirupsen/logrus"
)

type postgresSurveyRepo struct {
	Connection *sql.DB
}

type surveyWithQuestions struct {
	surveyId          int
	surveyName        string
	surveyDescription interface{}
	questionId        int
	question          string
	questionType      string
	answerId          interface{}
	answer            interface{}
}

func NewPostgresSurveyRepository(db *sql.DB) data.SurveyLocalRepository {
	return &postgresSurveyRepo{
		Connection: db,
	}
}

func (p *postgresSurveyRepo) FetchAll(ctx context.Context, user domain.UserAnswer) (result []domain.Survey, err error) {
	query := `SELECT s.id, s.name, s.description, s.active, ljua.resolve_id, ljua.question_answered, count(sq.question_id) AS total_question
			FROM survey s
			LEFT JOIN (
				SELECT ua.resolve_id, count(ua.question) AS question_answered, ua.survey, ua.email, ua.document, ua.document_type
				FROM survey as s1
				LEFT JOIN user_answer ua ON ua.survey = s1.id
				WHERE ua.document_type = $1 AND ua.document = $2 AND ua.email = $3
				GROUP BY ua.survey, ua.email, ua.survey, ua.resolve_id, ua.document, ua.document_type
				) AS ljua ON ljua.survey = s.id
			INNER JOIN survey_question sq on s.id = sq.survey_id
			WHERE s.active = true
			GROUP BY s.id, s.name, s.description, s.active, ljua.resolve_id, ljua.question_answered`
	rows, err := database.NewRowsByQueryContext(p.Connection, ctx, query, user.DocumentTypeCode, user.Document, user.Email)
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.Survey, 0)
	for rows.Next() {
		r := domain.Survey{}
		if err := rows.Scan(&r.Id, &r.Name, &r.Description, &r.IsActive, &r.ResolveAttempt, &r.QuestionsAnswered, &r.TotalQuestions); err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}

func (p *postgresSurveyRepo) FetchActiveSurveyById(c context.Context, surveyId string) (result domain.SurveyWithQuestions, err error) {
	query := `SELECT s.id AS survey_id,
				   s.name,
                   s.description,
				   q.id AS question_id,
				   q.question,
				   qt.type,
				   a.id AS option_id,
				   a.option
			FROM survey s
					 INNER JOIN survey_question sq ON s.id = sq.survey_id
					 INNER JOIN question q ON q.id = sq.question_id
					 INNER JOIN question_type qt ON q.question_type = qt.id
					 LEFT JOIN answer_options ao ON q.answer_options = ao.code
					 LEFT JOIN answer_option a ON ao.answer_option = a.id
			WHERE s.id = $1 AND s.active = true
			ORDER BY q.id, a.id ASC;`
	rows, err := database.NewRowsByQueryContext(p.Connection, c, query, surveyId)
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()
	if err != nil {
		logrus.Error(err)
		return
	}

	var queryResult []surveyWithQuestions
	for rows.Next() {
		r := surveyWithQuestions{}
		if err = rows.Scan(
			&r.surveyId,
			&r.surveyName,
			&r.surveyDescription,
			&r.questionId,
			&r.question,
			&r.questionType,
			&r.answerId,
			&r.answer,
		); err != nil {
			logrus.Error(err)
			return
		}

		queryResult = append(queryResult, r)
	}

	return groupAnswersByQuestions(queryResult), nil
}

// groupAnswersByQuestions this method group all answers of the question and store it in an array
// and same with the questions.
func groupAnswersByQuestions(surveys []surveyWithQuestions) (result domain.SurveyWithQuestions) {
	answers := make([]domain.Answer, 0)
	newIndex := 0
	for i, iSurvey := range surveys {
		result.Id = iSurvey.surveyId
		result.Name = iSurvey.surveyName
		result.Description = iSurvey.surveyDescription

		if newIndex < i {
			continue
		}

		for j, jSurvey := range surveys[newIndex:] {
			w := j + newIndex
			if w < len(surveys)-1 && iSurvey.questionId == surveys[w].questionId || w == i {
				if jSurvey.answerId != nil {
					answers = append(answers, domain.Answer{
						Id:    jSurvey.answerId,
						Value: jSurvey.answer,
					})
				}
			} else {
				newIndex += j
				break
			}
		}

		if i >= len(surveys)-1 || iSurvey.questionId != surveys[i+1].questionId {
			result.Questions = append(result.Questions, domain.QuestionWithAnswers{
				Id:       iSurvey.questionId,
				Question: iSurvey.question,
				Type:     iSurvey.questionType,
				Answers:  answers,
			})
			answers = make([]domain.Answer, 0)
		}
	}

	return
}

func (p *postgresSurveyRepo) NewQuestionAnswer(c context.Context, ua domain.UserAnswer) error {
	query := `INSERT INTO user_answer (email, document_type, document, question, answer, survey, resolve_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT (email, document_type, document, question, survey, resolve_id)
				DO UPDATE SET answer = EXCLUDED.answer;`
	stmt, err := p.Connection.PrepareContext(c, query)
	if err != nil {
		return err
	}
	defer func() {
		if err = stmt.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	_, err = stmt.ExecContext(c, ua.Email, ua.DocumentTypeCode, ua.Document, ua.Question, ua.Answer, ua.Survey, ua.ResolveAttempt)
	if err != nil {
		logrus.Error(err)
	}

	return err
}
