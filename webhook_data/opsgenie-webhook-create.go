package webhook_data

type OpsgenieWebhookCreate struct {
	Alert OpsgenieWebhookCreateAlert `json:"alert" binding:"required"`

	Action string `json:"action" binding:"required"`
}

type OpsgenieWebhookCreateAlert struct {
	Message string `json:"message" binding:"required"`
}
