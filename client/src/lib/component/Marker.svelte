<script>
  import { getContext, setContext, onMount } from 'svelte';
	import { browser } from '$app/env'
  
	let classNames = undefined;
  export { classNames as class };

	export let marker = undefined;

	export let width = 30;
	export let height = 30;
	export let latLng;
 
	const layerGroup = getContext('layerGroup')();
	setContext('layer', () => marker);

	let L;
	onMount(async () => {
		if (browser) {
			L = await import('leaflet')
		}
	})

  function createMarker(markerElement) {
		let icon = L.divIcon({ 
			html: markerElement, 
			className: 'map-marker',
			iconSize: L.point(width, height)
		});
		marker = L.marker(latLng, { icon }).addTo(layerGroup);

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
	{#if L}
		<div use:createMarker class={classNames}>
			{#if marker}
				<slot />
			{/if}
		</div>
	{/if}
</div>
