package controllers

import (
	"API/models"
	"API/services"
	"API/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		token, err := services.Login(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		utils.WriteJson(w, token)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		err = services.Register(user)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		message := models.Message{Message: "User registered successfully"}
		utils.WriteJson(w, message)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func CreateChallange(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if user.Role == "admin" {
		if r.Method == "POST" {
			var challange models.Challange
			err := json.NewDecoder(r.Body).Decode(&challange)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			err = services.CreateChallange(challange)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			message := models.Message{
				Message: "Challange created successfully"}
			utils.WriteJson(w, message)

		} else if r.Method == "GET" {
			challanges, err := services.GetAllChallanges()
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			utils.WriteJson(w, challanges)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	} else {
		http.Error(w, "Admin Değilsin "+user.Role, http.StatusUnauthorized)
		return
	}

}
func DeleteUpdataSelectChallange(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if user.Role == "admin" {
		id := strings.Split(r.URL.Path, "/")
		challangeId := id[len(id)-1]
		challangeIdInt, err := strconv.Atoi(challangeId)
		if err != nil {
			http.Error(w, "Invalid challange ID", http.StatusBadRequest)
			return
		}

		if r.Method == "DELETE" {

			err = services.DeleteChallange(challangeIdInt)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			http.Error(w, "Challange deleted successfully "+challangeId, http.StatusOK)
		} else if r.Method == "PUT" {

			var challange models.Challange
			err = json.NewDecoder(r.Body).Decode(&challange)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			challange.Id = challangeIdInt
			err = services.UpdateChallange(challange)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			message := models.Message{Message: "Challange Güncellendi"}
			utils.WriteJson(w, message)
		} else if r.Method == "GET" {
			challange, err := services.GetChallangeById(challangeIdInt)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			utils.WriteJson(w, challange)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

}
func AdminQuest(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if user.Role == "admin" {
		var challangeIdInt int
		var questIdInt int
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) == 7 {
			challangeId := parts[len(parts)-2]
			challangeIdInt, err = strconv.Atoi(challangeId)
			if err != nil {
				http.Error(w, "Invalid challange ID", http.StatusBadRequest)
				return
			}
			questId := parts[len(parts)-1]
			questIdInt, err = strconv.Atoi(questId)
			if err != nil {
				http.Error(w, "Invalid quest ID", http.StatusBadRequest)
				return
			}
		} else if len(parts) == 6 {
			challangeId := parts[len(parts)-1]
			challangeIdInt, err = strconv.Atoi(challangeId)
			if err != nil {
				http.Error(w, "Invalid challange ID", http.StatusBadRequest)
				return
			}
		} else {
			http.Error(w, "Hatalı parametre", http.StatusBadRequest)
			return
		}

		if r.Method == "POST" {
			var quest models.Quest
			err := json.NewDecoder(r.Body).Decode(&quest)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			err = services.AddQuestion(quest, challangeIdInt)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			message := models.Message{Message: "Question created successfully"}
			utils.WriteJson(w, message)

		} else if r.Method == "DELETE" {
			err := services.DeleteQuestion(questIdInt)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			message := models.Message{Message: "Question delete successfully"}
			utils.WriteJson(w, message)
		} else if r.Method == "PUT" {
			var quest models.Quest
			err := json.NewDecoder(r.Body).Decode(&quest)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			err = services.UpdateQuestion(quest)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			message := models.Message{Message: "Question update successfully" + quest.Question}
			utils.WriteJson(w, message)
		} else if r.Method == "GET" {
			if len(parts) == 7 {
				quest, err := services.GetQuestionId(questIdInt)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				utils.WriteJson(w, quest)
			} else if len(parts) == 6 {
				quest, err := services.GetAllQuestions(challangeIdInt)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				utils.WriteJson(w, quest)
			}

		} else {
			http.Error(w, "Geçersiz Method", http.StatusBadRequest)
			return
		}
	}

}
func UserChallange(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if user.Role == "user" {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var UserChallange models.USER_CHALLANGE
		err := json.NewDecoder(r.Body).Decode(&UserChallange)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		err = services.UserAttendChallange(UserChallange)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

}
func GetChallangeById(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if user.Role == "user" {
		id := strings.Split(r.URL.Path, "/")
		challangeId := id[len(id)-1]
		challangeIdInt, err := strconv.Atoi(challangeId)
		if err != nil {
			http.Error(w, "Invalid challange ID", http.StatusBadRequest)
			return
		}
		challange, err := services.GetChallangeById(challangeIdInt)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		utils.WriteJson(w, challange)
	}

}
func UserQuest(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if user.Role == "user" {
		id := strings.Split(r.URL.Path, "/")
		challangeId := id[len(id)-2]
		challangeIdInt, err := strconv.Atoi(challangeId)
		if err != nil {
			http.Error(w, "Invalid challange ID", http.StatusBadRequest)
			return
		}
		questId := id[len(id)-1]
		questIdInt, err := strconv.Atoi(questId)
		if err != nil {
			http.Error(w, "Invalid challange ID", http.StatusBadRequest)
			return
		}
		if r.Method == "POST" {
			var quest models.Quest
			err := json.NewDecoder(r.Body).Decode(&quest)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			quest.Id = questIdInt
			err = services.UserQuest(quest, challangeIdInt, user.Id)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			message := models.Message{Message: "Question answer successfully"}
			utils.WriteJson(w, message)

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

}
func UserQuestAll(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if user.Role == "user" {
		id := strings.Split(r.URL.Path, "/")
		challangeId := id[len(id)-1]
		challangeIdInt, err := strconv.Atoi(challangeId)
		if err != nil {
			http.Error(w, "Invalid challange ID", http.StatusBadRequest)
			return
		}
		if r.Method == "GET" {
			quest, err := services.UserGetAllQuest(challangeIdInt)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			utils.WriteJson(w, quest)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

}
func UserScore(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ExtractToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if user.Role == "user" {
		if r.Method == "GET" {
			score, err := services.UserScore()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			utils.WriteJson(w, score)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

}
