<script>
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
	import HiOutlineServer from 'svelte-icons-pack/hi/HiOutlineServer';
	import HiSolidServer from 'svelte-icons-pack/hi/HiSolidServer';
	import HiOutlineTemplate from 'svelte-icons-pack/hi/HiOutlineTemplate';
	import HiSolidTemplate from 'svelte-icons-pack/hi/HiSolidTemplate';

	let routes = [
		{ name: 'Dashboard', href: '/', icon: HiOutlineChartPie, currentIcon: HiSolidChartPie },
		{
			name: 'Infrastructure',
			href: '/infrastructure',
			icon: HiOutlineServer,
			currentIcon: HiSolidServer
		},
		{
			name: 'Applications',
			href: '/applications',
			icon: HiOutlineTemplate,
			currentIcon: HiSolidTemplate
		},
		{ name: 'Repositories', href: '/repositories', icon: HiOutlineInbox, currentIcon: HiSolidInbox }
	];
</script>

<div class="h-full">
	<div class="h-24 bg-primary-500 absolute -z-10 w-full" />
	<div class="bottom-0 top-24 bg-gray-100 absolute -z-10 w-full" />
	<AppShell>
		<svelte:fragment slot="sidebarLeft">
			<Nav {routes} />
		</svelte:fragment>
	</AppShell>
	<div class="absolute p-4 left-64 right-0 top-24 bottom-0 overflow-y-scroll hide-scrollbar">
		{#each routes as route}
			{#if $page.route.id === route.href}
				<h1 class="text-2xl font-bold">{route.name}</h1>
			{/if}
		{/each}
		<slot />
	</div>
</div>
