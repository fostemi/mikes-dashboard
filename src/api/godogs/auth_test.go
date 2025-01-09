package main

import (
  "github.com/cucumber/godog"
  "github.com/go-resty/resty/v2"
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
func InitializeScenario(ctx *godog.ScenarioContext) {
  ctx.Step()
}
