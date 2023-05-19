<script lang='ts'>
	import '../theme.postcss';
	import '@skeletonlabs/skeleton/styles/all.css';
	import '../app.postcss';
	import { AppShell } from '@skeletonlabs/skeleton';
	import Nav from '$lib/components/base/Nav.svelte';
	import { page } from '$app/stores';
	import HiOutlineChartPie from 'svelte-icons-pack/hi/HiOutlineChartPie';
	import HiSolidChartPie from 'svelte-icons-pack/hi/HiSolidChartPie';
	import HiOutlineInbox from 'svelte-icons-pack/hi/HiOutlineInbox';
	import HiSolidInbox from 'svelte-icons-pack/hi/HiSolidInbox';
	import HiOutlineTemplate from 'svelte-icons-pack/hi/HiOutlineTemplate';
	import HiSolidTemplate from 'svelte-icons-pack/hi/HiSolidTemplate';
	import HiOutlineCube from 'svelte-icons-pack/hi/HiOutlineCube';
	import HiSolidCube from 'svelte-icons-pack/hi/HiSolidCube';
	import HiOutlineCollection from 'svelte-icons-pack/hi/HiOutlineCollection';
	import HiSolidCollection from 'svelte-icons-pack/hi/HiSolidCollection';
	import type { NavRoute } from '$lib/types';


	let routes: NavRoute[] = [
		{ id: '1',name: 'Dashboard', href: '/', icon: HiOutlineChartPie, currentIcon: HiSolidChartPie },
		{
			id: '2',
			name: 'Deployments',
			href: '/deployments',
			icon: HiOutlineCube,
			currentIcon: HiSolidCube
		},
		{
			id: '3',
			name: 'Stages',
			href: '/stages',
			icon: HiOutlineCollection,
			currentIcon: HiSolidCollection
		},
		{
			id: '4',
			name: 'Applications',
			href: '/applications',
			icon: HiOutlineTemplate,
			currentIcon: HiSolidTemplate
		},
		{ id: '5', name: 'Repositories', href: '/repositories', icon: HiOutlineInbox, currentIcon: HiSolidInbox }
	];
</script>

<div class="h-full">
	<div class=" h-32 bg-primary-500 absolute -z-10 w-full" />
	<div class="bottom-0 top-32 bg-gray-100 absolute -z-10 w-full" />
	<AppShell>
		<svelte:fragment slot="sidebarLeft">
			<Nav {routes} />
		</svelte:fragment>
	</AppShell>
	<div
		class="absolute pr-4 pt-4 pb-4 left-64 right-0 top-32 bottom-0 overflow-y-scroll hide-scrollbar"
	>
		{#each routes as route}
			{#if $page.route.id === route.href}
				<h1 class="h1 text-primary-500 font-bold mb-6">
					{route.name}
				</h1>
			{/if}
		{/each}
		<slot />
	</div>
</div>
