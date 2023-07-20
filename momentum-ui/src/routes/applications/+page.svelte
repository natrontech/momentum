<script lang="ts">
    import { metadata } from "$lib/stores/metadata";
    import { createApplicationClient } from "$lib/components/client-factories/client-factory";
    import type { ApplicationsApi } from "$lib/momentum-core-client/api";
    import TableViewSelector from "$lib/components/base/TableViewSelector.svelte";
    import { onDestroy, onMount } from "svelte";
    import { ActiveRepository } from "$lib/stores/activeRepository";
  
    $metadata.title = "Applications";

    $: repository = $ActiveRepository;
    let appService: ApplicationsApi = createApplicationClient(repository);

    appService.

    onMount(async () => {
      
    });
  
    onDestroy(() => {
        
    });
  
  </script>
  
  <TableViewSelector />
  
  <!-- <div class="grid gap-6 {$tableView ? 'grid-cols-2' : ''}">
    {#each repositories as repository}
      <RepositoryCard {repository} />
    {/each}
  </div> -->
  