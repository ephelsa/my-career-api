package server

import (
	"bytes"
	"encoding/json"
	"ephelsa/my-career/internal/env"
	sharedDomain "ephelsa/my-career/pkg/shared/domain"
	sharedServer "ephelsa/my-career/pkg/shared/infrastructure/server"
	"ephelsa/my-career/pkg/survey/data"
	"ephelsa/my-career/pkg/survey/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type handler struct {
	Repository            data.SurveyLocalRepository
	ClassifierModelConfig env.ClassifierModel
}

func NewSurveyServer(remote *fiber.App, repo data.SurveyLocalRepository, classifierModel env.ClassifierModel) data.SurveyServerRepository {
	handler := &handler{
		Repository:            repo,
		ClassifierModelConfig: classifierModel,
	}

	survey := remote.Group("/survey")

	survey.Get("/:id/questions-with-answers", handler.FetchActiveSurveyById)
	survey.Post("/by-user", handler.FetchAll)
	survey.Post("/bulk-answers", handler.BulkQuestionAnswer)
	survey.Post("/classify", handler.ClassifySurveyAnswersByAttempt)
	survey.Put("/:id/answer", handler.NewQuestionAnswer)

	return handler
}

func (h *handler) FetchAll(c *fiber.Ctx) error {
	body := c.Body()
	userInfo := domain.UserAnswer{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	result, err := h.Repository.FetchAll(c.Context(), userInfo)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	if len(result) == 0 {
		return sharedServer.NotFound(c, sharedDomain.Error{
			Message: sharedDomain.ResourceEmpty,
			Details: sharedDomain.ResourcesEmpty.Error(),
		})
	}

	return sharedServer.OK(c, result)
}

func (h *handler) FetchActiveSurveyById(c *fiber.Ctx) error {
	surveyId := c.Params("id")
	result, err := h.Repository.FetchActiveSurveyById(c.Context(), surveyId)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	if len(result.Questions) == 0 {
		return sharedServer.NotFound(c, sharedDomain.Error{
			Message: sharedDomain.ResourceEmpty,
			Details: sharedDomain.ResourcesEmpty.Error(),
		})
	}

	return sharedServer.OK(c, result)
}

func (h *handler) NewQuestionAnswer(c *fiber.Ctx) error {
	surveyId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}
	body := c.Body()
	userAnswer := domain.UserAnswer{}
	if err := json.Unmarshal(body, &userAnswer); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}
	userAnswer.Survey = surveyId

	if err := h.Repository.NewQuestionAnswer(c.Context(), userAnswer); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, sharedDomain.SuccessStatus)
}

func (h *handler) BulkQuestionAnswer(c *fiber.Ctx) error {
	body := c.Body()
	var bulkUserAnswer []*domain.UserAnswer
	if err := json.Unmarshal(body, &bulkUserAnswer); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	if err := h.Repository.BulkQuestionAnswer(c.Context(), bulkUserAnswer); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, sharedDomain.SuccessStatus)
}

func (h *handler) ClassifySurveyAnswersByAttempt(c *fiber.Ctx) error {
	body := c.Body()
	userAnswer := domain.UserAnswer{}
	if err := json.Unmarshal(body, &userAnswer); err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	result, err := h.Repository.FetchUserAnswers(c.Context(), userAnswer)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	classification, err := h.consumeClassifierService(result)
	if err != nil {
		return sharedServer.InternalServerError(c, sharedDomain.Error{
			Message: sharedDomain.UnexpectedError,
			Details: err.Error(),
		})
	}

	return sharedServer.OK(c, classification)
}

func (h *handler) consumeClassifierService(a []domain.UserAnswer) ([]domain.ClassifierAnswerResponse, error) {
	if len(a) < 30 {
		return nil, sharedDomain.ResourcesInvalid
	}

	classify := domain.ClassifyAnswersRequest{
		LikeChallenges:               a[0].ClassifierValue,
		AffinityMath:                 a[1].ClassifierValue,
		AffinityPhysics:              a[2].ClassifierValue,
		AffinityChemistry:            a[3].ClassifierValue,
		AffinityHumanSci:             a[4].ClassifierValue,
		AffinityBiologySci:           a[5].ClassifierValue,
		AffinityEng:                  a[6].ClassifierValue,
		InterestNewResults:           a[7].ClassifierValue,
		LikeMentalGames:              a[8].ClassifierValue,
		LikeImproveOpportunities:     a[9].ClassifierValue,
		InterestHowThingWorks:        a[10].ClassifierValue,
		LikeIndividualWork:           a[11].ClassifierValue,
		HowManyCreative:              a[12].ClassifierValue,
		AreYouLeader:                 a[13].ClassifierValue,
		HowManyConcentration:         a[14].ClassifierValue,
		AreYouAutodidact:             a[15].ClassifierValue,
		VocacionalOrientationInside:  a[16].ClassifierValue,
		VocacionalOrientationOutside: a[17].ClassifierValue,
		LikePracticalUsages:          a[18].ClassifierValue,
		EnjoyOlderPeople:             a[19].ClassifierValue,
		HowManyOrganized:             a[20].ClassifierValue,
		FeelParentalSupport:          a[21].ClassifierValue,
		HowManyParentalSupport:       a[22].ClassifierValue,
		LikeTeach:                    a[23].ClassifierValue,
		LikeDissarm:                  a[24].ClassifierValue,
		ExperienceTeamWorks:          a[25].ClassifierValue,
		AssimilateText:               a[26].ClassifierValue,
		LikePaint:                    a[27].ClassifierValue,
		LikeBuild:                    a[28].ClassifierValue,
		LikeRead:                     a[29].ClassifierValue,
	}
	requestBody, err := json.Marshal(classify)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	client := http.Client{
		Timeout: 60 * time.Second,
	}
	url := fmt.Sprintf("http://%s:%s", h.ClassifierModelConfig.URL, h.ClassifierModelConfig.Port)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	request.Header.Set("Content-type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	var objResponse []domain.ClassifierAnswerResponse
	err = json.Unmarshal(body, &objResponse)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return objResponse, nil
}
