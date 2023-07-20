<script lang="ts">
  import BackButton from "$lib/components/base/BackButton.svelte";
    import { RepositoryClientFactory } from "$lib/components/client-factories/repository-client-factory";
  import type { RepositoriesResponse } from "$lib/pocketbase/generated-types";
  import { metadata } from "$lib/stores/metadata";
  import type { PageData } from "./$types";
  export let data: PageData;

  const repoPathBuilder = new RepositoryClientFactory();

  $: repository = data.record as RepositoriesResponse;

  $metadata.title = data.record.name;

</script>

<div class="">
  <BackButton />
  <h2 class="h2 text-primary-500 font-bold mb-6">
    {repository.name}
  </h2>

  <a href="{repoPathBuilder.buildBasePath(repository)}/swagger/index.html">{repoPathBuilder.buildBasePath(repository)}/swagger/index.html</a>
</div>
