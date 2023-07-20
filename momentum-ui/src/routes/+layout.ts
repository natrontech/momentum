import { client } from "$lib/pocketbase";
import type { RepositoriesResponse } from "$lib/pocketbase/generated-types";

// turn off SSR - we're JAMstack here
export const ssr = false;
// Prerendering turned off. Turn it on if you know what you're doing.
export const prerender = false;
// trailing slashes make relative paths much easier
export const trailingSlash = "always";

export const load = async ({url}) => {
    
    const pathname = url;
    const records: RepositoriesResponse[] = await client.collection("repositories").getFullList();

    return {
        pathname,
        records
    }
}
