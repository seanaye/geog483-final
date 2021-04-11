import { browser } from '$app/env';
import { createClient, defaultExchanges, subscriptionExchange } from '@urql/svelte';
import { sessionStore } from './urql';
import ws from 'ws';
import { SubscriptionClient } from 'subscriptions-transport-ws';

const uri = 'localhost:8080/query';

const subscriptionClient = new SubscriptionClient(
	`ws://${uri}`,
	{ reconnect: true },
	browser ? undefined : ws
);

export const client = createClient({
	url: `http://${uri}`,
	fetchOptions: () => {
		const token = sessionStore.data?.createSession?.token;
		return {
			headers: { authorization: token ? `Bearer ${token}` : '' }
		};
	},
	exchanges: [
		...defaultExchanges,
		subscriptionExchange({
			forwardSubscription(operation) {
				return subscriptionClient.request(operation);
			}
		})
	]
});
