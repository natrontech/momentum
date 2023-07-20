<script lang="ts">
    import type { RepositoriesResponse } from "$lib/pocketbase/generated-types";
    import { ActiveRepository } from "$lib/stores/activeRepository";
    import { load } from "../../../routes/repositories/+page";

    $: repo = ActiveRepository
  
  </script>
  
  <div>
    
  </div>
  