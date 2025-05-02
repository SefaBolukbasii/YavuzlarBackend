package routes

import (
	"API/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/api/v1/login", controllers.Login)
	http.HandleFunc("/api/v1/register", controllers.Register)
	http.HandleFunc("/api/v1/admin/challange", controllers.CreateChallange)
	http.HandleFunc("/api/v1/admin/challange/", controllers.DeleteUpdataSelectChallange)
	http.HandleFunc("/api/v1/admin/quest/", controllers.AdminQuest)
	http.HandleFunc("/api/v1/user/challange", controllers.UserChallange)
	http.HandleFunc("/api/v1/user/challange/", controllers.GetChallangeById)
	http.HandleFunc("/api/v1/user/quest/answer/", controllers.UserQuest)
	http.HandleFunc("/api/v1/user/quest/", controllers.UserQuestAll)
	http.HandleFunc("/api/v1/user/quest/score", controllers.UserScore)

}
