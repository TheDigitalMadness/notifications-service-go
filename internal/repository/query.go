package repository

const (
	createNotificationQuery = `
	INSERT INTO "notification"
	("publicType", "userId", "type", "message")
	VALUES
	($1, $2, $3, $4)
	`

	getByUserID = `
	SELECT *
	FROM "notification"
	WHERE "userId" = $1
	`

	getAdminNotifications = `
	SELECT *
	FROM "notifications"
	WHERE "publicType" = $1
	SKIP $2
	LIMIT $3
	`

	setReadQuery = `
	UPDATE "notification"
	SET "read" = true
	WHERE id = ANY($1)
	`
)
