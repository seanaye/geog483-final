<script lang="ts">
	import { mutation } from '@urql/svelte'
	import { sessionStore } from '$lib/util/urql'
	import Input from '$lib/component/Input.svelte'
	import { location } from '$lib/store/location'
	let name = ''

	const mut = mutation(sessionStore)

	let prom: Promise<void>|null = null
	async function createSession() {
		await mut({ name, x: $location.x, y: $location.y })
		return
	}

	function handleClick () {
		prom = createSession()
	}
</script>

<div class="container px-4">
	<div class="bg-white rounded-lg shadow-lg w-full md:w-3/4 lg:w-1/2 my-4 flex flex-col p-4 mx-auto">
		<Input label="What's your name?" bind:value={name} />
		<div class="flex- flex-row">
			<button on:click={handleClick} type="button" class="inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
				Submit
			</button>	
			{#await prom}
				<svg class="animate-spin h-5 w-5 mr-3 ..." viewBox="0 0 24 24"></svg>
			{:catch}
				<div>There was an error creating your user session</div>
			{/await}
		</div>
	</div>
</div>
