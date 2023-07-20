<script lang="ts">
  
  import type { ModelsApplication } from "$lib/momentum-core-client/api";

  export let application: ModelsApplication;

</script>

<div>
  <div class="flex relative">
    <a
      class="flex flex-col justify-center items-center w-full relative"
      href={"/applications/" + application.id}
      >
      <h3 class="text-lg leading-6 font-medium text-gray-900">
        {application.name}
      </h3>
    </a>
  </div>
</div>
