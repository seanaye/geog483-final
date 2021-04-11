<script lang="ts">
  import type { User } from '$lib/types/api'
  import Marker from './Marker.svelte'
  import Circle from './Circle.svelte'
  import Popup from './Popup.svelte'

  export let user: User
</script>
<Marker lat={user.coords.y} lng={user.coords.x} width={30} height={30}>
  <svg style="width:30px;height:30px" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
  </svg>
  
  <Popup class="w-1/2">
    <div class="w-full bg-white">
      name: {user.name}, lat: {user.coords.y}, lng: {user.coords.x}
    </div>
  </Popup>
</Marker> 
<Circle lat={user.coords.y} lng={user.coords.x} radius={user.radius} />

