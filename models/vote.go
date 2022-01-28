package models

type Vote struct {
	UserId int64 `json:"user_id"`
	PostId int64 `json:"post_id"`
	State  int64 `json:"state"`
}

func (v *Vote) Validate() map[string]string {
	var errMsgs map[string]string = map[string]string{}

	if v.State != 1 && v.State != -1 {
		errMsgs["state"] = "state must be 1 or -1"
	}

	return errMsgs
}
