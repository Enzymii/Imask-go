package controller

func InitRouter() {
	// user
	E.GET("/status", GetStatus)

	E.POST("/login", Login)
	E.POST("/register", Register)

	// oss
	E.GET("/collections", GetCollection)
	E.GET("/upload_url", GetUploadSignature)
	E.GET("/download_url", GetDownloadURL)

	E.POST("/uploaded", UploadFinished)

	// task
	E.GET("/task", GetTasks)
	E.GET("/task/my", GetMyTasks)
	E.POST("/task/create", CreateTask)

	// annotation
	E.POST("/annotation/create", CreateAnnotation)
	E.POST("/annotation/status", UpdateAnnotationStatus)
	E.GET("/annotation", GetAnnotation)
}
