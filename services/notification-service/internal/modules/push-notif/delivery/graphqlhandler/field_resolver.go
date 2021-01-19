package graphqlhandler

import "monorepo/services/notification-service/internal/modules/push-notif/domain"

type pushInputResolver struct {
	Payload *domain.PushNotifRequestPayload
}

type scheduleNotifInputResolver struct {
	Payload struct {
		ScheduledAt string
		Data        *domain.PushNotifRequestPayload
	}
}

type subscribeInputResolver struct {
	Token string
	Topic string
}

type inputTopicEvent struct {
	ID      string
	Message string
	Topic   string
}
