<script>
  /** @type {import('./$types').PageData} */
  export let data;
	import BackButton from '$lib/components/base/BackButton.svelte';
  import { CodeBlock } from '@skeletonlabs/skeleton';

  let mockupYAML = `
image: nginx:1.19.6
replicas: 1
resources:
  cpu: 100m
  memory: 128Mi
env:
  - name: TEST
    value: test
ingress:
  enabled: true
  hostName: test.momentum.io
  path: /
  tls:
    enabled: true
    secretName: test-tls
  `

</script>

<BackButton />

<h2 class="h2 text-primary-500 font-bold mb-6">
  {data.props.deployment.name}
</h2>

<CodeBlock language="yaml" code={mockupYAML}></CodeBlock>
