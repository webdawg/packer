// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package vagrant

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildGroupName         *string                `mapstructure:"packer_build_group_name" cty:"packer_build_group_name" hcl:"packer_build_group_name"`
	PackerBuildName              *string                `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType            *string                `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug                  *bool                  `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce                  *bool                  `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError                *string                `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars               map[string]string      `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars          []string               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	CompressionLevel             *int                   `mapstructure:"compression_level" cty:"compression_level" hcl:"compression_level"`
	Include                      []string               `mapstructure:"include" cty:"include" hcl:"include"`
	OutputPath                   *string                `mapstructure:"output" cty:"output" hcl:"output"`
	Override                     map[string]interface{} `cty:"override" hcl:"override"`
	VagrantfileTemplate          *string                `mapstructure:"vagrantfile_template" cty:"vagrantfile_template" hcl:"vagrantfile_template"`
	VagrantfileTemplateGenerated *bool                  `mapstructure:"vagrantfile_template_generated" cty:"vagrantfile_template_generated" hcl:"vagrantfile_template_generated"`
	ProviderOverride             *string                `mapstructure:"provider_override" cty:"provider_override" hcl:"provider_override"`
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
		"packer_build_group_name":        &hcldec.AttrSpec{Name: "packer_build_group_name", Type: cty.String, Required: false},
		"packer_build_name":              &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":            &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                   &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                   &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":          &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":     &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"compression_level":              &hcldec.AttrSpec{Name: "compression_level", Type: cty.Number, Required: false},
		"include":                        &hcldec.AttrSpec{Name: "include", Type: cty.List(cty.String), Required: false},
		"output":                         &hcldec.AttrSpec{Name: "output", Type: cty.String, Required: false},
		"override":                       &hcldec.AttrSpec{Name: "override", Type: cty.Map(cty.String), Required: false},
		"vagrantfile_template":           &hcldec.AttrSpec{Name: "vagrantfile_template", Type: cty.String, Required: false},
		"vagrantfile_template_generated": &hcldec.AttrSpec{Name: "vagrantfile_template_generated", Type: cty.Bool, Required: false},
		"provider_override":              &hcldec.AttrSpec{Name: "provider_override", Type: cty.String, Required: false},
	}
	return s
}
