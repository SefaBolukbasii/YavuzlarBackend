package repositories

import (
	db "API/config"
	"API/models"
	"API/utils"
)

func Login(user models.User) (string, error) {
	row := db.Db.QueryRow("SELECT id,username,role FROM users WHERE username = $1 AND sifre = $2", user.Username, user.Password)
	var kullanici models.User
	err := row.Scan(&kullanici.Id, &kullanici.Username, &kullanici.Role)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateJwt(kullanici.Id, kullanici.Username, kullanici.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
func Register(user models.User) error {
	_, err := db.Db.Exec("INSERT INTO users (username, sifre, role) VALUES ($1, $2, $3)", user.Username, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}
func CreateChallange(challange models.Challange) error {
	var chaID int
	err := db.Db.QueryRow("INSERT INTO challange (name) VALUES ($1) RETURNING id", challange.Name).Scan(&chaID)
	if err != nil {
		return err
	}
	for _, question := range challange.Questions {
		err = AddQuestion(question, chaID)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteChallange(id int) error {
	_, err := db.Db.Exec("DELETE FROM challange WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateChallange(challange models.Challange) error {
	_, err := db.Db.Exec("UPDATE challange SET name = $1 WHERE id = $2", challange.Name, challange.Id)
	if err != nil {
		return err
	}
	return nil
}
func GetAllChallanges() ([]models.Challange, error) {
	rows, err := db.Db.Query("SELECT * FROM challange")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var challanges []models.Challange
	for rows.Next() {
		var challange models.Challange
		err := rows.Scan(&challange.Id, &challange.Name)
		if err != nil {
			return nil, err
		}
		challanges = append(challanges, challange)
	}
	return challanges, nil
}
func GetChallangeById(id int) (models.Challange, error) {
	row := db.Db.QueryRow("SELECT * FROM challange WHERE id = $1", id)
	var challange models.Challange
	err := row.Scan(&challange.Id, &challange.Name)
	if err != nil {
		return models.Challange{}, err
	}
	return challange, nil
}
func AddQuestion(question models.Quest, challangeId int) error {
	_, err := db.Db.Exec("INSERT INTO quest(question, answer,challange_id) VALUES ($1, $2,$3)", question.Question, question.Answer, challangeId)
	if err != nil {
		return err
	}
	return nil
}
func DeleteQuestion(id int) error {
	_, err := db.Db.Exec("DELETE FROM quest WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateQuestion(question models.Quest) error {
	_, err := db.Db.Exec("UPDATE quest SET question = $1, answer = $2 WHERE id = $3", question.Question, question.Answer, question.Id)
	if err != nil {
		return err
	}
	return nil
}
func GetAllQuestions(challange_id int) ([]models.Quest, error) {
	rows, err := db.Db.Query("SELECT id,question,answer FROM quest WHERE challange_id = $1", challange_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Quest
	for rows.Next() {
		var question models.Quest
		err := rows.Scan(&question.Id, &question.Question, &question.Answer)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}
func GetQuestionId(id int) (models.Quest, error) {
	row := db.Db.QueryRow("SELECT id,question,answer FROM quest WHERE id = $1", id)
	var question models.Quest
	err := row.Scan(&question.Id, &question.Question, &question.Answer)
	if err != nil {
		return models.Quest{}, err
	}
	return question, nil
}
func UserAttendChallange(UserChallange models.USER_CHALLANGE) error {
	_, err := db.Db.Exec("INSERT INTO USER_CHALLANGE (user_id, challange_id) VALUES ($1, $2)", UserChallange.UserId, UserChallange.ChallangeId)
	if err != nil {
		return err
	}
	return nil
}
func UserQuest(quest models.Quest, challangeId int, userId int) error {
	row := db.Db.QueryRow("SELECT id,question,answer FROM quest WHERE id = $1", quest.Id)
	var question models.Quest
	err := row.Scan(&question.Id, &question.Question, &question.Answer)
	if err != nil {
		return err
	}
	if question.Answer == quest.Answer {
		_, err := db.Db.Exec(`
			INSERT INTO USER_CHALLANGE (user_id, challange_id, score) 
			VALUES ($1, $2, $3) 
			ON CONFLICT (user_id, challange_id) 
			DO UPDATE SET score = USER_CHALLANGE.score + 10`, userId, challangeId, 10)
		if err != nil {
			return err
		}

	}
	return nil
}
func UserGetAllQuest(challangeId int) ([]models.Quest, error) {
	rows, err := db.Db.Query("SELECT id,question FROM quest WHERE challange_id = $1", challangeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Quest
	for rows.Next() {
		var question models.Quest
		err := rows.Scan(&question.Id, &question.Question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}
func UserScore() ([]models.ScoreTable, error) {
	rows, err := db.Db.Query("SELECT challange.name,users.username,USER_CHALLANGE.score FROM challange " +
		"INNER JOIN USER_CHALLANGE ON challange.id = USER_CHALLANGE.challange_id " +
		"INNER JOIN users ON users.id = USER_CHALLANGE.user_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []models.ScoreTable
	for rows.Next() {
		var score models.ScoreTable
		err := rows.Scan(&score.Challange.Name, &score.User.Username, &score.Score)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}
