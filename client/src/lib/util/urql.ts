import { operationStore } from '@urql/svelte'
import { CreateSessionDocument } from '$lib/types/api'

export const sessionStore = operationStore(CreateSessionDocument)
