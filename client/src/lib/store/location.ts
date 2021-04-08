import { readable } from 'svelte/store'

export const location = readable<{x: number; y: number;}>({x: 0, y: 0}, function start (set) {
  const id = navigator.geolocation.watchPosition(
    function success(pos) {
      const crd = {
        x: pos.coords.longitude,
        y: pos.coords.latitude
      }
      set(crd)
    }
  )

  return function stop () {
    navigator.geolocation.clearWatch(id)
  }
})
