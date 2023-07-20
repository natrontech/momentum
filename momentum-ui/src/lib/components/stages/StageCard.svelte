<script lang="ts">
  import type { ModelsStage } from "$lib/momentum-core-client/model/models";

  export let stage: ModelsStage;

</script>

<div>
  <div class="flex relative">
    <a
      class="flex flex-col justify-center items-center w-full relative"
      href={"/stages/" + stage.id}
      >
      <h3 class="text-lg leading-6 font-medium text-gray-900">
        {stage.name}
      </h3>
    </a>
  </div>
</div>
