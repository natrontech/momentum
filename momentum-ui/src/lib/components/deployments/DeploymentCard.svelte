<script lang="ts">
	import { DeploymentStatus, type Deployment } from '$lib/types';
	export let deployment: Deployment;
</script>

<a class="card rounded-md" href={'deployments/' + deployment.id}>
	<div class="px-4 py-5 sm:px-6">
		<h3 class="text-lg leading-6 font-medium text-gray-900">
			{deployment.name}
		</h3>
		<span
			class="badge variant-filled {deployment.status == DeploymentStatus.Running
				? 'bg-green-500'
				: 'bg-red-500'}">{deployment.status}</span
		>
	</div>
</a>
