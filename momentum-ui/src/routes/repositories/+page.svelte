<script lang="ts">
  import { metadata } from "$lib/app/store";
  import { watch } from "$lib/pocketbase";
  import type { RepositoriesResponse } from "$lib/pocketbase/generated-types";

  $metadata.title = "Repositories";

  const repositories = watch<RepositoriesResponse>("repositories", {
    sort: "-updated"
  });

  console.log(repositories);
</script>
