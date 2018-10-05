package barong

import (
	"errors"
	"fmt"
)

type Session struct {
	SessionID string

	// that's how it looks like in JSON
	// "warden.user.account.key":[[1],"$2a$11$5vp7Ujjifr11zgjEu1GSwu"]
	WardenUserAccountKey []interface{} `json:"warden.user.account.key"`

	// that's how it looks like in JSON
	// "warden.user.account.session":{}
	WardenUserAccountSession interface{} `json:"warden.user.account.session"`

	CSRFToken string `yaml:"_csrf_token"`
}

func (session *Session) GetUserID() (int64, error) {
	if len(session.WardenUserAccountKey) < 1 {
		return 0, errors.New("invalid warden.user.account.key data")
	}

	userData, ok := session.WardenUserAccountKey[0].([]interface{})
	if !ok {
		return 0, fmt.Errorf(
			"invalid first item in warden.user.account.key, "+
				"expected []interface{} but got %T",
			session.WardenUserAccountKey[0],
		)
	}

	if len(userData) != 1 {
		return 0, fmt.Errorf(
			"invalid first item in warden.user.account.key, "+
				"expected 1 int value but got %v values",
			len(userData),
		)
	}

	// everything is Number for JSON
	userID, ok := userData[0].(float64)
	if !ok {
		return 0, fmt.Errorf(
			"invalid user id item: %T",
			userData[0],
		)
	}

	return int64(userID), nil
}
