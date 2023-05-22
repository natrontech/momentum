<script lang="ts">
  import {
    RepositoriesStatusOptions,
    type RepositoriesResponse
  } from "$lib/pocketbase/generated-types";
  import { RefreshCw, Ban, Check, Link, MenuSquare, Settings2 } from "lucide-svelte";
  export let repository: RepositoriesResponse;

  function onHandleStatusClick(event: any): void {
    console.log(event);
  }

</script>

<div class="card rounded-md border-2
  {
    repository.status === RepositoriesStatusOptions['UP-TO-DATE']
      ? 'border-green-300'
      : repository.status === RepositoriesStatusOptions.ERROR
      ? 'border-red-300'
      : repository.status === RepositoriesStatusOptions.PENDING
      ? 'border-yellow-300'
      : 'border-gray-300'
  }
">
  <div class="flex">
    <button
      on:click={onHandleStatusClick}
      class="p-6 rounded-l-md cursor-pointer w-52 group hover:w-full transition-all duration-300 ease-in-out
      {repository.status === RepositoriesStatusOptions['UP-TO-DATE']
        ? 'bg-green-300'
        : repository.status === RepositoriesStatusOptions.ERROR
        ? 'bg-red-300'
        : repository.status === RepositoriesStatusOptions.PENDING
        ? 'bg-yellow-300'
        : 'bg-gray-300'}
    "
    >

        <div class="flex flex-col justify-center items-center ">
          {#if repository.status === RepositoriesStatusOptions.PENDING}
          <RefreshCw />
          {:else if repository.status === RepositoriesStatusOptions.ERROR}
          <Ban />
          {:else if repository.status === RepositoriesStatusOptions['UP-TO-DATE']}
          <Check />
          {:else if repository.status === RepositoriesStatusOptions.SYNCING}
          <RefreshCw class="animate-spin" />
          {/if}
          <p class=" text-xs">
            {repository.status}
          </p>
        </div>
    </button>
    <div class="flex flex-col justify-center items-center w-full relative">
      <h3 class="text-lg leading-6 font-medium text-gray-900 ">
        {repository.name}
      </h3>
      <a href={"/repositories/" + repository.id} class="absolute right-2 top-2">
        <Settings2 class="text-gray-900" />
      </a>
    </div>
  </div>
</div>

<!-- <div class="absolute left-0 right-0 w-full top-0">
  {#if repository.status === RepositoriesStatusOptions["UP-TO-DATE"]}
    <span class="badge variant-filled bg-green-500 w-full rounded-t-md">
      <Check />{@html "&nbsp;"}
      {repository.status}
    </span>
  {:else if repository.status === RepositoriesStatusOptions.ERROR}
    <span class="badge variant-filled bg-red-500 w-full rounded-t-md">
      <Ban />{@html "&nbsp;"}
      {repository.status}
    </span>
  {:else if repository.status === RepositoriesStatusOptions.SYNCING}
    <span class="badge variant-filled bg-yellow-500 w-full rounded-t-md">
      <RefreshCw class="animate-spin" />{@html "&nbsp;"}
      {repository.status}
    </span>
  {:else if repository.status === RepositoriesStatusOptions.PENDING}
    <span class="badge variant-filled bg-gray-500 w-full rounded-l-md">
      <RefreshCw />{@html "&nbsp;"}
      {repository.status}
    </span>
  {/if}
</div> -->
