package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceSpannerInstance() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(ResourceSpannerInstance().Schema)

	addRequiredFieldsToSchema(dsSchema, "name")
	addOptionalFieldsToSchema(dsSchema, "config")       // not sure why this is configurable
	addOptionalFieldsToSchema(dsSchema, "display_name") // not sure why this is configurable
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceSpannerInstanceRead,
		Schema: dsSchema,
	}
}

func dataSourceSpannerInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	id, err := ReplaceVars(d, config, "{{project}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceSpannerInstanceRead(d, meta)
}
