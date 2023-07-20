/**
* This file was @generated using pocketbase-typegen
*/

export enum Collections {
	Repositories = "repositories",
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

export enum RepositoriesStatusOptions {
	"UP" = "UP",
	"ERROR" = "ERROR",
	"SYNC" = "SYNC",
}
export type RepositoriesRecord = {
	name?: string
	coreHost?: string
	corePort?: string
	coreBasePath?: string
	status?: RepositoriesStatusOptions
}

export type UsersRecord = {
	name?: string
	avatar?: string
}

// Response types include system fields and match responses from the PocketBase API
export type RepositoriesResponse = Required<RepositoriesRecord> & BaseSystemFields
export type UsersResponse = Required<UsersRecord> & AuthSystemFields

// Types containing all Records and Responses, useful for creating typing helper functions

export type CollectionRecords = {
	repositories: RepositoriesRecord
	users: UsersRecord
}

export type CollectionResponses = {
	repositories: RepositoriesResponse
	users: UsersResponse
}