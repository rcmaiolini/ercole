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
package job

import (
	"context"

	"google.golang.org/api/cloudresourcemanager/v1"
)

func (job *GcpDataRetrieveJob) Getprojects(ctx context.Context) ([]*cloudresourcemanager.Project, error) {
	client, err := cloudresourcemanager.NewService(ctx, *job.Opt)
	if err != nil {
		return nil, err
	}

	projects, err := client.Projects.List().Do()
	if err != nil {
		return nil, err
	}

	return projects.Projects, nil
}
