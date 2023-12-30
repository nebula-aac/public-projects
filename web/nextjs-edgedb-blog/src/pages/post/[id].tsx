import React from "react";
import { GetServerSidePropsContext, InferGetServerSidePropsType } from "next";
import e from '../../../dbschema/edgeql-js';
import { client } from "../api/post";

export const getServerSideProps = async (
    context?: GetServerSidePropsContext
) => {
    const post = await e
        .select(e.BlogPost, (post) => ({
            id: true,
            title: true,
            content: true,
            filter_single: e.op(
                post.id,
                '=',
                e.uuid(context!.params!.id as string)
            ),
        }))
        .run(client);
    return { props: { post: post! } };
};

export type GetPost = InferGetServerSidePropsType<typeof getServerSideProps>;

const Post = (props: GetPost) => {
    return (
        <div
            style={{
                margin: 'auto',
                width: '100%',
                maxWidth: '600px',
            }}
        >
            <h1 style={{ padding: '50px 0px' }}>{props.post.title}</h1>
            <p style={{ color: '#666' }}>{props.post.content}</p>
        </div>
    );
}

export default Post;