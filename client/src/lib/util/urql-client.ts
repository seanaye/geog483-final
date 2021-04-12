import { browser } from '$app/env';
import {
	makeOperation,
	createClient,
	cacheExchange,
	dedupExchange,
	fetchExchange,
	subscriptionExchange
} from '@urql/svelte';
import { authExchange } from '@urql/exchange-auth';
import { sessionStore } from './urql';
import ws from 'ws';
import { SubscriptionClient } from 'subscriptions-transport-ws';

const uri = 'geomsg.seanaye.ca/query';

const makeHeader = () => {
	const token = sessionStore.data?.createSession?.token;
	return {
		headers: { Authorization: token ? `Bearer ${token}` : '' }
	};
};

const makeSubClient = () =>
	new SubscriptionClient(
		`wss://${uri}`,
		{
			reconnect: true,
			connectionParams: makeHeader
		},
		browser ? undefined : ws
	);

let subClient;

export const client = createClient({
	url: `https://${uri}`,
	fetchOptions: makeHeader,
	exchanges: [
		dedupExchange,
		cacheExchange,
		authExchange({
			getAuth: ({ authState }) => {
				return new Promise((resolve) => {
					if (!authState) {
						const token: string | undefined = sessionStore.data?.createSession?.token;
						resolve(token ? { token } : null);
					}
					resolve(null);
				});
			},
			addAuthToOperation: ({ authState, operation }) => {
				if (!authState) {
					return operation;
				}

				const fetchOptions =
					typeof operation.context.fetchOptions === 'function'
						? operation.context.fetchOptions()
						: operation.context.fetchOptions || {};

				let auth: { token: string } = authState as { token: string };

				return makeOperation(operation.kind, operation, {
					...operation.context,
					fetchOptions: {
						...fetchOptions,
						headers: {
							...fetchOptions.headers,
							Authorization: `Bearer ${auth.token}`
						}
					}
				});
			},
			didAuthError: ({ error }) => {
				console.log({ error });
				return !!error.graphQLErrors;
			}
		}),
		fetchExchange,
		subscriptionExchange({
			forwardSubscription(operation) {
				if (!subClient) {
					subClient = makeSubClient();
				}
				return subClient.request(operation);
			}
		})
	]
});
