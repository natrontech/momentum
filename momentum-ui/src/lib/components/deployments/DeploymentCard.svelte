<script lang="ts">
  import type { ModelsDeployment } from "$lib/momentum-core-client/model/models";

  export let deployment: ModelsDeployment;

</script>

<div>
  <div class="flex relative">
    <a
      class="flex flex-col justify-center items-center w-full relative"
      href={"/deployments/" + deployment.id}
      >
      <h3 class="text-lg leading-6 font-medium text-gray-900">
        {deployment.name}
      </h3>
    </a>
  </div>
</div>
