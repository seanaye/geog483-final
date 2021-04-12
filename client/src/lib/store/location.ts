import { readable } from 'svelte/store'

// this will hold the watcher for the device location
let watcherId: number;

export const location = readable<{x: number; y: number;}>({x: 0, y: 0}, function start (set) {
  // define a function to update the store when given a
  // GeolocationPosition
  // Read more at https://developer.mozilla.org/en-US/docs/Web/API/GeolocationPosition
  const posCallback = (pos: GeolocationPosition) => {
    const crd = {
      x: pos.coords.longitude,
      y: pos.coords.latitude
    }
    console.log( crd )
    set(crd)
  }

  // get the current position now
  navigator.geolocation.getCurrentPosition(posCallback)

  // if we are not already watching the position change
  // create a watcher which will update when the position changes
  if (!watcherId) {
    watcherId = navigator.geolocation.watchPosition(posCallback)
  }

  // cleanup function
  return function stop () {
    // clear the watcher
    navigator.geolocation.clearWatch(watcherId)
  }
})
