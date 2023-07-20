<script lang="ts">
  // @ts-nocheck
  import { page } from "$app/stores";
  import type { NavRoute } from "$lib/types";

  export let routes: NavRoute[] = [];
</script>

<div class="h-full p-4 bg-transparent">
  <div class="bg-white shadow-md h-full rounded-lg w-56 z-20 relative">
    <div class="flex flex-shrink-0 items-center pt-4 px-8 flex-col">
      <img class="h-12 w-auto" src="/images/momentum-logo.png" alt="Momentum Logo" />
      <p class="mt-2 text-xl font-bold text-primary-500 uppercase">Momentum</p>
    </div>

    <!-- Navigation -->
    <nav class="mt-3 px-3 pt-6">
      <div class="space-y-4">
        {#each routes as route}
          <a
            href={route.href}
            class="text-gray-500 unstyled border-2 rounded-md group flex items-center hover:text-gray-600 pl-9 py-2 text-sm font-medium
						{$page.route.id.split('/')[1]?.includes(route.name.toLowerCase()) ||
            ($page.route.id === '/' && route.name === 'Dashboard')
              ? 'text-gray-700 border-gray-700'
              : ''}
"
            aria-current="page"
          >
            {#if $page.route.id
              .split("/")[1]
              ?.includes(route.name.toLowerCase()) || ($page.route.id === "/" && route.name === "Dashboard")}
              <!-- <Icon className="w-5 h-5" src={route.currentIcon} />{@html '&nbsp;'} -->
              <svelte:component this={route.icon} class="w-5 h-5" />{@html "&nbsp;"}
              <div class="block">{route.name}</div>
            {:else}
              <svelte:component this={route.icon} class="w-5 h-5" />{@html "&nbsp;"}{route.name}
            {/if}
          </a>
        {/each}
      </div>
    </nav>
  </div>
</div>
