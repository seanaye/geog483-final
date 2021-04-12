import { writable } from 'svelte/store'
import type { MsgFragment } from '$lib/types/api'


function createMessage () {
  const { subscribe, update } = writable<MsgFragment[]>([])

  const append = (msg: MsgFragment) => {
    update((data) => {
      return [...data, msg]
    })
  }

  return {
    subscribe,
    append
  }
}

export const messages = createMessage()

