<script lang="ts">
  import "../theme.postcss";
  import "@skeletonlabs/skeleton/styles/all.css";
  import "../app.postcss";
  import { AppShell } from "@skeletonlabs/skeleton";
  import Nav from "$lib/components/base/Nav.svelte";
  import { page } from "$app/stores";
  import { FolderGit2, LayoutDashboard, FileCode, Layers, Package } from "lucide-svelte";
  import type { NavRoute } from "$lib/types";
  import { onMount } from "svelte";
  import { fly } from "svelte/transition";
  import { metadata } from "$lib/app/store";
  import { site } from "$lib/config";
  import { beforeNavigate } from "$app/navigation";

  $: title = $metadata.title ? $metadata.title + " | " + site.name : site.name;
  $: description = $metadata.description ?? site.description;
  // reset metadata on navigation so that the new page inherits nothing from the old page
  beforeNavigate(() => {
    $metadata = {};
  });

  let ready = false;
  onMount(() => (ready = true));

  let routes: NavRoute[] = [
    {
      id: "1",
      name: "Dashboard",
      href: "/",
      icon: LayoutDashboard
    },
    {
      id: "2",
      name: "Deployments",
      href: "/deployments",
      icon: Package
    },
    {
      id: "3",
      name: "Stages",
      href: "/stages",
      icon: Layers
    },
    {
      id: "4",
      name: "Applications",
      href: "/applications",
      icon: FileCode
    },
    {
      id: "5",
      name: "Repositories",
      href: "/repositories",
      icon: FolderGit2
    }
  ];
</script>

<svelte:head>
  <title>{title}</title>
  <meta name="description" content={description} />
</svelte:head>

<div class="h-full">
  <div class=" h-32 bg-primary-500 absolute -z-10 w-full">
    <div class="absolute pr-4 pt-4 pb-4 left-64 right-0 top-12">
      {#each routes as route}
        {#if ($page.route.id && $page.route.id
            .split("/")[1]
            ?.includes(route.name.toLowerCase())) || ($page.route.id === "/" && route.name === "Dashboard")}
          <h1 class="h1 text-white font-bold mb-6">
            <svelte:component this={route.icon} class="w-14 h-14 inline -mt-2" />
            {route.name}
          </h1>
        {/if}
      {/each}
    </div>
  </div>
  <div class="bottom-0 top-32 bg-gray-100 absolute -z-10 w-full" />
  <AppShell>
    <svelte:fragment slot="sidebarLeft">
      <Nav {routes} />
    </svelte:fragment>
  </AppShell>
  <div
    class="absolute pr-4 pt-4 pb-4 left-64 right-0 top-32 bottom-0 overflow-y-scroll hide-scrollbar"
  >
    {#if ready}
      <div transition:fly={{ y: 100, duration: 200 }}>
        <slot />
      </div>
    {/if}
  </div>
</div>
