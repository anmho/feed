package users

import (
	"database/sql"
	feedv1 "feed/gen/protos/feed/v1"
	"fmt"
)




func CreateUser(db *sql.DB, user *feedv1.User) error {
	_, err := db.Exec(
	`
	INSERT INTO users (
		user_id, name
	) VALUES (
		$1, $2
	)
	;
	`,
	user.UserId, user.Name,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

