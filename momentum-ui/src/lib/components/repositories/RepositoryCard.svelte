<script lang="ts">
  import {
    RepositoriesStatusOptions,
    type RepositoriesResponse
  } from "$lib/pocketbase/generated-types";
  import { RefreshCw, Ban, Check, Lock } from "lucide-svelte";

  export let repository: RepositoriesResponse;

  function onHandleStatusClick(event: any): void {
    console.log(event);
  }
</script>

<div
  class="card rounded-md border-2 transition-all duration-300 ease-in-out
  {repository.status === RepositoriesStatusOptions.UP
    ? 'border-green-400 hover:border-green-500'
    : repository.status === RepositoriesStatusOptions.ERROR
    ? 'border-red-400 hover:border-red-500'
    : repository.status === RepositoriesStatusOptions.SYNC
    ? 'border-yellow-400 hover:border-yellow-500'
    : 'border-gray-400 hover:border-gray-500'}
"
>
  <div class="flex relative">
    <button
      on:click={onHandleStatusClick}
      class="p-6 rounded-l-md cursor-pointer w-52 group hover:w-full transition-all duration-300 ease-in-out
      {repository.status === RepositoriesStatusOptions.UP
        ? 'bg-green-400 hover:bg-green-500'
        : repository.status === RepositoriesStatusOptions.ERROR
        ? 'bg-red-400 hover:bg-red-500'
        : repository.status === RepositoriesStatusOptions.SYNC
        ? 'bg-yellow-400 hover:bg-yellow-500'
        : 'bg-gray-400 hover:bg-gray-500'}
    "
    >
      <div class="flex flex-col justify-center items-center ">
        {#if repository.status === RepositoriesStatusOptions.SYNC}
          <RefreshCw
            class="text-white group-hover:rotate-180 transition-all duration-300 ease-in-out"
          />
        {:else if repository.status === RepositoriesStatusOptions.ERROR}
          <Ban class="text-white" />
        {:else if repository.status === RepositoriesStatusOptions.UP}
          <Check class="text-white" />
        {:else if repository.status === RepositoriesStatusOptions.SYNC}
          <RefreshCw class="animate-spin text-white" />
        {/if}
        <p class=" text-xs text-white">
          {repository.status}
        </p>
      </div>
    </button>
    <a
      class="flex flex-col justify-center items-center w-full relative"
      href={"/repositories/" + repository.id}
    >
      <h3 class="text-lg leading-6 font-medium text-gray-900">
        {repository.name}
      </h3>
    </a>
    <button class="absolute right-2 bottom-2">
      <Lock class="text-gray-500 hover:text-gray-900 transition-all duration-150 ease-in-out" />
    </button>
  </div>
</div>
