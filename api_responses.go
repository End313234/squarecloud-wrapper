package squarecloud

type BaseResponse[T any] struct {
	Success  string `json:"success"`
	Response T      `json:"response"`
}

type Error struct {
	Status string `json:"status"`
	Code   string `json:"code"`
}

type GetCurrentUserRawResponse struct {
	User         RawUser       `json:"user"`
	Applications []Application `json:"applications"`
}

type GetCurrentUserResponse = BaseResponse[GetCurrentUserRawResponse]
