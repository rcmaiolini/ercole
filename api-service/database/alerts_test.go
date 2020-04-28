// Copyright (c) 2020 Sorint.lab S.p.A.
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

package database

import (
	"context"
	"testing"

	"github.com/amreo/ercole-services/model"
	"github.com/amreo/ercole-services/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongodbSuite) TestSearchAlerts() {
	defer m.db.Client.Database(m.dbname).Collection("alerts").DeleteMany(context.TODO(), bson.M{})
	m.InsertAlert(model.Alert{
		ID:            utils.Str2oid("5ea6a65bb2e36eb58da2f67c"),
		AlertCode:     model.AlertCodeNewOption,
		AlertSeverity: model.AlertSeverityCritical,
		AlertStatus:   model.AlertStatusNew,
		Date:          utils.P("2020-04-15T08:46:58.475+02:00"),
		Description:   "The database ERCOLE on test-db has enabled new features (Diagnostics Pack) on server",
		OtherInfo: map[string]interface{}{
			"Dbname": "ERCOLE",
			"Features": []interface{}{
				"Diagnostics Pack",
			},
			"Hostname": "test-db",
		},
	})
	m.InsertAlert(model.Alert{
		ID:            utils.Str2oid("5e96ade270c184faca93fe1b"),
		AlertCode:     model.AlertCodeNewServer,
		AlertSeverity: model.AlertSeverityNotice,
		AlertStatus:   model.AlertStatusAck,
		Date:          utils.P("2020-04-10T08:46:58.38+02:00"),
		Description:   "The server 'rac1_x' was added to ercole",
		OtherInfo: map[string]interface{}{
			"Hostname": "rac1_x",
		},
	})

	m.T().Run("should_be_paging", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{}, "", false, 0, 1, "", "", utils.MIN_TIME, utils.MAX_TIME)
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{
			map[string]interface{}{
				"Content": []interface{}{
					map[string]interface{}{
						"_id":           utils.Str2oid("5ea6a65bb2e36eb58da2f67c"),
						"AlertCode":     model.AlertCodeNewOption,
						"AlertSeverity": model.AlertSeverityCritical,
						"AlertStatus":   model.AlertStatusNew,
						"Date":          utils.P("2020-04-15T08:46:58.475+02:00"),
						"Description":   "The database ERCOLE on test-db has enabled new features (Diagnostics Pack) on server",
						"OtherInfo": map[string]interface{}{
							"Dbname": "ERCOLE",
							"Features": []interface{}{
								"Diagnostics Pack",
							},
							"Hostname": "test-db",
						},
						"Hostname": "test-db",
					},
				},
				"Metadata": map[string]interface{}{
					"Empty":         false,
					"First":         true,
					"Last":          false,
					"Number":        0,
					"Size":          1,
					"TotalElements": 2,
					"TotalPages":    2,
				},
			},
		}

		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})
	m.T().Run("should_be_sorting", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{}, "AlertSeverity", true, -1, -1, "", "", utils.MIN_TIME, utils.MAX_TIME)
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{
			map[string]interface{}{
				"_id":           utils.Str2oid("5e96ade270c184faca93fe1b"),
				"AlertCode":     model.AlertCodeNewServer,
				"AlertSeverity": model.AlertSeverityNotice,
				"AlertStatus":   model.AlertStatusAck,
				"Date":          utils.P("2020-04-10T08:46:58.38+02:00"),
				"Description":   "The server 'rac1_x' was added to ercole",
				"OtherInfo": map[string]interface{}{
					"Hostname": "rac1_x",
				},
				"Hostname": "rac1_x",
			},
			map[string]interface{}{
				"_id":           utils.Str2oid("5ea6a65bb2e36eb58da2f67c"),
				"AlertCode":     model.AlertCodeNewOption,
				"AlertSeverity": model.AlertSeverityCritical,
				"AlertStatus":   model.AlertStatusNew,
				"Date":          utils.P("2020-04-15T08:46:58.475+02:00"),
				"Description":   "The database ERCOLE on test-db has enabled new features (Diagnostics Pack) on server",
				"OtherInfo": map[string]interface{}{
					"Dbname": "ERCOLE",
					"Features": []interface{}{
						"Diagnostics Pack",
					},
					"Hostname": "test-db",
				},
				"Hostname": "test-db",
			},
		}
		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})
	m.T().Run("should_filter_by_status", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{}, "", false, -1, -1, "", model.AlertStatusNew, utils.MIN_TIME, utils.MAX_TIME)
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{
			map[string]interface{}{
				"_id":           utils.Str2oid("5ea6a65bb2e36eb58da2f67c"),
				"AlertCode":     model.AlertCodeNewOption,
				"AlertSeverity": model.AlertSeverityCritical,
				"AlertStatus":   model.AlertStatusNew,
				"Date":          utils.P("2020-04-15T08:46:58.475+02:00"),
				"Description":   "The database ERCOLE on test-db has enabled new features (Diagnostics Pack) on server",
				"OtherInfo": map[string]interface{}{
					"Dbname": "ERCOLE",
					"Features": []interface{}{
						"Diagnostics Pack",
					},
					"Hostname": "test-db",
				},
				"Hostname": "test-db",
			},
		}

		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})
	m.T().Run("should_filter_by_severity", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{}, "", false, -1, -1, model.AlertSeverityCritical, "", utils.MIN_TIME, utils.MAX_TIME)
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{
			map[string]interface{}{
				"_id":           utils.Str2oid("5ea6a65bb2e36eb58da2f67c"),
				"AlertCode":     model.AlertCodeNewOption,
				"AlertSeverity": model.AlertSeverityCritical,
				"AlertStatus":   model.AlertStatusNew,
				"Date":          utils.P("2020-04-15T08:46:58.475+02:00"),
				"Description":   "The database ERCOLE on test-db has enabled new features (Diagnostics Pack) on server",
				"OtherInfo": map[string]interface{}{
					"Dbname": "ERCOLE",
					"Features": []interface{}{
						"Diagnostics Pack",
					},
					"Hostname": "test-db",
				},
				"Hostname": "test-db",
			},
		}

		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})
	m.T().Run("should_filter_by_from", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{}, "", false, -1, -1, "", "", utils.P("2020-04-13T08:46:58.38+02:00"), utils.MAX_TIME)
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{
			map[string]interface{}{
				"_id":           utils.Str2oid("5ea6a65bb2e36eb58da2f67c"),
				"AlertCode":     model.AlertCodeNewOption,
				"AlertSeverity": model.AlertSeverityCritical,
				"AlertStatus":   model.AlertStatusNew,
				"Date":          utils.P("2020-04-15T08:46:58.475+02:00"),
				"Description":   "The database ERCOLE on test-db has enabled new features (Diagnostics Pack) on server",
				"OtherInfo": map[string]interface{}{
					"Dbname": "ERCOLE",
					"Features": []interface{}{
						"Diagnostics Pack",
					},
					"Hostname": "test-db",
				},
				"Hostname": "test-db",
			},
		}

		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})
	m.T().Run("should_filter_by_to", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{}, "", false, -1, -1, "", "", utils.MIN_TIME, utils.P("2020-04-13T08:46:58.38+02:00"))
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{
			map[string]interface{}{
				"_id":           utils.Str2oid("5e96ade270c184faca93fe1b"),
				"AlertCode":     model.AlertCodeNewServer,
				"AlertSeverity": model.AlertSeverityNotice,
				"AlertStatus":   model.AlertStatusAck,
				"Date":          utils.P("2020-04-10T08:46:58.38+02:00"),
				"Description":   "The server 'rac1_x' was added to ercole",
				"OtherInfo": map[string]interface{}{
					"Hostname": "rac1_x",
				},
				"Hostname": "rac1_x",
			},
		}

		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})
	m.T().Run("should_search1", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{"foobar"}, "", false, -1, -1, "", "", utils.MIN_TIME, utils.MAX_TIME)
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{}

		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})
	m.T().Run("should_search2", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{"added", model.AlertCodeNewServer, model.AlertSeverityNotice, "rac1_x"}, "", false, -1, -1, "", "", utils.MIN_TIME, utils.MAX_TIME)
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{
			map[string]interface{}{
				"_id":           utils.Str2oid("5e96ade270c184faca93fe1b"),
				"AlertCode":     model.AlertCodeNewServer,
				"AlertSeverity": model.AlertSeverityNotice,
				"AlertStatus":   model.AlertStatusAck,
				"Date":          utils.P("2020-04-10T08:46:58.38+02:00"),
				"Description":   "The server 'rac1_x' was added to ercole",
				"OtherInfo": map[string]interface{}{
					"Hostname": "rac1_x",
				},
				"Hostname": "rac1_x",
			},
		}

		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})

	m.T().Run("should_search3", func(t *testing.T) {
		out, err := m.db.SearchAlerts([]string{"ERCOLE", "Diagnostics Pack"}, "", false, -1, -1, "", "", utils.MIN_TIME, utils.MAX_TIME)
		m.Require().NoError(err)
		var expectedOut interface{} = []interface{}{
			map[string]interface{}{
				"_id":           utils.Str2oid("5ea6a65bb2e36eb58da2f67c"),
				"AlertCode":     model.AlertCodeNewOption,
				"AlertSeverity": model.AlertSeverityCritical,
				"AlertStatus":   model.AlertStatusNew,
				"Date":          utils.P("2020-04-15T08:46:58.475+02:00"),
				"Description":   "The database ERCOLE on test-db has enabled new features (Diagnostics Pack) on server",
				"OtherInfo": map[string]interface{}{
					"Dbname": "ERCOLE",
					"Features": []interface{}{
						"Diagnostics Pack",
					},
					"Hostname": "test-db",
				},
				"Hostname": "test-db",
			},
		}

		assert.JSONEq(m.T(), utils.ToJSON(expectedOut), utils.ToJSON(out))
	})
}
