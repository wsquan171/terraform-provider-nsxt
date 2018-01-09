/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type LogicalRouterStatusPerNode struct {

	// A service router's HA status on an edge node
	HighAvailabilityStatus string `json:"high_availability_status"`

	// id of the service router where the router status is retrieved.
	ServiceRouterId string `json:"service_router_id,omitempty"`

	// id of the transport node where the router status is retrieved.
	TransportNodeId string `json:"transport_node_id"`
}
