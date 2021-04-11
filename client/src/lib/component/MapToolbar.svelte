<script lang="ts">
	import{ onDestroy } from 'svelte'
	import Slider from '$lib/component/Slider.svelte'
	import { radiusStore } from '$lib/util/urql'
	import { mutation } from '@urql/svelte'

	const mut = mutation(radiusStore)

	let value = 500

	let timer: NodeJS.Timeout;
	function debounceMut () {
		clearTimeout(timer)
		timer = setTimeout(() => {
			mut({rad: value})
		}, 300)
	}
	$: if (value) debounceMut()

	onDestroy(() => clearTimeout(timer))
	
</script>

<div class="bg-white shadow-lg p-4 rounded-lg">
	<Slider bind:value={value} />
</div>
