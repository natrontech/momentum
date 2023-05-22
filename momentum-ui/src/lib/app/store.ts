import { writable } from "svelte/store";

export const toggleTableView = writable(false);

export interface Metadata {
    title?: string;
    description?: string;
}

export const metadata = writable<Metadata>({});
