<script lang="ts">
  import { metadata } from "$lib/stores/metadata";
  import type { RepositoriesResponse } from "$lib/pocketbase/generated-types";
  import type { PageData } from "./$types";
  import RepositoryCard from "$lib/components/repositories/RepositoryCard.svelte";
  import tableView from "$lib/stores/tableView";
  import TableViewSelector from "$lib/components/base/TableViewSelector.svelte";

  $metadata.title = "Repositories";

  export let data: PageData;

  // parse data to RepositoriesRecord[]
  $: repositories = data.records as RepositoriesResponse[];
</script>

<TableViewSelector />

<div class="grid gap-6 {$tableView ? 'grid-cols-2' : ''}">
  {#each repositories as repository}
    <RepositoryCard {repository} />
  {/each}
</div>
