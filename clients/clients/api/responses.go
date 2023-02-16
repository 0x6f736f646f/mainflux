package api

import (
	"fmt"
	"net/http"

	"github.com/mainflux/mainflux"
	"github.com/mainflux/mainflux/clients/clients"
)

var (
	_ mainflux.Response = (*tokenRes)(nil)
	_ mainflux.Response = (*viewClientRes)(nil)
	_ mainflux.Response = (*createClientRes)(nil)
	_ mainflux.Response = (*createClientsRes)(nil)
	_ mainflux.Response = (*deleteClientRes)(nil)
	_ mainflux.Response = (*clientsPageRes)(nil)
	_ mainflux.Response = (*viewMembersRes)(nil)
	_ mainflux.Response = (*memberPageRes)(nil)
)

type pageRes struct {
	Limit  uint64 `json:"limit,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
	Total  uint64 `json:"total"`
	Level  uint64 `json:"level"`
	Name   string `json:"name"`
}

type createClientRes struct {
	clients.Client
	created bool
}

func (res createClientRes) Code() int {
	if res.created {
		return http.StatusCreated
	}

	return http.StatusOK
}

func (res createClientRes) Headers() map[string]string {
	if res.created {
		return map[string]string{
			"Location": fmt.Sprintf("/clients/%s", res.ID),
		}
	}

	return map[string]string{}
}

func (res createClientRes) Empty() bool {
	return false
}

type createClientsRes struct {
	clients []clients.Client
	created bool
}

func (res createClientsRes) Code() int {
	if res.created {
		return http.StatusCreated
	}

	return http.StatusOK
}

func (res createClientsRes) Headers() map[string]string {
	return map[string]string{}
}

func (res createClientsRes) Empty() bool {
	return false
}

type tokenRes struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessType   string `json:"access_type,omitempty"`
}

func (res tokenRes) Code() int {
	return http.StatusCreated
}

func (res tokenRes) Headers() map[string]string {
	return map[string]string{}
}

func (res tokenRes) Empty() bool {
	return res.AccessToken == "" || res.RefreshToken == ""
}

type updateClientRes struct {
	clients.Client
}

func (res updateClientRes) Code() int {
	return http.StatusOK
}

func (res updateClientRes) Headers() map[string]string {
	return map[string]string{}
}

func (res updateClientRes) Empty() bool {
	return false
}

type viewClientRes struct {
	clients.Client
}

func (res viewClientRes) Code() int {
	return http.StatusOK
}

func (res viewClientRes) Headers() map[string]string {
	return map[string]string{}
}

func (res viewClientRes) Empty() bool {
	return false
}

type clientsPageRes struct {
	pageRes
	Clients []viewClientRes `json:"clients"`
}

func (res clientsPageRes) Code() int {
	return http.StatusOK
}

func (res clientsPageRes) Headers() map[string]string {
	return map[string]string{}
}

func (res clientsPageRes) Empty() bool {
	return false
}

type viewMembersRes struct {
	clients.Client
}

func (res viewMembersRes) Code() int {
	return http.StatusOK
}

func (res viewMembersRes) Headers() map[string]string {
	return map[string]string{}
}

func (res viewMembersRes) Empty() bool {
	return false
}

type memberPageRes struct {
	pageRes
	Members []viewMembersRes `json:"members"`
}

func (res memberPageRes) Code() int {
	return http.StatusOK
}

func (res memberPageRes) Headers() map[string]string {
	return map[string]string{}
}

func (res memberPageRes) Empty() bool {
	return false
}

type deleteClientRes struct {
	clients.Client
}

func (res deleteClientRes) Code() int {
	return http.StatusOK
}

func (res deleteClientRes) Headers() map[string]string {
	return map[string]string{}
}

func (res deleteClientRes) Empty() bool {
	return false
}
