// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import {createClient} from 'edgedb';
import e, {$infer} from '../../../dbschema/edgeql-js';

export const client = createClient();

const selectedPosts = e.select(e.BlogPost, () => ({
  id: true,
  title: true,
  content: true,
}))

export type Posts = $infer<typeof selectedPosts>;

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  /*
  const posts = await client.query(`select BlogPost {
    id,
    title,
    content
  };`);
  */

  const posts = await selectedPosts.run(client);

  res.status(200).json(posts);
}
