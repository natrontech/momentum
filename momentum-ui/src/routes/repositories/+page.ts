import { client } from "$lib/pocketbase";
import type { RepositoriesRecord } from "$lib/pocketbase/generated-types";
import type { PageLoad } from "../$types";

export const load: PageLoad = async () => {
    const records: RepositoriesRecord[] = await client.collection("repositories").getFullList();
    return {
        records
    };
};
