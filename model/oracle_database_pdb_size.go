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
package model

type OracleDatabasePdbSize struct {
	SegmentsSize       float64 `json:"segmentsSize" bson:"segmentsSize"`
	DatafileSize       float64 `json:"datafileSize" bson:"datafileSize"`
	Allocable          float64 `json:"allocable" bson:"allocable"`
	SGATarget          float64 `json:"sgaTarget" bson:"sgaTarget"`
	PGAAggregateTarget float64 `json:"pgaAggregateTarget" bson:"pgaAggregateTarget"`
}
