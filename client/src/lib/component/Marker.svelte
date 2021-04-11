<script>
  import { getContext, setContext, onMount } from 'svelte';
	import { browser } from '$app/env'
  
	let classNames = undefined;
  export { classNames as class };

	export let marker = undefined;

	export let width = 30;
	export let height = 30;
	export let lat = 0
	export let lng = 0
	const layerGroup = getContext('layerGroup')();
	setContext('layer', () => marker);

	let L;
	onMount(async () => {
		if (browser) {
			L = await import('leaflet')
		}
	})

	$: if (marker) {
		marker.setLatLng([lat, lng]).update()
	}

  function createMarker(markerElement) {
		let icon = L.divIcon({ 
			html: markerElement, 
			className: 'map-marker',
			iconSize: L.point(width, height)
		});
		marker = L.marker([lat, lng], { icon }).addTo(layerGroup);

    return {
      destroy() {
        if (marker) {
          marker.remove();
          marker = undefined;
        }
      },
    };
  }
</script>

<div class="hidden">
	{#if L && lat && lng}
		<div use:createMarker class={classNames}>
			{#if marker}
				<slot />
			{/if}
		</div>
	{/if}
</div>
