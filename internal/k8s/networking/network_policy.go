/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package networking

import (
	"context"

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type networkPolicy struct {
	k8s.CommonNamespacedFields
	v1.NetworkPolicySpec
}

// NetworkPolicyColumns TODO
func NetworkPolicyColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&networkPolicy{})
}

// NetworkPoliciesGenerate TODO
func NetworkPoliciesGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		nps, err := k8s.GetClient().NetworkingV1().NetworkPolicies(metav1.NamespaceAll).List(context.TODO(), options)
		if err != nil {
			return nil, err
		}

		for _, np := range nps.Items {
			item := &networkPolicy{
				CommonNamespacedFields: k8s.GetCommonNamespacedFields(np.ObjectMeta),
				NetworkPolicySpec:      np.Spec,
			}
			results = append(results, k8s.ToMap(item))
		}

		if nps.Continue == "" {
			break
		}
		options.Continue = nps.Continue
	}

	return results, nil
}
