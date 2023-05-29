package activity

type MessageController struct {
	Service *ActivityService
}

func NewUserController(service *ActivityService) *MessageController {
	return &MessageController{
		Service: service,
	}
}
