<script type="ts">
	// @ts-nocheck
	import Icon from 'svelte-icons-pack/Icon.svelte';
	import { fade } from 'svelte/transition';
	import { page } from '$app/stores';
	import { clickOutside } from '$lib/utils/clickOutside';
	export let routes = [];
	import HiOutlineLogout from 'svelte-icons-pack/hi/HiOutlineLogout';

	let open = false;

	function toggle() {
		open = !open;
	}
</script>

<div class="h-full p-4 bg-transparent">
	<div class="bg-white shadow-md h-full rounded-lg w-56 z-20">
		<div class="flex flex-shrink-0 items-center pt-4 px-8 ">
			<img class="h-12 w-auto" src="./images/momentum-logo.png" alt="Momentum Logo" />
			<span class="ml-2 text-xl font-semibold text-primary-500">Momentum</span>
		</div>

		<!-- Navigation -->
		<nav class="mt-4 px-3 pt-6">
			<div class="space-y-4">
				{#each routes as route}
					<a
						href={route.href}
						class="text-gray-500 unstyled border-2 rounded-md group flex items-center hover:text-gray-600 pl-9 py-2 text-sm font-medium {$page
							.route.id === route.href
							? 'text-gray-600 border-gray-600'
							: ''}"
						aria-current="page"
					>
						{#if $page.route.id === route.href}
							<Icon className="w-5 h-5" src={route.currentIcon} />{@html '&nbsp;'}
							<div class="block">{route.name}</div>
						{:else}
							<Icon className="w-5 h-5" src={route.icon} />{@html '&nbsp;'}{route.name}
						{/if}
					</a>
				{/each}
			</div>
		</nav>
		<div class="bottom-6 absolute w-full left-20 ">
			<a
				href="/logout"
				class="text-gray-500 unstyled group flex items-center hover:text-gray-600 px-2 py-2 text-sm font-medium"
				aria-current="page"
			>
				<Icon className="w-5 h-5" src={HiOutlineLogout} />{@html '&nbsp;'}Sign-out
			</a>
		</div>
	</div>
</div>

<style>
	.profile-icon {
		color: #fff;
	}
</style>
