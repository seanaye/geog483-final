<script lang="ts">
  import { readable } from 'svelte/store'
  import { messages } from '$lib/store/message'

  function timeSince (time: number, start: number) {
    const seconds = Math.floor((time - start))

    let interval = Math.floor(seconds / 31536000) // seconds in a year
    if (interval > 1) return `${interval}y`

    interval = Math.floor(seconds / 2592000) // seconds in a month
    if (interval > 1) return `${interval}m`

    interval = Math.floor(seconds / 86400)
    if (interval > 1) return `${interval}d`

    interval = Math.floor(seconds / 3600)
    if (interval > 1) return `${interval}h`

    interval = Math.floor(seconds / 60)
    if (interval > 1) return `${interval}min`

    return (seconds > 0) ? `${seconds}s` : 'Just now'
  }

  function getEpoch () {
    const now = new Date().getTime()
    return Math.floor(now / 1000)
  }

  const timeStore = readable(getEpoch(), function start(set) {
    const id = setInterval(() => {
      set(getEpoch())
    }, 60000)

    return function stop () {
      clearTimeout(id)
    }
  })

</script>

<!-- This example requires Tailwind CSS v2.0+ -->
<ul class="divide-y divide-gray-200">
  {#each $messages as message}
    <li class="py-4">
      <div class="flex space-x-3">
        <div class="flex-1 space-y-1">
          <div class="flex items-center justify-between">
            <h3 class="text-sm font-medium">{message.user.name}</h3>
            <p class="text-sm text-gray-500">{timeSince(message.time, $timeStore)}</p>
          </div>
          <p class="text-sm text-gray-500">{message.content}</p>
        </div>
      </div>
    </li>
  {/each}
</ul>
