package views

import "sesi8-latihan/server/models"

type GetAllPeopleSwagger struct {
	Response
	Payload []models.Person
}

type CreatePeopleSwagger struct {
	Response
	Payload models.Person
}

type UpdatePeopleSwagger struct {
	Response
	Payload models.Person
}

type DeletePeopleSwagger struct {
	Response
}
