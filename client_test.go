package squarecloud

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/End313234/squarecloud-wrapper/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()

	// os.Remove("./logs.log")

	os.Exit(exitCode)
}

func TestCreateClientWithoutLoggingSuccess(t *testing.T) {
	assert := assert.New(t)

	client := NewClient(utils.GetEnv("API_KEY"), nil)
	assert.NotEmpty(client)
	assert.Nil(client.Logger)
}

func TestCreateClientWithLoggingSuccess(t *testing.T) {
	assert := assert.New(t)

	file, err := os.OpenFile("./logs.log", os.O_APPEND|os.O_CREATE, os.ModePerm)
	assert.NoError(err)

	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime)
	client := NewClient(utils.GetEnv("API_KEY"), logger)

	assert.NotEmpty(client)
	assert.NotNil(client.Logger)
	assert.FileExists("./logs.log")
}

func TestConnectToTheClientWithoutLoggingSuccess(t *testing.T) {
	assert := assert.New(t)

	client := NewClient(utils.GetEnv("API_KEY"), nil)
	assert.NotEmpty(client)
	assert.Nil(client.Logger)

	err := client.Connect()
	assert.NoError(err)
}

func TestConnectToTheClientWithLoggingSuccess(t *testing.T) {
	assert := assert.New(t)

	file, err := os.OpenFile("./logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	assert.NoError(err)

	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime)
	client := NewClient(utils.GetEnv("API_KEY"), logger)

	assert.NotEmpty(client)
	assert.NotNil(client.Logger)
	assert.FileExists("./logs.log")
	fmt.Println(client.Logger != nil)

	err = client.Connect()
	assert.NoError(err)
	assert.Condition(func() bool {
		content, err := ioutil.ReadFile("./logs.log")
		assert.NoError(err)

		stringifiedContent := string(content)
		return stringifiedContent != ""
	})
}

func TestFetchCurrentUser(t *testing.T) {
	assert := assert.New(t)

	client := NewClient(utils.GetEnv("API_KEY"), nil)
	assert.NotEmpty(client)
	assert.Nil(client.Logger)

	err := client.Connect()
	assert.NoError(err)

	currentUser, err := client.FetchCurrentUser()
	assert.NoError(err)
	assert.Equal(utils.GetEnv("DISCORD_USER_ID"), currentUser.Id)
	assert.NotEmpty(client.Cache.Users)
	assert.Equal(currentUser, client.Cache.Users[0])
}
