package config

const API = "/api"
const API_VERSION_BETA = API + "/beta"

const API_FILE_BY_ID = API_VERSION_BETA + "/file/:id"
const API_FILE_ADD = API_VERSION_BETA + "/file"
const API_FILE_UPDATE = API_VERSION_BETA + "/file/:id"
const API_FILE_LINE_OVERWRITTENBY = API_VERSION_BETA + "/file/:id/line/:lineNumber/overwritten-by"

const API_ARTEFACT_BY_ID = API_VERSION_BETA + "/artefact/:id"
const API_ARTEFACT_APPLICATIONS = API_VERSION_BETA + "/applications"
const API_ARTEFACT_STAGES = API_VERSION_BETA + "/stages"
const API_ARTEFACT_DEPLOYMENTS = API_VERSION_BETA + "/deployments"

const API_TEMPLATES_PREFIX = API_VERSION_BETA + "/templates"
const API_TEMPLATES_APPLICATIONS = API_TEMPLATES_PREFIX + "/applications"
const API_TEMPLATES_STAGES = API_TEMPLATES_PREFIX + "/stages"
const API_TEMPLATES_DEPLOYMENTS = API_TEMPLATES_PREFIX + "/deployments"
