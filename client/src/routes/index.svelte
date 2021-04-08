<script lang="ts">
	import Leaflet from '$lib/component/Leaflet.svelte'
	import Control from '$lib/component/Control.svelte'
	import MapToolbar from '$lib/component/MapToolbar.svelte'
	import Marker from '$lib/component/Marker.svelte'
	import Popup from '$lib/component/Popup.svelte'

	let map;

	const markerLocations = [
		[29.8283, -96.5795],
		[37.8283, -90.5795],
		[43.8283, -102.5795],
		[48.40, -122.5795],
		[43.60, -79.5795],
		[36.8283, -100.5795],
		[38.40, -122.5795],
	];
	
	
	const initialView = [39.8283, -98.5795];
	
	let eye = true;
	let showLines = true;

	function resizeMap() {
	  if(map) { map.invalidateSize(); }
  }
	
	function resetMapView() {
		map.setView(initialView, 5);
	}

</script>
<svelte:window on:resize={resizeMap} />

<!-- Can just use an import statement for this, when outside the REPL -->
<link rel="stylesheet" href="https://unpkg.com/leaflet@1.6.0/dist/leaflet.css"
   integrity="sha512-xwE/Az9zrjBIphAcBb3F6JVqxf46+CDLwfLMHloNu6KEQCAWi6HcDUbeOfBIptF7tcCzusKFjFw2yuvEpDL9wQ=="
   crossorigin=""/>

<Leaflet bind:map view={initialView} zoom={4}>
	<Control position="topright">
		<MapToolbar bind:eye bind:lines={showLines} on:click-reset={resetMapView} />
	</Control>
	
	{#if eye}
		{#each markerLocations as latLng}
			<Marker {latLng} width={30} height={30}>
				<svg style="width:30px;height:30px" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" viewBox="0 0 24 24" stroke="currentColor"><path d="M8 14v3m4-3v3m4-3v3M3 21h18M3 10h18M3 7l9-4 9 4M4 10h16v11H4V10z"></path></svg>
				
				<Popup>A popup!</Popup>
			</Marker> 
		{/each}
	{/if}
</Leaflet>
