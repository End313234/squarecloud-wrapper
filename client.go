package squarecloud

import (
	"errors"
	"fmt"
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
		client.Logger.Printf("Unsuccessful request to /user: %s\n", err.Error())

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

	client.Cache.Users.addsToCacheIfTargetDoesNotExist(currentUser, func() {
		client.Logger.Printf("User of ID %s has been cached\n", currentUser.Id)
	})

	client.Logger.Println("Successful request to /user")

	return currentUser, nil
}

// Searches for an user in the local cache and makes a request
// to the SquareCloud API if the user couldn't be found in the
// cache
func (client *Client) GetUser(id string) (User, error) {
	for _, user := range client.Cache.Users {
		if user.Id == id {
			return user, nil
		}
	}

	var fetchedUser GetUserResponse
	err := client.Requester.Get(fmt.Sprintf("/user/%s", id), &fetchedUser)
	if err != nil {
		return User{}, err
	}

	user := User{
		Id:            fetchedUser.Response.User.Id,
		Tag:           fetchedUser.Response.User.Tag,
		Email:         fetchedUser.Response.User.Email,
		Plan:          fetchedUser.Response.User.Plan,
		IsBlocklisted: fetchedUser.Response.User.IsBlocklisted,
		Applications:  fetchedUser.Response.Applications,
	}

	client.Logger.Printf("Successful request to /user/%s\n", user.Id)

	client.Cache.Users.addsToCacheIfTargetDoesNotExist(user, func() {
		client.Logger.Printf("User of ID %s has been cached\n", user.Id)
	})

	return user, nil
}
