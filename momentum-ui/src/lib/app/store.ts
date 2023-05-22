import { writable } from "svelte/store";

export interface Metadata {
    title?: string;
    description?: string;
}

export const metadata = writable<Metadata>({});

export const toggleTableView = writable(false);

export const loading = writable(false);
