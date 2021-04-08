<script lang="ts">
  export let label: string
  export let value: string
  export let hint: string|undefined
</script>
<div>
  <label for="input" class="block text-sm font-medium text-gray-700">{label}</label>
  <div class="mt-1">
    <input
        type="text"
        name="input"
        bind:value={value}
        class="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
    >
  </div>
    <p class="mt-2 text-sm text-gray-500" id="email-description">{hint || ''}</p>
</div>

