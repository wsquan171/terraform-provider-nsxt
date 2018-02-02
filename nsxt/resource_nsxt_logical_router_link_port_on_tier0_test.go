/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/vmware/go-vmware-nsxt"
	"net/http"
	"testing"
)

func TestNSXLogicalRouterLinkPortOnTier0Basic(t *testing.T) {

	name := fmt.Sprintf("test-nsx-port-on-tier0")
	tier0RouterName := Tier0RouterDefaultName
	updateName := fmt.Sprintf("%s-update", name)
	testResourceName := "nsxt_logical_router_link_port_on_tier0.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNSXLogicalRouterLinkPortOnTier0CheckDestroy(state, name)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNSXLogicalRouterLinkPortOnTier0CreateTemplate(name, tier0RouterName),
				Check: resource.ComposeTestCheckFunc(
					testAccNSXLogicalRouterLinkPortOnTier0Exists(name, testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "display_name", name),
					resource.TestCheckResourceAttr(testResourceName, "description", "Acceptance Test"),
					resource.TestCheckResourceAttr(testResourceName, "tags.#", "2"),
				),
			},
			{
				Config: testAccNSXLogicalRouterLinkPortOnTier0UpdateTemplate(updateName, tier0RouterName),
				Check: resource.ComposeTestCheckFunc(
					testAccNSXLogicalRouterLinkPortOnTier0Exists(updateName, testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "display_name", updateName),
					resource.TestCheckResourceAttr(testResourceName, "description", "Acceptance Test Update"),
					resource.TestCheckResourceAttr(testResourceName, "tags.#", "1"),
				),
			},
		},
	})
}

func TestNSXLogicalRouterLinkPortOnTier0WithRelay(t *testing.T) {
	// Note: this test will pass only with NSX 2.2 & up. Before that dhcp relay is not supported here
	name := fmt.Sprintf("test-nsx-port-on-tier0")
	tier0RouterName := Tier0RouterDefaultName
	updateName := fmt.Sprintf("%s-update", name)
	testResourceName := "nsxt_logical_router_link_port_on_tier0.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccNSXLogicalRouterLinkPortOnTier0CheckDestroy(state, name)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNSXLogicalRouterLinkPortOnTier0WithRelayCreateTemplate(name, tier0RouterName),
				Check: resource.ComposeTestCheckFunc(
					testAccNSXLogicalRouterLinkPortOnTier0Exists(name, testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "display_name", name),
					resource.TestCheckResourceAttr(testResourceName, "description", "Acceptance Test"),
					resource.TestCheckResourceAttr(testResourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(testResourceName, "service_bindings.#", "1"),
					resource.TestCheckResourceAttr(testResourceName, "service_bindings.0.target_type", "LogicalService"),
					resource.TestCheckResourceAttr(testResourceName, "service_bindings.0.target_display_name", "srv"),
				),
			},
			{
				Config: testAccNSXLogicalRouterLinkPortOnTier0WithRelayUpdateTemplate(updateName, tier0RouterName),
				Check: resource.ComposeTestCheckFunc(
					testAccNSXLogicalRouterLinkPortOnTier0Exists(updateName, testResourceName),
					resource.TestCheckResourceAttr(testResourceName, "display_name", updateName),
					resource.TestCheckResourceAttr(testResourceName, "description", "Acceptance Test Update"),
					resource.TestCheckResourceAttr(testResourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(testResourceName, "service_bindings.#", "1"),
					resource.TestCheckResourceAttr(testResourceName, "service_bindings.0.target_type", "LogicalService"),
					resource.TestCheckResourceAttr(testResourceName, "service_bindings.0.target_display_name", "srv"),
				),
			},
		},
	})
}

func testAccNSXLogicalRouterLinkPortOnTier0Exists(displayName string, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		nsxClient := testAccProvider.Meta().(*nsxt.APIClient)

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("NSX logical router link port on Tier0 resource %s not found in resources", resourceName)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("NSX logical router link port on Tier0 resource ID not set in resources")
		}

		resource, responseCode, err := nsxClient.LogicalRoutingAndServicesApi.ReadLogicalRouterLinkPortOnTier0(nsxClient.Context, resourceID)
		if err != nil {
			return fmt.Errorf("Error while retrieving logical router link port on Tier0 ID %s. Error: %v", resourceID, err)
		}

		if responseCode.StatusCode != http.StatusOK {
			return fmt.Errorf("Error while checking verifying logical router link port on Tier0 existence. HTTP returned %d", resourceID, responseCode)
		}

		if displayName == resource.DisplayName {
			return nil
		}
		return fmt.Errorf("NSX logical router link port on Tier0 %s not found", displayName)
	}
}

func testAccNSXLogicalRouterLinkPortOnTier0CheckDestroy(state *terraform.State, displayName string) error {

	nsxClient := testAccProvider.Meta().(*nsxt.APIClient)
	for _, rs := range state.RootModule().Resources {

		if rs.Type != "nsxt_logical_router_link_port_on_tier0" {
			continue
		}

		resourceID := rs.Primary.Attributes["id"]
		resource, responseCode, err := nsxClient.LogicalRoutingAndServicesApi.ReadLogicalRouterLinkPortOnTier0(nsxClient.Context, resourceID)
		if err != nil {
			if responseCode.StatusCode != http.StatusOK {
				return nil
			}
			return fmt.Errorf("Error while retrieving logical router link port on Tier0 %s. Error: %v", resourceID, err)
		}

		if displayName == resource.DisplayName {
			return fmt.Errorf("NSX logical router link port on Tier0 %s still exists", displayName)
		}
	}
	return nil
}

func testAccNSXLogicalRouterLinkPortOnTier0CreateTemplate(name string, tier0RouterName string) string {
	return fmt.Sprintf(`
data "nsxt_logical_tier0_router" "TIER0RTR" {
    display_name = "%s"
}

resource "nsxt_logical_router_link_port_on_tier0" "test" {
	display_name = "%s"
	description = "Acceptance Test"
	logical_router_id = "${data.nsxt_logical_tier0_router.TIER0RTR.id}"
	tags = [{scope = "scope1"
        	 tag = "tag1"},
        	{scope = "scope2"
    	     tag = "tag2"}
	]
}`, tier0RouterName, name)
}

func testAccNSXLogicalRouterLinkPortOnTier0UpdateTemplate(name string, tier0RouterName string) string {
	return fmt.Sprintf(`
data "nsxt_logical_tier0_router" "TIER0RTR" {
    display_name = "%s"
}

resource "nsxt_logical_router_link_port_on_tier0" "test" {
	display_name = "%s"
	description = "Acceptance Test Update"
	logical_router_id = "${data.nsxt_logical_tier0_router.TIER0RTR.id}"
	tags = [{scope = "scope3"
    	     tag = "tag3"}
	]
}`, tier0RouterName, name)
}

func testAccNSXLogicalRouterLinkPortOnTier0WithRelayCreateTemplate(name string, tier0RouterName string) string {
	return testAccNSXLogicalRouterDownlinkPortRelayTemplate() + fmt.Sprintf(`
data "nsxt_logical_tier0_router" "TIER0RTR" {
    display_name = "%s"
}

resource "nsxt_logical_router_link_port_on_tier0" "test" {
	display_name = "%s"
	description = "Acceptance Test"
	logical_router_id = "${data.nsxt_logical_tier0_router.TIER0RTR.id}"
	service_bindings = [{target_id = "${nsxt_dhcp_relay_service.DRS1.id}"
			             target_type = "LogicalService"}
	]
	tags = [{scope = "scope1"
        	 tag = "tag1"},
        	{scope = "scope2"
    	     tag = "tag2"}
	]
}`, tier0RouterName, name)
}

func testAccNSXLogicalRouterLinkPortOnTier0WithRelayUpdateTemplate(name string, tier0RouterName string) string {
	return testAccNSXLogicalRouterDownlinkPortRelayTemplate() + fmt.Sprintf(`
data "nsxt_logical_tier0_router" "TIER0RTR" {
    display_name = "%s"
}

resource "nsxt_logical_router_link_port_on_tier0" "test" {
	display_name = "%s"
	description = "Acceptance Test Update"
	logical_router_id = "${data.nsxt_logical_tier0_router.TIER0RTR.id}"
	service_bindings = [{target_id = "${nsxt_dhcp_relay_service.DRS1.id}"
			             target_type = "LogicalService"}
	]
	tags = [{scope = "scope3"
    	     tag = "tag3"}
	]
}`, tier0RouterName, name)
}