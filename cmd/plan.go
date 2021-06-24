package main

import "github.com/gruntwork-io/terratest/modules/terraform"

func InitAndPlan(options *terraform.Options) string {
	out, err := InitAndPlanE(options)
	if err != nil {
		return ""
	}

	return out
}

func InitAndPlanE(options *terraform.Options) (string, error) {
	if _, err := InitE(options); err != nil {
		return "", err
	}

	return PlanE(options)
}

func Plan(options *terraform.Options) string {
	out, err := PlanE(options)
	if err != nil {
		return ""
	}

	return out
}

func PlanE(options *terraform.Options) (string, error) {
	return RunTerraformCommandE(options, terraform.FormatArgs(options, "plan", "-input=false", "-lock=false")...)
}

func TgPlanAll(options *terraform.Options) string {
	out, err := TgPlanAllE(options)
	if err != nil {
		return ""
	}

	return out
}

func TgPlanAllE(options *terraform.Options) (string, error) {
	if options.TerraformBinary != "terragrunt" {
		logger.Fatalf("terragrunt must be set as the TerraformBinary to use this method")
	}

	return RunTerraformCommandE(options, terraform.FormatArgs(options, "run-all", "plan", "--input=false", "--lock=true")...)
}