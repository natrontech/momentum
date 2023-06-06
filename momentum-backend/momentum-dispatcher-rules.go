package main

import "github.com/pocketbase/pocketbase/core"

const TABLE_NAME_APPLICATIONS = "applications"
const TABLE_NAME_DEPLOYMENTS = "deployments"
const TABLE_NAME_REPOSITORIES = "repositories"
const TABLE_NAME_STAGES = "stages"
const TABLE_NAME_TEMPLATES = "templates"
const TABLE_NAME_REPOSITORY_CREDENTIALS = "repositoryCredentials"
const TABLE_NAME_HELM_REPOSITORIES = "helmRepositories"
const TABLE_NAME_HELM_REPOSITORY_CREDENTIALS = "helmRepositoryCredentials"
const TABLE_NAME_HOOKS = "hooks"
const TABLE_NAME_KEYVALUE = "keyValue"
const TABLE_NAME_SECRET_KEY_VALUES = "secretKeyValues"

func setupCreateRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{TABLE_NAME_APPLICATIONS, dummyAction},
	}
}

func setupUpdateRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{TABLE_NAME_APPLICATIONS, dummyAction},
	}
}

func setupDeleteRules() []*MomentumDispatcherRule {
	return []*MomentumDispatcherRule{
		{TABLE_NAME_APPLICATIONS, dummyAction},
	}
}

func dummyAction(e *core.ModelEvent) error {
	return nil
}
