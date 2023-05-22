<script lang="ts">
  import BackButton from "$lib/components/base/BackButton.svelte";
  import type { RepositoriesResponse } from "$lib/pocketbase/generated-types";
  import type { PageData } from "./$types";
  export let data: PageData;

  $: repository = data.record as RepositoriesResponse;
</script>

<div class="">
  <BackButton />
  <h2 class="h2 text-primary-500 font-bold mb-6">
    {repository.name}
  </h2>
</div>
