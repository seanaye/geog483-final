import type { User } from '$lib/types/api'
import { writable } from 'svelte/store'


function createUsers () {
  const { subscribe, set, update } = writable<User[]>([])

  const remove = (id: string) => {
    update((data) => {
      return data.filter(user => user.id !== id)
    })
  }

  const addSingle = (u: User) => {
    update((data) => {
      const existIndex = data.findIndex(user => user.id === u.id)
      if (existIndex === -1) {
        return [...data, u]
      } else {
        data.splice(existIndex, 1, u)
        return [...data]
      }
    })
  }

  return {
    subscribe,
    set,
    addSingle,
    remove
  }
}

export const users = createUsers()

