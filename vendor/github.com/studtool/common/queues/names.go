package queues

const (
	RegistrationEmailsQueueName = "registration_emails_queue"

	CreatedUsersQueueName = "created_users_queue"
	DeletedUsersQueueName = "deleted_users_queue"

	EmailUpdatesQueueName    = "email_updates_queue"
	PasswordUpdatesQueueName = "password_updates_queue"

	AvatarsToCreateQueueName = "created_users_avatars_queue"
	AvatarsToDeleteQueueName = "created_users_avatars_queue"

	ProfilesToCreateQueueName = "profiles_to_create"
	ProfilesToDeleteQueueName = "profiles_to_delete"

	DocumentUsersToCreateQueueName = "document_users_to_create"
	DocumentUsersToDeleteQueueName = "document_users_to_delete"
)
