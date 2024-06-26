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
	"sync"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/ercole-io/ercole/v2/model"
)

func (job *GcpDataRetrieveJob) GetDisk(ctx context.Context, projectID, diskname, zone string) (*computepb.Disk, error) {
	c, err := compute.NewDisksRESTClient(ctx, *job.Opt)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	req := &computepb.GetDiskRequest{
		Project: projectID,
		Disk:    diskname,
		Zone:    zone,
	}

	return c.Get(ctx, req)
}

func (job *GcpDataRetrieveJob) FetchGcpStorageDisk(ctx context.Context, gcpDisk model.GcpDisk, seqValue uint64, wg *sync.WaitGroup, ch chan<- model.GcpRecommendation) {
	defer wg.Done()

	maxReadIops, err := job.IsMaxReadIopsStorageOptimizable(ctx, gcpDisk)
	if err != nil {
		job.Log.Error(err)
		return
	}

	maxWriteIops, err := job.IsMaxWriteIopsStorageOptimizable(ctx, gcpDisk)
	if err != nil {
		job.Log.Error(err)
		return
	}

	maxReadThroughput, err := job.IsMaxThroughputReadStorageOptimizable(ctx, gcpDisk)
	if err != nil {
		job.Log.Error(err)
		return
	}

	maxWriteThroughput, err := job.IsMaxThroughputWriteStorageOptimizable(ctx, gcpDisk)
	if err != nil {
		job.Log.Error(err)
		return
	}

	optimizable := maxReadIops || maxWriteIops || maxReadThroughput || maxWriteThroughput

	if optimizable {
		ch <- model.GcpRecommendation{
			SeqValue:    seqValue,
			CreatedAt:   time.Now(),
			ProfileID:   gcpDisk.ProfileID,
			InstanceID:  gcpDisk.InstanceID,
			Category:    "Storage",
			Suggestion:  "Resize Oversized Disk",
			ProjectID:   gcpDisk.ProjectId,
			ProjectName: gcpDisk.Project.Name,
			ObjectType:  "Disk",
			Details:     map[string]string{},
		}
	}
}
