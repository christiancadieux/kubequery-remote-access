/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package discovery

import (
	"context"

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	"k8s.io/apimachinery/pkg/version"
)

type info struct {
	version.Info
}

// InfoColumns TODO
func InfoColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&info{})
}

// InfoGenerate TODO
func InfoGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	results := make([]map[string]string, 0)

	sv, err := k8s.GetClient().Discovery().ServerVersion()
	if err != nil {
		return nil, err
	}

	item := &info{
		Info: *sv,
	}
	results = append(results, k8s.ToMap(item))

	return results, nil
}
