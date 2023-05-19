<script lang="ts">
	import Icon from "svelte-icons-pack";
	import HiOutlineArrowSmLeft from "svelte-icons-pack/hi/HiOutlineArrowSmLeft";

	function onClick() {
		window.history.back();
	}
</script>

<button on:click={onClick}>
	<Icon className="inline w-5 h-5" src={HiOutlineArrowSmLeft} />
	Back
</button>
