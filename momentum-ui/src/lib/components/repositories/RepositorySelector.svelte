<script lang="ts">
    import type { RepositoriesResponse } from "$lib/pocketbase/generated-types";
    import { ActiveRepository } from "$lib/stores/activeRepository";

    export let repos: RepositoriesResponse[] = [];

    $: repo = $ActiveRepository
  
    let showDropdown = false;

    function setActiveRepository(repo: RepositoriesResponse) {
      ActiveRepository.set(repo);
      toggleDropdown();
    }

    function toggleDropdown() {
        showDropdown = !showDropdown;
    }
  </script>

  <div>
    <div class="dropdown-container">
        <div class="selected-option" on:click={toggleDropdown}>
          {#if repo}
            {repo.name}
          {:else}
            choose repository
          {/if}
        </div>
    
        {#if showDropdown}
          <ul class="dropdown">
            {#each repos as repo}
              <li on:click={() => setActiveRepository(repo)} on:keydown={() => setActiveRepository(repo)}>{repo.name}</li>
            {/each}
          </ul>
        {/if}
      </div>
</div>

<style>
    .dropdown-container {
      position: relative;
    }
  
    .selected-option {
      cursor: pointer;
      padding: 0.5rem;
      border: 1px solid #ccc;
      border-radius: 4px;
    }
  
    .dropdown {
      position: absolute;
      top: 100%;
      left: 0;
      list-style: none;
      padding: 0;
      margin: 0;
      border: 1px solid #ccc;
      border-radius: 4px;
      background-color: #fff;
    }
  
    .dropdown li {
      cursor: pointer;
      padding: 0.5rem;
    }
  
    .dropdown li:hover {
      background-color: #f0f0f0;
    }
  </style>