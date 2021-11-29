import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
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
  Time: any;
};

export type AddTag = {
  post_id: Scalars['ID'];
  tag_id: Scalars['ID'];
};

export type DeletePost = {
  id: Scalars['ID'];
};

export type Mutation = {
  __typename?: 'Mutation';
  addTag: Post;
  createPost: Post;
  deletePost: Scalars['ID'];
};


export type MutationAddTagArgs = {
  input?: InputMaybe<AddTag>;
};


export type MutationCreatePostArgs = {
  input?: InputMaybe<NewPost>;
};


export type MutationDeletePostArgs = {
  input?: InputMaybe<DeletePost>;
};

export type NewPost = {
  title: Scalars['String'];
};

export type Post = {
  __typename?: 'Post';
  createdAt: Scalars['Time'];
  id: Scalars['ID'];
  title: Scalars['String'];
  updatedAt: Scalars['Time'];
  user: User;
};

export type Query = {
  __typename?: 'Query';
  post: Post;
  posts: Array<Post>;
  tagPosts: Array<Post>;
  tags: Array<Tag>;
  users: Array<Maybe<User>>;
};


export type QueryPostArgs = {
  id: Scalars['ID'];
};


export type QueryTagPostsArgs = {
  tag_id: Scalars['ID'];
};


export type QueryTagsArgs = {
  input?: InputMaybe<Tags>;
};

export type Tag = {
  __typename?: 'Tag';
  createdAt: Scalars['Time'];
  id: Scalars['ID'];
  name: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export type Tags = {
  post_id: Scalars['ID'];
};

export type User = {
  __typename?: 'User';
  createdAt: Scalars['Time'];
  email: Scalars['String'];
  id: Scalars['ID'];
  name: Scalars['String'];
  picture: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export type CreatePostMutationVariables = Exact<{
  title: Scalars['String'];
}>;


export type CreatePostMutation = { __typename?: 'Mutation', createPost: { __typename?: 'Post', id: string, title: string, createdAt: any, updatedAt: any } };


export const CreatePostDocument = gql`
    mutation createPost($title: String!) {
  createPost(input: {title: $title}) {
    id
    title
    createdAt
    updatedAt
  }
}
    `;
export type CreatePostMutationFn = Apollo.MutationFunction<CreatePostMutation, CreatePostMutationVariables>;

/**
 * __useCreatePostMutation__
 *
 * To run a mutation, you first call `useCreatePostMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreatePostMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createPostMutation, { data, loading, error }] = useCreatePostMutation({
 *   variables: {
 *      title: // value for 'title'
 *   },
 * });
 */
export function useCreatePostMutation(baseOptions?: Apollo.MutationHookOptions<CreatePostMutation, CreatePostMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreatePostMutation, CreatePostMutationVariables>(CreatePostDocument, options);
      }
export type CreatePostMutationHookResult = ReturnType<typeof useCreatePostMutation>;
export type CreatePostMutationResult = Apollo.MutationResult<CreatePostMutation>;
export type CreatePostMutationOptions = Apollo.BaseMutationOptions<CreatePostMutation, CreatePostMutationVariables>;