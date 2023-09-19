package config

const API = "/api"
const API_VERSION_BETA = API + "/beta"

const API_FILE_BY_ID = API_VERSION_BETA + "/file/:id"
const API_DIR_BY_ID = API_VERSION_BETA + "/dir/:id"
const API_FILE_LINE_OVERWRITTENBY = API_VERSION_BETA + "/file/:id/line/:lineNumber/overwritten-by"
const API_FILE_LINE_OVERWRITES = API_VERSION_BETA + "/file/:id/line/:lineNumber/overwrites"

const API_ARTEFACT_BY_ID = API_VERSION_BETA + "/artefact/:id"
const API_APPLICATIONS = API_VERSION_BETA + "/applications"
const API_STAGES = API_VERSION_BETA + "/stages"
const API_DEPLOYMENTS = API_VERSION_BETA + "/deployments"
