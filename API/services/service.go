package services

import (
	"API/models"
	"API/repositories"
)

func Login(user models.User) (string, error) {
	token, err := repositories.Login(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
func Register(user models.User) error {
	err := repositories.Register(user)
	if err != nil {
		return err
	}
	return nil
}
func CreateChallange(challange models.Challange) error {

	err := repositories.CreateChallange(challange)
	if err != nil {
		return err
	}
	return nil
}
func DeleteChallange(id int) error {
	err := repositories.DeleteChallange(id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateChallange(challange models.Challange) error {
	err := repositories.UpdateChallange(challange)
	if err != nil {
		return err
	}
	return nil
}
func GetChallangeById(id int) (models.Challange, error) {
	challange, err := repositories.GetChallangeById(id)
	if err != nil {
		return models.Challange{}, err
	}
	return challange, nil
}
func GetAllChallanges() ([]models.Challange, error) {
	challanges, err := repositories.GetAllChallanges()
	if err != nil {
		return nil, err
	}
	return challanges, nil
}

func AddQuestion(question models.Quest, challangeId int) error {
	err := repositories.AddQuestion(question, challangeId)
	if err != nil {
		return err
	}
	return nil
}
func DeleteQuestion(id int) error {
	err := repositories.DeleteQuestion(id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateQuestion(question models.Quest) error {
	err := repositories.UpdateQuestion(question)
	if err != nil {
		return err
	}
	return nil
}
func GetAllQuestions(challange_id int) ([]models.Quest, error) {
	questions, err := repositories.GetAllQuestions(challange_id)
	if err != nil {
		return nil, err
	}
	return questions, nil
}
func GetQuestionId(id int) (models.Quest, error) {
	question, err := repositories.GetQuestionId(id)
	if err != nil {
		return models.Quest{}, err
	}
	return question, nil
}
func UserAttendChallange(UserChallange models.USER_CHALLANGE) error {
	err := repositories.UserAttendChallange(UserChallange)
	if err != nil {
		return err
	}
	return nil
}
func UserQuest(quest models.Quest, challangeId int, userId int) error {
	err := repositories.UserQuest(quest, challangeId, userId)
	if err != nil {
		return err
	}
	return nil
}
func UserGetAllQuest(challangeId int) ([]models.Quest, error) {
	questions, err := repositories.UserGetAllQuest(challangeId)
	if err != nil {
		return nil, err
	}
	return questions, nil
}
func UserScore() ([]models.ScoreTable, error) {
	score, err := repositories.UserScore()
	if err != nil {
		return nil, err
	}
	return score, nil
}
