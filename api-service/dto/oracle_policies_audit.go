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
package dto

import (
	"github.com/ercole-io/ercole/v2/utils"
)

type OraclePoliciesAudit struct {
	List []string `json:"policiesAudit"`
}

func (opa OraclePoliciesAudit) Response(configured, policies []string) map[string][]string {
	if len(configured) == 0 {
		return map[string][]string{"N/A": nil}
	}

	isEnable := true

	for _, c := range configured {
		if !isEnable {
			break
		}

		isEnable = utils.Contains(policies, c)
	}

	s := []string{}

	if isEnable {
		for _, c := range configured {
			if utils.Contains(policies, c) {
				s = append(s, c)
			}
		}

		return map[string][]string{"GREEN": s}
	}

	for _, c := range configured {
		if !utils.Contains(policies, c) {
			s = append(s, c)
		}
	}

	return map[string][]string{"RED": s}
}
