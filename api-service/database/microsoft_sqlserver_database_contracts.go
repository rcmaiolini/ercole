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

package database

import (
	"context"

	"github.com/ercole-io/ercole/v2/model"
	"github.com/ercole-io/ercole/v2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const sqlServerDbContractsCollection = "ms_sqlserver_database_contracts"

func (md *MongoDatabase) InsertSqlServerDatabaseContract(contract model.SqlServerDatabaseContract) (*model.SqlServerDatabaseContract, error) {
	res, err := md.Client.Database(md.Config.Mongodb.DBName).Collection(sqlServerDbContractsCollection).
		InsertOne(context.TODO(), contract)
	if err != nil {
		return nil, utils.NewError(err, "DB ERROR")
	}

	contract.ID = res.InsertedID.(primitive.ObjectID)

	return &contract, nil
}

func (md *MongoDatabase) UpdateSqlServerDatabaseContract(contract model.SqlServerDatabaseContract) error {
	result, err := md.Client.Database(md.Config.Mongodb.DBName).Collection(sqlServerDbContractsCollection).
		ReplaceOne(context.TODO(), bson.M{
			"_id": contract.ID,
		}, contract)
	if err != nil {
		return utils.NewError(err, "DB ERROR")
	}

	if result.MatchedCount != 1 {
		return utils.ErrContractNotFound
	}

	return nil
}

func (md *MongoDatabase) RemoveSqlServerDatabaseContract(id primitive.ObjectID) error {
	res, err := md.Client.Database(md.Config.Mongodb.DBName).Collection(sqlServerDbContractsCollection).
		DeleteOne(context.TODO(), bson.M{
			"_id": id,
		})
	if err != nil {
		return utils.NewError(err, "DB ERROR")
	}

	if res.DeletedCount == 0 {
		return utils.ErrContractNotFound
	}

	return nil
}

func (md *MongoDatabase) ListSqlServerDatabaseContracts(locations []string) ([]model.SqlServerDatabaseContract, error) {
	ctx := context.TODO()
	out := make([]model.SqlServerDatabaseContract, 0)

	cur, err := md.Client.Database(md.Config.Mongodb.DBName).Collection(sqlServerDbContractsCollection).
		Aggregate(ctx, filterExistingLocations(locations))
	if err != nil {
		return nil, utils.NewError(err, "DB ERROR")
	}

	if err = cur.All(ctx, &out); err != nil {
		return nil, utils.NewError(err, "Decode ERROR")
	}

	return out, nil
}
