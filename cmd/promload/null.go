// PromHouse
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/Percona-Lab/PromHouse/prompb"
)

// nullClient does nothing.
type nullClient struct {
	l *logrus.Entry
}

func newNullClient() *nullClient {
	return &nullClient{
		l: logrus.WithField("client", fmt.Sprintf("null")),
	}
}

func (client *nullClient) writeTS(ts []*prompb.TimeSeries) error {
	return nil
}

func (client *nullClient) close() error {
	return nil
}

// check interfaces
var (
	_ tsWriter = (*nullClient)(nil)
)
