import { operationStore } from '@urql/svelte';
import {
	CreateSessionDocument,
	UpdateRadiusDocument,
	UpdateCoordsDocument,
	DelUsersDocument,
	SubUsersDocument,
	EndSessionDocument,
	GetUsersDocument,
	SubMessagesDocument,
	SendMessageDocument
} from '$lib/types/api';

import type {
	CreateSessionMutation,
	CreateSessionMutationVariables,
	UpdateRadiusMutation,
	UpdateRadiusMutationVariables,
	UpdateCoordsMutation,
	UpdateCoordsMutationVariables,
	DelUsersSubscription,
	DelUsersSubscriptionVariables,
	SubUsersSubscription,
	SubUsersSubscriptionVariables,
	EndSessionMutation,
	EndSessionMutationVariables,
	GetUsersQuery,
	GetUsersQueryVariables,
	SubMessagesSubscription,
	SubMessagesSubscriptionVariables,
	SendMessageMutation,
	SendMessageMutationVariables
} from '$lib/types/api';

export const sessionStore = operationStore<CreateSessionMutation, CreateSessionMutationVariables>(
	CreateSessionDocument
);

export const radiusStore = operationStore<UpdateRadiusMutation, UpdateRadiusMutationVariables>(
	UpdateRadiusDocument
);

export const updateCoordsStore = operationStore<
	UpdateCoordsMutation,
	UpdateCoordsMutationVariables
>(UpdateCoordsDocument);

export const usersUpdatesStore = operationStore<
	SubUsersSubscription,
	SubUsersSubscriptionVariables
>(SubUsersDocument);

export const delUsersStore = operationStore<DelUsersSubscription, DelUsersSubscriptionVariables>(
	DelUsersDocument
);

export const endSessionStore = operationStore<EndSessionMutation, EndSessionMutationVariables>(
	EndSessionDocument
);

export const getUsersStore = operationStore<GetUsersQuery, GetUsersQueryVariables>(
	GetUsersDocument
);

export const messagesStore = operationStore<
	SubMessagesSubscription,
	SubMessagesSubscriptionVariables
>(SubMessagesDocument);

export const sendMessageStore = operationStore<
	SendMessageMutation,
	SendMessageMutationVariables
>(SendMessageDocument)
