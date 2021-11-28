export interface Post {
  id: number;
  title: string;
  createdAt: string;
  updatedAt: string;
}

export interface Posts {
  posts: Post[];
}
