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

export type Mutation = {
  __typename?: 'Mutation';
  createSession: Session;
  updateRadius: User;
  updateName: User;
  connect: Scalars['Boolean'];
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


export type MutationConnectArgs = {
  id: Scalars['ID'];
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
  newUsers: User;
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
