<script lang="ts">
	import Input from '$lib/component/Input.svelte'
	import Leaflet from '$lib/component/Leaflet.svelte'
	import Control from '$lib/component/Control.svelte'
	import MapToolbar from '$lib/component/MapToolbar.svelte'
	import UserMarker from '$lib/component/UserMarker.svelte'
	import { onDestroy, onMount } from 'svelte'
	import { query, mutation, subscription } from '@urql/svelte'
	import { users } from '$lib/store/users'
	import { location } from '$lib/store/location'
	import {
		updateCoordsStore,
		delUsersStore,
		usersUpdatesStore,
		endSessionStore,
		getUsersStore,
		messagesStore
} from '$lib/util/urql'
 
	let map;
	
	const initialView = [43.466667, -80.516670];
	
	let eye = true;

	// fetch and add initial users
	const usersStore = query(getUsersStore)
	$: {
		const usersPayload = $usersStore.data?.users
		if (usersPayload && !$users.length) {
			users.set(usersPayload)
		}
	}

	// subscribe to user updates
	const usersUpdate = subscription(usersUpdatesStore)

	// when a user update comes from the server,
	// put it in the users store
	$: {
		const usersPayload = $usersUpdate.data?.users
		console.log({ usersPayload })
		if (usersPayload) {
			users.addSingle(usersPayload)
		}
	}

	// subscribe to users which are deleted
	const usersDel = subscription(delUsersStore)

	// when the server tells us a user has left
	// remove it from the users store
	$: {
		const deletedId = $usersDel.data?.delUsers
		users.remove(deletedId)
	}

	// send update when our watched location changes
	const updateLocation = mutation(updateCoordsStore)
	$: {
		updateLocation($location);
		console.log('location mutation');
	}

	// subscribe to incoming messages
	const messages = subscription(messagesStore)
	$: {
		const subData = $messages.data
		console.log({ subData })
	}

	function resizeMap() {
	  if(map) { map.invalidateSize(); }
  }
	
	function resetMapView() {
		map.setView(initialView, 5);
	}



	// cleanup the user session
	const endSession = mutation(endSessionStore)
	onDestroy(endSession)
	onMount(() => {
		window.onbeforeunload = () => endSession()
	})

</script>
<svelte:window on:resize={resizeMap} />

<!-- Can just use an import statement for this, when outside the REPL -->
<link rel="stylesheet" href="https://unpkg.com/leaflet@1.6.0/dist/leaflet.css"
   integrity="sha512-xwE/Az9zrjBIphAcBb3F6JVqxf46+CDLwfLMHloNu6KEQCAWi6HcDUbeOfBIptF7tcCzusKFjFw2yuvEpDL9wQ=="
   crossorigin=""/>

<Leaflet bind:map view={initialView} zoom={12}>
	<Control position="topright">
		<MapToolbar  />
	</Control>
	
	{#if eye}
		{#each $users as user (user.id) }
			<UserMarker {user} />
		{/each}
	{/if}
	<Control position="bottomleft">
		<div class="w-screen -ml-2 -mb-3">
			<div class="p-4 bg-white flex flex-row items-end w-11-12 mx-auto">
				<Input label="Message" />
				<button type="button" class="h-11 mb-2 inline-flex items-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
					Send
				</button>	
			</div>
		</div>
	</Control>
</Leaflet>
