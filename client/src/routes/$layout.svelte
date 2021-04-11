<script lang="ts">
	import '../app.postcss';
	import CreateSession from '$lib/component/CreateSession.svelte'
	import { client } from '$lib/util/urql-client'
	import { sessionStore } from '$lib/util/urql'
	import { setClient } from '@urql/svelte'

	let token = ''

	setClient(client)
	$: token = $sessionStore.data?.createSession?.token || ''

</script>

{#if token}
<div class="absolute top-0 bottom-0 left-0 right-0">
	<slot />
</div>
{:else}
<div class="bg-gray-200 absolute top-0 bottom-0 left-0 right-0 grid content-center place-items-center grid-rows-none grid-cols-none">
	<CreateSession />
</div>
{/if}

