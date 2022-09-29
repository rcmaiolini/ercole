// Copyright (c) 2022 Sorint.lab S.p.A.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package service is a package that provides methods for querying data
package service

import (
	"github.com/ercole-io/ercole/v2/model"
)

func (as *APIService) InsertRole(role model.Role) (*model.Role, error) {
	err := as.Database.InsertRole(role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (as *APIService) UpdateRole(role model.Role) (*model.Role, error) {
	if err := as.Database.UpdateRole(role); err != nil {
		return nil, err
	}

	return &role, nil
}

func (as *APIService) GetRole(name string) (*model.Role, error) {
	role, err := as.Database.GetRole(name)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (as *APIService) GetRoles() ([]model.Role, error) {
	roles, err := as.Database.GetRoles()
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (as *APIService) DeleteRole(name string) error {
	if err := as.Database.DeleteRole(name); err != nil {
		return err
	}

	return nil
}
