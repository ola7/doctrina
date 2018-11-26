package service

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"../dbclient"
	"../model"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/h2non/gock.v1"
)

func init() {
	gock.InterceptClient(client)
}

func TestGetUserWrongPath(t *testing.T) {

	Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})

}

func TestGetUser(t *testing.T) {

	// make sure http intercept is turned off after this test
	defer gock.Off()

	// turn on http intercept here
	gock.New("http://quotes-service:8080").
		Get("api/quote").MatchParam("strength", "4").
		Reply(200).
		BodyString(`{"quote":"May the source be with you. Always.","ipAddress":"10.0.0.5:8080","language":"en"}`)

	mockRepo := &dbclient.MockBoltClient{}

	mockRepo.On("QueryUser", "123").Return(model.User{Id: "123", Name: "Name_123"}, nil)
	mockRepo.On("QueryUser", "456").Return(model.User{}, fmt.Errorf("Some error"))

	DBClient = mockRepo

	Convey("Given a HTTP request for /users/123", t, func() {
		req := httptest.NewRequest("GET", "/users/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)

				user := model.User{}
				json.Unmarshal(resp.Body.Bytes(), &user)
				So(user.Id, ShouldEqual, "123")
				So(user.Name, ShouldEqual, "Name_123")
				So(user.Quote.Text, ShouldEqual, "May the source be with you. Always.")
			})
		})
	})

	Convey("Given a HTTP request for /users/456", t, func() {
		req := httptest.NewRequest("GET", "/users/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}
