import { writable } from 'svelte/store';

export function localStore(key, value) {
  const data = typeof localStorage != 'undefined' ? localStorage.getItem(key) : null;
  const store = writable(value);
  if (data !== null) {
    store.set(JSON.parse(data));
  }
  store.subscribe((val) => {
    if (typeof localStorage == 'undefined') {
      return;
    }
    localStorage.setItem(key, JSON.stringify(val));
  });

  return store;
}
