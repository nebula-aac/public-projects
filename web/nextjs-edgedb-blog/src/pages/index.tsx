import Head from 'next/head'
import styles from '#/styles/Home.module.css'
import { NextPage } from 'next'
import { useEffect, useState } from 'react';
import { Posts } from './api/post';

/* Remove after inferring Posts in /api/post
type Post = {
  id: string;
  title: string;
  content: string;
}
*/

const HomePage: NextPage = () => {
  /*
  const posts: Post[] = [
    {
      id: 'post1',
      title: 'This one weird trick makes using databases fun',
      content: 'Use EdgeDB',
    },
    {
      id: 'post2',
      title: 'How to build a blog with EdgeDB and Next.js',
      content: "Let's start by scaffolding our app with `create-next-app`.",
    },
  ];
  */

  const [posts, setPosts] = useState<Posts[] | null>(null);

  useEffect(() => {
    fetch(`/api/post`)
      .then((result) => result.json())
      .then(setPosts)
  }, [])

  if (!posts) return <p>Loading...</p>

  return (
    <div className={styles.container}>
      <Head>
        <title>My Blog</title>
        <meta name="description" content="An awesome blog" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>Blog</h1>
        <div style={{ height: '50px' }}></div>
        {posts.map((post) => {
          return (
            <a href={`/post/${post?.id}`} key={post?.id}>
              <div className={styles.card}>
                <p>{post?.title}</p>
              </div>
            </a>
          );
        })}
      </main>
    </div>
  );
};

export default HomePage;