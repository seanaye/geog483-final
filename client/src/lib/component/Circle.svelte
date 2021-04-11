<script lang="ts">
  import { getContext, setContext, onMount } from 'svelte';
	import { browser } from '$app/env'

  let classNames = undefined;
  export { classNames as class }

  export let circle = undefined;
  export let lat = 0
  export let lng = 0
  export let radius = 0

  let L;
  onMount(async () => {
    if (browser) {
      L = await import('leaflet')
    }
  })

  $: if (circle) {
    circle.setLatLng([lat, lng])
    circle.setRadius(radius)
  }

  const layerGroup = getContext<()=>void>('layerGroup')();
  setContext('layer', () => circle)


  function createCircle(circleElement) {
    circle = L.circle([lat, lng], {
      color: "red",
      fillColor: '#f03',
      fillOpacity: 0.5,
      radius
    }).addTo(layerGroup)
    console.log({ circle })
    return {
      destroy () {
        if (circle) {
          circle.remove()
          circle = undefined
        }
      }
    }
  }
</script>

<div class="hidden">
  {#if L && lat && lng}
    <div use:createCircle class={classNames}>
      {#if circle}
        <slot />
        <div>hi</div>
      {/if}
    </div>
  {/if}
</div>
