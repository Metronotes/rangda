package barong

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSession_GetUserID_ShoulldReturnError_IfInvalidData(t *testing.T) {
	test := assert.New(t)

	testcases := []string{
		`[]`,
		`[1]`,
		`[1,2]`,
		`[1,[]]`,
	}

	for _, testcase := range testcases {
		var session Session
		err := json.Unmarshal(
			[]byte(`{"session_id":"1658bb7d75d566e2954a0c4c227ea43e",`+
				`"warden.user.account.key":`+testcase+`,`+
				`"warden.user.account.session":{},"_csrf_token":`+
				`"g0X8U3cE4kwgEUDfN7CLMm7u7nXbPUOtXvTNSZKeZWw="}`),
			&session,
		)
		test.NoError(err)

		id, err := session.GetUserID()
		test.Zero(id)
		test.Error(err)
	}
}

func TestSession_GetUserID_ShoulldReturnID_IfValidData(t *testing.T) {
	test := assert.New(t)

	var session Session
	err := json.Unmarshal(
		[]byte(`{"session_id":"1658bb7d75d566e2954a0c4c227ea43e",`+
			`"warden.user.account.key":[[1],"$2a$11$5vp7Ujjifr11zgjEu1GSwu"],`+
			`"warden.user.account.session":{},"_csrf_token":`+
			`"g0X8U3cE4kwgEUDfN7CLMm7u7nXbPUOtXvTNSZKeZWw="}`),
		&session,
	)
	test.NoError(err)

	id, err := session.GetUserID()
	test.NoError(err)
	test.Equal(int64(1), id)
}
