// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package digitaloceanimport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildGroupName *string           `mapstructure:"packer_build_group_name" cty:"packer_build_group_name" hcl:"packer_build_group_name"`
	PackerBuildName      *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType    *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug          *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce          *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError        *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars       map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars  []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	APIToken             *string           `mapstructure:"api_token" cty:"api_token" hcl:"api_token"`
	SpacesKey            *string           `mapstructure:"spaces_key" cty:"spaces_key" hcl:"spaces_key"`
	SpacesSecret         *string           `mapstructure:"spaces_secret" cty:"spaces_secret" hcl:"spaces_secret"`
	SpacesRegion         *string           `mapstructure:"spaces_region" cty:"spaces_region" hcl:"spaces_region"`
	SpaceName            *string           `mapstructure:"space_name" cty:"space_name" hcl:"space_name"`
	ObjectName           *string           `mapstructure:"space_object_name" cty:"space_object_name" hcl:"space_object_name"`
	SkipClean            *bool             `mapstructure:"skip_clean" cty:"skip_clean" hcl:"skip_clean"`
	Tags                 []string          `mapstructure:"image_tags" cty:"image_tags" hcl:"image_tags"`
	Name                 *string           `mapstructure:"image_name" cty:"image_name" hcl:"image_name"`
	Description          *string           `mapstructure:"image_description" cty:"image_description" hcl:"image_description"`
	Distribution         *string           `mapstructure:"image_distribution" cty:"image_distribution" hcl:"image_distribution"`
	ImageRegions         []string          `mapstructure:"image_regions" cty:"image_regions" hcl:"image_regions"`
	Timeout              *string           `mapstructure:"timeout" cty:"timeout" hcl:"timeout"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_group_name":    &hcldec.AttrSpec{Name: "packer_build_group_name", Type: cty.String, Required: false},
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"api_token":                  &hcldec.AttrSpec{Name: "api_token", Type: cty.String, Required: false},
		"spaces_key":                 &hcldec.AttrSpec{Name: "spaces_key", Type: cty.String, Required: false},
		"spaces_secret":              &hcldec.AttrSpec{Name: "spaces_secret", Type: cty.String, Required: false},
		"spaces_region":              &hcldec.AttrSpec{Name: "spaces_region", Type: cty.String, Required: false},
		"space_name":                 &hcldec.AttrSpec{Name: "space_name", Type: cty.String, Required: false},
		"space_object_name":          &hcldec.AttrSpec{Name: "space_object_name", Type: cty.String, Required: false},
		"skip_clean":                 &hcldec.AttrSpec{Name: "skip_clean", Type: cty.Bool, Required: false},
		"image_tags":                 &hcldec.AttrSpec{Name: "image_tags", Type: cty.List(cty.String), Required: false},
		"image_name":                 &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_description":          &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"image_distribution":         &hcldec.AttrSpec{Name: "image_distribution", Type: cty.String, Required: false},
		"image_regions":              &hcldec.AttrSpec{Name: "image_regions", Type: cty.List(cty.String), Required: false},
		"timeout":                    &hcldec.AttrSpec{Name: "timeout", Type: cty.String, Required: false},
	}
	return s
}
