import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions =  {}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Message = {
  __typename?: 'Message';
  id: Scalars['ID'];
  isSender: Scalars['Boolean'];
  roomId: Scalars['ID'];
  text: Scalars['String'];
  userID: Scalars['ID'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createRoom?: Maybe<Room>;
  postMessage?: Maybe<Message>;
};


export type MutationCreateRoomArgs = {
  description?: Maybe<Scalars['String']>;
  name: Scalars['String'];
};


export type MutationPostMessageArgs = {
  roomId: Scalars['ID'];
  text: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  messages: Array<Message>;
  rooms: Array<Room>;
};


export type QueryMessagesArgs = {
  roomId: Scalars['String'];
};

export type Room = {
  __typename?: 'Room';
  description?: Maybe<Scalars['String']>;
  disabled: Scalars['Boolean'];
  id: Scalars['ID'];
  name: Scalars['String'];
  own_messages: Array<Message>;
};

export type Subscription = {
  __typename?: 'Subscription';
  messagePosted: Message;
};


export type SubscriptionMessagePostedArgs = {
  roomId: Scalars['ID'];
};

export type PostMessageMutationVariables = Exact<{
  roomId: Scalars['ID'];
  text: Scalars['String'];
}>;


export type PostMessageMutation = { __typename?: 'Mutation', postMessage?: { __typename?: 'Message', id: string, text: string, roomId: string } | null | undefined };

export type CreateRoomMutationVariables = Exact<{
  name: Scalars['String'];
  description?: Maybe<Scalars['String']>;
}>;


export type CreateRoomMutation = { __typename?: 'Mutation', createRoom?: { __typename?: 'Room', id: string, name: string, disabled: boolean, description?: string | null | undefined } | null | undefined };

export type RoomsQueryVariables = Exact<{ [key: string]: never; }>;


export type RoomsQuery = { __typename?: 'Query', rooms: Array<{ __typename?: 'Room', id: string, name: string, disabled: boolean, description?: string | null | undefined, own_messages: Array<{ __typename: 'Message' }> }> };

export type MessagesQueryVariables = Exact<{
  roomId: Scalars['String'];
}>;


export type MessagesQuery = { __typename?: 'Query', messages: Array<{ __typename?: 'Message', id: string, text: string, roomId: string, userID: string, isSender: boolean }> };

export type MessagePostedSubscriptionVariables = Exact<{
  roomId: Scalars['ID'];
}>;


export type MessagePostedSubscription = { __typename?: 'Subscription', messagePosted: { __typename?: 'Message', id: string, roomId: string, text: string } };


export const PostMessageDocument = gql`
    mutation postMessage($roomId: ID!, $text: String!) {
  postMessage(roomId: $roomId, text: $text) {
    id
    text
    roomId
  }
}
    `;
export type PostMessageMutationFn = Apollo.MutationFunction<PostMessageMutation, PostMessageMutationVariables>;

/**
 * __usePostMessageMutation__
 *
 * To run a mutation, you first call `usePostMessageMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `usePostMessageMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [postMessageMutation, { data, loading, error }] = usePostMessageMutation({
 *   variables: {
 *      roomId: // value for 'roomId'
 *      text: // value for 'text'
 *   },
 * });
 */
export function usePostMessageMutation(baseOptions?: Apollo.MutationHookOptions<PostMessageMutation, PostMessageMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<PostMessageMutation, PostMessageMutationVariables>(PostMessageDocument, options);
      }
export type PostMessageMutationHookResult = ReturnType<typeof usePostMessageMutation>;
export type PostMessageMutationResult = Apollo.MutationResult<PostMessageMutation>;
export type PostMessageMutationOptions = Apollo.BaseMutationOptions<PostMessageMutation, PostMessageMutationVariables>;
export const CreateRoomDocument = gql`
    mutation createRoom($name: String!, $description: String) {
  createRoom(name: $name, description: $description) {
    id
    name
    disabled
    description
  }
}
    `;
export type CreateRoomMutationFn = Apollo.MutationFunction<CreateRoomMutation, CreateRoomMutationVariables>;

/**
 * __useCreateRoomMutation__
 *
 * To run a mutation, you first call `useCreateRoomMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateRoomMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createRoomMutation, { data, loading, error }] = useCreateRoomMutation({
 *   variables: {
 *      name: // value for 'name'
 *      description: // value for 'description'
 *   },
 * });
 */
export function useCreateRoomMutation(baseOptions?: Apollo.MutationHookOptions<CreateRoomMutation, CreateRoomMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateRoomMutation, CreateRoomMutationVariables>(CreateRoomDocument, options);
      }
export type CreateRoomMutationHookResult = ReturnType<typeof useCreateRoomMutation>;
export type CreateRoomMutationResult = Apollo.MutationResult<CreateRoomMutation>;
export type CreateRoomMutationOptions = Apollo.BaseMutationOptions<CreateRoomMutation, CreateRoomMutationVariables>;
export const RoomsDocument = gql`
    query rooms {
  rooms {
    id
    name
    disabled
    description
    own_messages {
      __typename
    }
  }
}
    `;

/**
 * __useRoomsQuery__
 *
 * To run a query within a React component, call `useRoomsQuery` and pass it any options that fit your needs.
 * When your component renders, `useRoomsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useRoomsQuery({
 *   variables: {
 *   },
 * });
 */
export function useRoomsQuery(baseOptions?: Apollo.QueryHookOptions<RoomsQuery, RoomsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<RoomsQuery, RoomsQueryVariables>(RoomsDocument, options);
      }
export function useRoomsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<RoomsQuery, RoomsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<RoomsQuery, RoomsQueryVariables>(RoomsDocument, options);
        }
export type RoomsQueryHookResult = ReturnType<typeof useRoomsQuery>;
export type RoomsLazyQueryHookResult = ReturnType<typeof useRoomsLazyQuery>;
export type RoomsQueryResult = Apollo.QueryResult<RoomsQuery, RoomsQueryVariables>;
export const MessagesDocument = gql`
    query messages($roomId: String!) {
  messages(roomId: $roomId) {
    id
    text
    roomId
    userID
    isSender
  }
}
    `;

/**
 * __useMessagesQuery__
 *
 * To run a query within a React component, call `useMessagesQuery` and pass it any options that fit your needs.
 * When your component renders, `useMessagesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMessagesQuery({
 *   variables: {
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useMessagesQuery(baseOptions: Apollo.QueryHookOptions<MessagesQuery, MessagesQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<MessagesQuery, MessagesQueryVariables>(MessagesDocument, options);
      }
export function useMessagesLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<MessagesQuery, MessagesQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<MessagesQuery, MessagesQueryVariables>(MessagesDocument, options);
        }
export type MessagesQueryHookResult = ReturnType<typeof useMessagesQuery>;
export type MessagesLazyQueryHookResult = ReturnType<typeof useMessagesLazyQuery>;
export type MessagesQueryResult = Apollo.QueryResult<MessagesQuery, MessagesQueryVariables>;
export const MessagePostedDocument = gql`
    subscription messagePosted($roomId: ID!) {
  messagePosted(roomId: $roomId) {
    id
    roomId
    text
  }
}
    `;

/**
 * __useMessagePostedSubscription__
 *
 * To run a query within a React component, call `useMessagePostedSubscription` and pass it any options that fit your needs.
 * When your component renders, `useMessagePostedSubscription` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMessagePostedSubscription({
 *   variables: {
 *      roomId: // value for 'roomId'
 *   },
 * });
 */
export function useMessagePostedSubscription(baseOptions: Apollo.SubscriptionHookOptions<MessagePostedSubscription, MessagePostedSubscriptionVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useSubscription<MessagePostedSubscription, MessagePostedSubscriptionVariables>(MessagePostedDocument, options);
      }
export type MessagePostedSubscriptionHookResult = ReturnType<typeof useMessagePostedSubscription>;
export type MessagePostedSubscriptionResult = Apollo.SubscriptionResult<MessagePostedSubscription>;