<script lang="ts">
	import DeploymentCard from '$lib/components/deployments/DeploymentCard.svelte';
	import { mockDeployments } from '$lib/mock-data';
	import { toggleTableView } from '$lib/store';
	import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';
	import { getContext } from 'svelte';
	import Icon from 'svelte-icons-pack';

	import HiOutlineTable from 'svelte-icons-pack/hi/HiOutlineTable';
	import HiOutlineViewGrid from 'svelte-icons-pack/hi/HiOutlineViewGrid';

	let toggleTableViewValue = false;

	toggleTableView.subscribe((value) => {
		toggleTableViewValue = value;
	});

	function setToggleTableView() {
		toggleTableView.update((value) => !value);
	}

</script>

<div class="flex justify-between items-center mb-6">
	<RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
		<RadioItem bind:group={toggleTableViewValue} on:click={setToggleTableView} name="justify" value={false}><Icon src={HiOutlineTable} /></RadioItem>
    <RadioItem bind:group={toggleTableViewValue} on:click={setToggleTableView} name="justify" value={true}><Icon src={HiOutlineViewGrid} /></RadioItem>
	</RadioGroup>
</div>

<div class='grid gap-6 {toggleTableViewValue ? 'grid-cols-2' : ''}'>
	{#each mockDeployments as deployment}
		<DeploymentCard {deployment} />
	{/each}
</div>
