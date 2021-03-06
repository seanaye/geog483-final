<div class="absolute top-0 bottom-0 left-0 right-0 grid grid-rows-6 grid-cols-1">
	<div class="row-span-4 col-span-6">
		<Leaflet bind:map view={[43.466667, -80.516670]} zoom={12}>
			<Control position="topright">
				<MapToolbar  />
			</Control>
			
			{#each $users as user (user.id) }
				<UserMarker {user} />
			{/each}
		</Leaflet>
	</div>
	<div class="grid grid-cols-6 grid-rows-4 w-full row-span-2">
		<div class="col-span-6 p-3 row-span-3 overflow-y-auto h-full">
			<UserMessages />
		</div>
		<div class="col-span-4 md:col-span-5 self-end ml-2">
			<Input label="Message" bind:value={msg} />
		</div>
		<button
			type="button"
			on:click={sendMsg}
			class="col-span-2 md:col-span-1 self-end mb-2 mx-2 h-11 text-center px-4 py-2 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
		>
			Send
		</button>	
	</div>
</div>

<script lang="ts">
	import Input from '$lib/component/Input.svelte'
	import Leaflet from '$lib/component/Leaflet.svelte'
	import Control from '$lib/component/Control.svelte'
	import MapToolbar from '$lib/component/MapToolbar.svelte'
	import UserMarker from '$lib/component/UserMarker.svelte'
	import UserMessages from '$lib/component/UserMessages.svelte'
	import { onDestroy, onMount } from 'svelte'
	import { query, mutation, subscription } from '@urql/svelte'
	import { users } from '$lib/store/users'
	import { location } from '$lib/store/location'
	import { messages } from '$lib/store/message'
	import {
		updateCoordsStore,
		delUsersStore,
		usersUpdatesStore,
		endSessionStore,
		getUsersStore,
		messagesStore,
		sendMessageStore
} from '$lib/util/urql'
 
	let map;
	
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
		console.log( usersPayload )
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
	}

	// subscribe to incoming messages
	const messagesSub = subscription(messagesStore)
	$: {
		const subData = $messagesSub.data?.messages
		if (subData) messages.append(subData)
	}


	// hold value of message input
	let msg = ''
	const sendMsgMut = mutation(sendMessageStore)

	async function sendMsg () {
		if (msg) {
			await sendMsgMut({ content: msg })
			msg = ''
		}
	}


	function resizeMap() {
	  if(map) { map.invalidateSize(); }
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

