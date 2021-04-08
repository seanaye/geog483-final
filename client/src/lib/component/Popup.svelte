<script>
  import { getContext, onMount } from 'svelte';
  import { browser } from '$app/env'
  let classNames = undefined;
  export { classNames as class };
  export let popup = undefined;
  let showContents = false;
  let popupOpen = false;
  const layer = getContext('layer')();

  let L;
	onMount(async () => {
		if (browser) {
			L = await import('leaflet')
		}
	})
  
	function createPopup(popupElement) {
    popup = L.popup().setContent(popupElement);
    layer.bindPopup(popup);
    layer.on('popupopen', () => {
      popupOpen = true;
      showContents = true;
    });
    layer.on('popupclose', () => {
      popupOpen = false;
      // Wait for the popup to completely fade out before destroying it.
      // Otherwise the fade out looks weird as the contents disappear too early.
      setTimeout(() => {
        if (!popupOpen) {
          showContents = false;
        }
      }, 500);
    });
    return {
      destroy() {
        if (popup) {
          layer.unbindPopup();
          popup.remove();
          popup = undefined;
        }
      },
    };
  }
</script>

<div class="hidden">
  {#if L}
    <div use:createPopup class={classNames}>
      {#if showContents}
        <slot />
      {/if}
    </div>
  {/if}
</div>
