package momentumconfig

const GENERIC_FIELD_KEYVALUES = "keyValues"

const TABLE_REPOSITORIES_NAME = "repositories"
const TABLE_REPOSITORIES_FIELD_ID = "id"
const TABLE_REPOSITORIES_FIELD_NAME = "name"
const TABLE_REPOSITORIES_FIELD_URL = "url"
const TABLE_REPOSITORIES_FIELD_APPLICATIONS = "applications"

const TABLE_APPLICATIONS_NAME = "applications"
const TABLE_APPLICATIONS_FIELD_ID = "id"
const TABLE_APPLICATIONS_FIELD_NAME = "name"
const TABLE_APPLICATIONS_FIELD_STAGES = "stages"
const TABLE_APPLICATIONS_FIELD_HELMREPO = "helmRepository"
const TABLE_APPLICATIONS_FIELD_PARENTREPOSITORY = "parentRepository"

const TABLE_DEPLOYMENTS_NAME = "deployments"
const TABLE_DEPLOYMENTS_FIELD_ID = "id"
const TABLE_DEPLOYMENTS_FIELD_NAME = "name"
const TABLE_DEPLOYMENTS_FIELD_DESCRIPTION = "description"
const TABLE_DEPLOYMENTS_FIELD_REPOSITORIES = "repositories"
const TABLE_DEPLOYMENTS_FIELD_PARENTSTAGE = "parentStage"

const TABLE_STAGES_NAME = "stages"
const TABLE_STAGES_FIELD_ID = "id"
const TABLE_STAGES_FIELD_NAME = "name"
const TABLE_STAGES_FIELD_DEPLOYMENTS = "deployments"
const TABLE_STAGES_FIELD_PARENTSTAGE = "parentStage"
const TABLE_STAGES_FIELD_PARENTAPPLICATION = "parentApplication"

const TABLE_TEMPLATES_NAME = "templates"
const TABLE_TEMPLATES_FIELD_ID = "id"
const TABLE_TEMPLATES_FIELD_NAME = "name"

const TABLE_REPOSITORYCREDENTIALS_NAME = "repositoryCredentials"

const TABLE_HELMREPOSITORIES_NAME = "helmRepositories"

const TABLE_HELMREPOSITORYCREDENTIALS_NAME = "helmRepositoryCredentials"

const TABLE_HOOKS_NAME = "hooks"

const TABLE_KEYVALUE_NAME = "keyValues"
const TABLE_KEYVALUE_FIELD_ID = "id"
const TABLE_KEYVALUE_FIELD_KEY = "key"
const TABLE_KEYVALUE_FIELD_VALUE = "value"
const TABLE_KEYVALUE_FIELD_PARENTSTAGE = "parentStage"
const TABLE_KEYVALUE_FIELD_PARENTDEPLOYMENT = "parentDeployment"

const TABLE_SECRETKEYVALUE_NAME = "secretKeyValues"
