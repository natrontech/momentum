/**
* This file was @generated using pocketbase-typegen
*/

export enum Collections {
	Hooks = "hooks",
	Repositories = "repositories",
	RepositoryCredentials = "repositoryCredentials",
	Users = "users",
}

// Alias types for improved usability
export type IsoDateString = string
export type RecordIdString = string
export type HTMLString = string

// System fields
export type BaseSystemFields<T = never> = {
	id: RecordIdString
	created: IsoDateString
	updated: IsoDateString
	collectionId: string
	collectionName: Collections
	expand?: T
}

export type AuthSystemFields<T = never> = {
	email: string
	emailVisibility: boolean
	username: string
	verified: boolean
} & BaseSystemFields<T>

// Record types for each collection

export enum HooksEventOptions {
	"insert" = "insert",
	"update" = "update",
	"delete" = "delete",
}

export enum HooksActionTypeOptions {
	"command" = "command",
	"post" = "post",
}
export type HooksRecord = {
	collection: string
	event: HooksEventOptions
	action_type: HooksActionTypeOptions
	action: string
	action_params?: string
	expands?: string
	disabled?: boolean
}

export enum RepositoriesStatusOptions {
	"PENDING" = "PENDING",
	"SYNCING" = "SYNCING",
	"UP-TO-DATE" = "UP-TO-DATE",
	"ERROR" = "ERROR",
}
export type RepositoriesRecord = {
	name: string
	url: string
	status: RepositoriesStatusOptions
	repositoryCredentials?: RecordIdString
}

export type RepositoryCredentialsRecord = {
	username?: string
	password?: string
}

export type UsersRecord = {
	name?: string
	avatar?: string
}

// Response types include system fields and match responses from the PocketBase API
export type HooksResponse = Required<HooksRecord> & BaseSystemFields
export type RepositoriesResponse<Texpand = unknown> = Required<RepositoriesRecord> & BaseSystemFields<Texpand>
export type RepositoryCredentialsResponse = Required<RepositoryCredentialsRecord> & BaseSystemFields
export type UsersResponse = Required<UsersRecord> & AuthSystemFields

// Types containing all Records and Responses, useful for creating typing helper functions

export type CollectionRecords = {
	hooks: HooksRecord
	repositories: RepositoriesRecord
	repositoryCredentials: RepositoryCredentialsRecord
	users: UsersRecord
}

export type CollectionResponses = {
	hooks: HooksResponse
	repositories: RepositoriesResponse
	repositoryCredentials: RepositoryCredentialsResponse
	users: UsersResponse
}
