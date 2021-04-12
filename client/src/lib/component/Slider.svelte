<script lang="ts">
  // minimum slider value in meters
  export let min = 1000
  // maximum slider value in meters
  export let max = 10000

  export let value:number

  function formatMeters (m: number) {
    if (m < 1000) return `${m} m`
    const km = Math.floor(m / 1000)
    return `${km} km`
  }
</script>
<div class="flex flex-row">
  <div>{formatMeters(min)}</div>
  <input type="range" {min} {max} bind:value={value} />
  <div>{formatMeters(max)}</div>
</div>
<div class="w-full text-center">{formatMeters(value)}</div>

