<script>
  import { getContext, onMount } from 'svelte';
  import { browser } from '$app/env'
  let classNames = undefined;
  export { classNames as class };

  let createControl;

  onMount(async () => {
    if (browser) {
      const L = await import('leaflet')

      class Control extends L.Control {
        constructor(
          el,
          position
        ) {
          super({ position });
          this.el = el;
        }
        onAdd() {
          L.DomEvent.disableClickPropagation(this.el)
          return this.el;
        }
        onRemove() {}
      }

      createControl = (container) => {
        control = new Control(container, position).addTo(map);
        return {
          destroy() {
            control.remove();
            control = undefined;
          },
        };
      }
    }
  })
  
	/** position: 'topleft' | 'topright' | 'bottomleft' | 'bottomright' */
	export let position;

	/** The control instance created by this component */
  export let control = undefined;
  const map = getContext('map')();
	
</script>

<div style="display:hidden">
  {#if createControl}
    <div use:createControl class={classNames}>
      {#if control}
        <slot {control} />
      {/if}
    </div>
  {/if}
</div>
