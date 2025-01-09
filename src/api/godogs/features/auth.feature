Feature: Authentication
  Scenario: signup request
    When I send a signup request to "/api/signup" with random data
    Then the response code should be 201
    # And the response should contain a "token"
