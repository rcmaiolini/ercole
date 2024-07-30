// Copyright (c) 2024 Sorint.lab S.p.A.
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
package service

import "errors"

func (as *APIService) GetOracleDatabasePoliciesAuditFlag(hostname, dbname string) (map[string][]string, error) {
	exist, err := as.Database.DbExist(hostname, dbname)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, errors.New("no document found")
	}

	policiesAudit, err := as.Database.FindOracleDatabasePoliciesAudit(hostname, dbname)
	if err != nil {
		return nil, err
	}

	return policiesAudit.Response(as.Config.APIService.OracleDatabasePoliciesAudit, policiesAudit.List), err
}