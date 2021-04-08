<script lang="ts">
	import '../app.postcss';
	import CreateSession from '$lib/component/CreateSession.svelte'
	import { initClient } from '@urql/svelte'
	import { sessionStore } from '$lib/util/urql'

	let token = ''

	initClient({
		url: 'http://localhost:8080/query',
		fetchOptions: () => {
			return {
				headers: { authorization: token ? `Bearer ${token}` : '' }
			}
		}
	})
	$: token = $sessionStore.data?.createSession?.token || ''

</script>

<div class="bg-gray-200 absolute top-0 bottom-0 left-0 right-0 grid place-items-center">
{#if token}
	<slot />
{:else}
	<CreateSession />
{/if}
</div>

