<script lang="ts">
  import { metadata } from "$lib/app/store";
  import type { RepositoriesRecord } from "$lib/pocketbase/generated-types";
  import type { PageData } from "./$types";

  $metadata.title = "Repositories";

  export let data: PageData;

  // parse data to RepositoriesRecord[]
  $: repositories = data.records as RepositoriesRecord[];

</script>

{#each repositories as repository}
  <div>
    <h1>{repository.name}</h1>
  </div>
{/each}
