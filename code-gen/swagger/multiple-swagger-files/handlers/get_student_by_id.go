package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/muly/golang/code-gen/swagger/multiple-swagger-files/gen/models"
	"github.com/muly/golang/code-gen/swagger/multiple-swagger-files/gen/restapi/operations"
)

func NewGetStudentByID() operations.GetStudentByIDHandler {
	return &getStudentByID{}
}

type getStudentByID struct{}

func (getStudentByID) Handle(operations.GetStudentByIDParams) middleware.Responder {

	s := models.StudentView{ID: 1, Name: "student 1"}

	return operations.NewGetStudentByIDOK().WithPayload(&s)
}
