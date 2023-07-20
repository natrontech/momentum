import type { RepositoriesResponse } from "$lib/pocketbase/generated-types";
import { writable, type Writable } from "svelte/store";

export const ActiveRepository : Writable<RepositoriesResponse> = writable()
