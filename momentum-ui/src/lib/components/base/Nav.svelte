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
	<div class="bg-white h-full rounded-lg w-56 z-20">
		<div class="flex flex-shrink-0 items-center pt-4 px-8 ">
			<img class="h-12 w-auto" src="./images/natrium-logo.png" alt="Momentum Logo" />
			<span class="ml-2 text-xl font-semibold text-primary-500">Momentum</span>
		</div>
		<div class="my-6 grid place-items-center">
			<button
				on:click|stopPropagation={() => toggle()}
				use:clickOutside
				on:click_outside={() => (open = false)}
				class="relative"
			>
				<img
					class="h-16 w-16 flex-shrink-0 rounded-full bg-gray-300 mb-4"
					src="https://images.unsplash.com/photo-1502685104226-ee32379fefbe?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=3&w=256&h=256&q=80"
					alt=""
				/>
				{#if open}
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<div
						on:click|stopPropagation={() => toggle()}
						transition:fade={{ duration: 100 }}
						class="absolute top-0 z-20 divide-y bg-secondary-700 w-16 h-16 rounded-full flex items-center justify-center"
						role="menu"
						aria-orientation="vertical"
						aria-labelledby="options-menu-button"
						tabindex="-1"
					>
						<div class="py-1" role="none">
							<a
								href="/profile"
								class="text-white block px-4 py-2 text-sm unstyled"
								role="menuitem"
								tabindex="-1"
								id="options-menu-item-0"
							>
							</a>
						</div>
					</div>
				{/if}
			</button>

			<div class="mx-auto w-auto">
				<span class="text-gray-500 text-sm block"> Welcome back, </span>
				<span class="text-gray-700 text-sm font-medium block"> Jessy Schwarz </span>
			</div>
		</div>

		<!-- Navigation -->
		<nav class="mt-6 px-3">
			<div class="space-y-4">
				{#each routes as route}
					<a
						href={route.href}
						class="text-primary-400 unstyled group flex items-center hover:text-primary-900 pl-9 py-2 text-sm font-medium {$page
							.route.id === route.href
							? 'text-primary-900'
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
				class="text-primary-400 unstyled group flex items-center hover:text-primary-900 px-2 py-2 text-sm font-medium"
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
