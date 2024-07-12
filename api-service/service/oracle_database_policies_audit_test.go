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

import (
	"testing"

	"github.com/ercole-io/ercole/v2/api-service/dto"
	"github.com/ercole-io/ercole/v2/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetOracleDatabasePoliciesAuditFlag_Green(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := NewMockMongoDatabaseInterface(mockCtrl)
	as := APIService{
		Database: db,
		Config: config.Configuration{
			APIService: config.APIService{
				OracleDatabasePoliciesAudit: []string{
					"policy0",
				},
			},
		},
	}

	expected := map[string][]string{"GREEN": {"policy0"}}

	db.EXPECT().FindOracleDatabasePoliciesAudit("hostname", "dbname").Return(
		&dto.OraclePoliciesAudit{List: []string{"policy0", "policy1"}}, nil)

	res, err := as.GetOracleDatabasePoliciesAuditFlag("hostname", "dbname")

	require.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestGetOracleDatabasePoliciesAuditFlag_Red(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := NewMockMongoDatabaseInterface(mockCtrl)
	as := APIService{
		Database: db,
		Config: config.Configuration{
			APIService: config.APIService{
				OracleDatabasePoliciesAudit: []string{
					"policy0",
				},
			},
		},
	}

	expected := map[string][]string{"RED": {"policy0"}}

	db.EXPECT().FindOracleDatabasePoliciesAudit("hostname", "dbname").Return(
		&dto.OraclePoliciesAudit{List: []string{"policy1"}}, nil)

	res, err := as.GetOracleDatabasePoliciesAuditFlag("hostname", "dbname")

	require.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestGetOracleDatabasePoliciesAuditFlag_NotAvailable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := NewMockMongoDatabaseInterface(mockCtrl)
	as := APIService{
		Database: db,
		Config:   config.Configuration{},
	}

	expected := map[string][]string{"N/A": nil}

	db.EXPECT().FindOracleDatabasePoliciesAudit("hostname", "dbname").Return(
		&dto.OraclePoliciesAudit{List: []string{"policy1"}}, nil)

	res, err := as.GetOracleDatabasePoliciesAuditFlag("hostname", "dbname")

	require.NoError(t, err)
	assert.Equal(t, expected, res)
}
