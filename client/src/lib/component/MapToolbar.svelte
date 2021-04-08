<script>
	import Slider from '$lib/component/Slider.svelte'
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();
	export let eye = true;
	function clickEye() {
		eye = !eye;
		dispatch('click-eye', eye);
	}
	
	export let lines = true;
	function clickLines() {
		lines = !lines;
		dispatch('click-lines', lines);
	}
</script>

<div class="bg-white shadow-lg px-4 rounded-lg">
	<Slider />
</div>
