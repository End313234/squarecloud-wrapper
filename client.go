package squarecloud

import (
	"errors"
	"log"

	"github.com/End313234/squarecloud-wrapper/internal/utils"
)

type Client struct {
	APIToken    string
	Logger      *log.Logger
	Requester   *requester
	isConnected bool
	Cache       Cache
}

// Instantiate a new version of squarecloud.Client
func NewClient(APIToken string, logger *log.Logger) *Client {
	return &Client{
		APIToken:  APIToken,
		Logger:    logger,
		Requester: newRequester(APIToken),
		Cache: Cache{
			Users: []User{},
		},
	}
}

func (client *Client) checkIfIsConnected() {
	if !client.isConnected {
		panic("client must be connected")
	}
}

func (client *Client) hasLogger() bool {
	return client.Logger != nil
}

// Makes a request to SquareCloud API to check if the API key
// is a valid one, even though the current user is not added
// to the cache in this method.
func (client *Client) Connect() error {
	var currentUser GetCurrentUserResponse
	err := client.Requester.Get("/user", &currentUser)
	if err != nil {
		return utils.ThrowSquareCloudAPIError(err.Error())
	}

	client.isConnected = true

	if client.hasLogger() {
		client.Logger.Println("The connection has been established")
	}

	return nil
}

// Fetches the current user and adds it to the cache
func (client *Client) FetchCurrentUser() (User, error) {
	currentUser := User{}

	if !client.isConnected {
		return currentUser, errors.New("client must be connected")
	}

	var currentRawUser GetCurrentUserResponse
	err := client.Requester.Get("/user", &currentRawUser)
	if err != nil {
		return currentUser, utils.ThrowSquareCloudAPIError(err.Error())
	}

	currentUser = User{
		Id:            currentRawUser.Response.User.Id,
		Tag:           currentRawUser.Response.User.Tag,
		Email:         currentRawUser.Response.User.Email,
		Plan:          currentRawUser.Response.User.Plan,
		IsBlocklisted: currentRawUser.Response.User.IsBlocklisted,
		Applications:  currentRawUser.Response.Applications,
	}

	if !client.Cache.Users.Contains(currentUser) {
		client.Cache.Users.Add(currentUser)
	}

	return currentUser, nil
}
