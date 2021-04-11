import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};


export type Coords = {
  __typename?: 'Coords';
  x: Scalars['Float'];
  y: Scalars['Float'];
};

export type CoordsInput = {
  x: Scalars['Float'];
  y: Scalars['Float'];
};

export type Message = {
  __typename?: 'Message';
  content: Scalars['String'];
  time: Scalars['Int'];
  user: User;
};

export type Mutation = {
  __typename?: 'Mutation';
  createSession: Session;
  endSession: Scalars['Boolean'];
  updateRadius: User;
  updateName: User;
  updateCoords: User;
  sendMessage: Scalars['Boolean'];
};


export type MutationCreateSessionArgs = {
  input: SessionInput;
};


export type MutationUpdateRadiusArgs = {
  radius: Scalars['Int'];
};


export type MutationUpdateNameArgs = {
  name: Scalars['String'];
};


export type MutationUpdateCoordsArgs = {
  input: CoordsInput;
};


export type MutationSendMessageArgs = {
  content: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  users: Array<Maybe<User>>;
};

export type Session = {
  __typename?: 'Session';
  token: Scalars['String'];
  user: User;
};

export type SessionInput = {
  name: Scalars['String'];
  x: Scalars['Float'];
  y: Scalars['Float'];
};

export type Subscription = {
  __typename?: 'Subscription';
  users: User;
  delUsers: Scalars['ID'];
  messages: Message;
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID'];
  name: Scalars['String'];
  radius: Scalars['Int'];
  coords?: Maybe<Coords>;
};

export type CreateSessionMutationVariables = Exact<{
  name: Scalars['String'];
  x: Scalars['Float'];
  y: Scalars['Float'];
}>;


export type CreateSessionMutation = (
  { __typename?: 'Mutation' }
  & { createSession: (
    { __typename?: 'Session' }
    & Pick<Session, 'token'>
    & { user: (
      { __typename?: 'User' }
      & Pick<User, 'id' | 'name' | 'radius'>
      & { coords?: Maybe<(
        { __typename?: 'Coords' }
        & Pick<Coords, 'x' | 'y'>
      )> }
    ) }
  ) }
);

export type DelUsersSubscriptionVariables = Exact<{ [key: string]: never; }>;


export type DelUsersSubscription = (
  { __typename?: 'Subscription' }
  & Pick<Subscription, 'delUsers'>
);

export type EndSessionMutationVariables = Exact<{ [key: string]: never; }>;


export type EndSessionMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'endSession'>
);

export type GetUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUsersQuery = (
  { __typename?: 'Query' }
  & { users: Array<Maybe<(
    { __typename?: 'User' }
    & Pick<User, 'id' | 'name' | 'radius'>
    & { coords?: Maybe<(
      { __typename?: 'Coords' }
      & Pick<Coords, 'x' | 'y'>
    )> }
  )>> }
);

export type SendMessageMutationVariables = Exact<{
  content: Scalars['String'];
}>;


export type SendMessageMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'sendMessage'>
);

export type SubMessagesSubscriptionVariables = Exact<{ [key: string]: never; }>;


export type SubMessagesSubscription = (
  { __typename?: 'Subscription' }
  & { messages: (
    { __typename?: 'Message' }
    & Pick<Message, 'content'>
    & { user: (
      { __typename?: 'User' }
      & Pick<User, 'name'>
    ) }
  ) }
);

export type SubUsersSubscriptionVariables = Exact<{ [key: string]: never; }>;


export type SubUsersSubscription = (
  { __typename?: 'Subscription' }
  & { users: (
    { __typename?: 'User' }
    & Pick<User, 'id' | 'name' | 'radius'>
    & { coords?: Maybe<(
      { __typename?: 'Coords' }
      & Pick<Coords, 'x' | 'y'>
    )> }
  ) }
);

export type UpdateCoordsMutationVariables = Exact<{
  x: Scalars['Float'];
  y: Scalars['Float'];
}>;


export type UpdateCoordsMutation = (
  { __typename?: 'Mutation' }
  & { updateCoords: (
    { __typename?: 'User' }
    & Pick<User, 'id' | 'radius'>
    & { coords?: Maybe<(
      { __typename?: 'Coords' }
      & Pick<Coords, 'x' | 'y'>
    )> }
  ) }
);

export type UpdateRadiusMutationVariables = Exact<{
  rad: Scalars['Int'];
}>;


export type UpdateRadiusMutation = (
  { __typename?: 'Mutation' }
  & { updateRadius: (
    { __typename?: 'User' }
    & Pick<User, 'id' | 'radius'>
    & { coords?: Maybe<(
      { __typename?: 'Coords' }
      & Pick<Coords, 'x' | 'y'>
    )> }
  ) }
);


export const CreateSessionDocument = gql`
    mutation createSession($name: String!, $x: Float!, $y: Float!) {
  createSession(input: {name: $name, x: $x, y: $y}) {
    token
    user {
      id
      name
      radius
      coords {
        x
        y
      }
    }
  }
}
    `;
export const DelUsersDocument = gql`
    subscription delUsers {
  delUsers
}
    `;
export const EndSessionDocument = gql`
    mutation endSession {
  endSession
}
    `;
export const GetUsersDocument = gql`
    query getUsers {
  users {
    id
    name
    radius
    coords {
      x
      y
    }
  }
}
    `;
export const SendMessageDocument = gql`
    mutation sendMessage($content: String!) {
  sendMessage(content: $content)
}
    `;
export const SubMessagesDocument = gql`
    subscription subMessages {
  messages {
    content
    user {
      name
    }
  }
}
    `;
export const SubUsersDocument = gql`
    subscription subUsers {
  users {
    id
    name
    radius
    coords {
      x
      y
    }
  }
}
    `;
export const UpdateCoordsDocument = gql`
    mutation UpdateCoords($x: Float!, $y: Float!) {
  updateCoords(input: {x: $x, y: $y}) {
    id
    radius
    coords {
      x
      y
    }
  }
}
    `;
export const UpdateRadiusDocument = gql`
    mutation UpdateRadius($rad: Int!) {
  updateRadius(radius: $rad) {
    id
    radius
    coords {
      x
      y
    }
  }
}
    `;