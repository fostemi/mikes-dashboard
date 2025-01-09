package main

import (
  "testing"

  "os"
  "fmt"
  "time"
  "context"
  "encoding/json"

  . "github.com/onsi/gomega"

  "github.com/cucumber/godog"
  "github.com/cucumber/godog/colors"
  "github.com/go-faker/faker/v4"
  "github.com/go-resty/resty/v2"
  "github.com/spf13/pflag"
)

type apiFeature struct {
  client *resty.Client
  response Response
  statusCode int
  authToken string
  err error
  user User
}

func (a *apiFeature) resetResponse(interface{}) {
  a.client = resty.New()
  a.response = Response{}
  a.statusCode = 0
  a.authToken = ""
  a.err = nil
  a.user = User{}
}

func (a *apiFeature) sendSignUpRequest(endpoint string) {
  password := faker.Password()
  signUpInput := SignUpRequest{
    Username: faker.Username(),
    Email: faker.Email(),
    Password: password,
    ConfirmPassword: password,
  }

  req := a.client.R().
    SetHeader("Content-Type", "application/json").
    SetBody(signUpInput)

  var res *resty.Response

  url := getUrl() + endpoint

  res, a.err = req.Post(url)
  Expect(a.err).To(BeNil())
  a.statusCode = res.StatusCode()
  a.user.Username = signUpInput.Username
  a.user.Password = password
}

func (a *apiFeature) theResponseCodeShouldBe(expectedStatusCode int) {
  Expect(a.statusCode).To(Equal(expectedStatusCode))
}

func (a *apiFeature) theResponseShouldContainA(key string) {
  res, err := json.Marshal(a.response)
  Expect(err).To(BeNil())

  actual := trimString(string(res))
  Expect(actual).To(ContainSubstring(key))
}

func InitializeScenario(ctx *godog.ScenarioContext) {
  api := &apiFeature{}
  ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
    api.resetResponse(sc)
    return ctx, nil
  })

  ctx.Step(`^I send a signup request to "([^"]*)" with random data$`, api.sendSignUpRequest)
  ctx.Step(`^the response code should be "([^"]*)"$`, api.theResponseCodeShouldBe)
  ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
  ctx.Step(`^the response should contain a"([^"]*)"$`, api.theResponseShouldContainA)
  ctx.Step(`^the response should contain a$`, api.theResponseShouldContainA)
}

var opts = godog.Options{
  Paths: []string{"features"},
  Output: colors.Colored(os.Stdout),
  Randomize: time.Now().UTC().UnixNano(),
  Format: "pretty",
}

func init() {
  godog.BindCommandLineFlags("godog.", &opts)
}

func TestFeatures(t *testing.T) {
  fmt.Println("Testing")
  pflag.Parse()

  RegisterFailHandler(func(message string, _ ...int){
    panic(message)
  })

  status := godog.TestSuite{
    ScenarioInitializer: InitializeScenario,
    Options: &opts,
  }

  if status.Run() != 0 {
    t.Fatal("non-zero status returned, failed to run feature tests")
  }
}
