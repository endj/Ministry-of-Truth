import { useThreadPosts } from '../../hooks/usePosts';
import { useLocation, useParams } from 'react-router';
import { LinkMode, Post } from '../../components/Post/Post';
import { useUser } from '../../hooks/useUser';


export const Thread = () => {
    const { id } = useParams();
    const { hash } = useLocation();
    const { user } = useUser()


    const { posts, error, loading } = useThreadPosts(id ?? '')

    const linkedPost = hash?.replace("#", "")

    return (
        <>
            {error && <p>Error...</p>}
            {loading && <p>Loading...</p>}
            <h3>Thread: {id}</h3>
            {posts?.sort((p1, p2) => p1.createdAt - p2.createdAt)
                .map(post => <Post
                    key={post.id}
                    post={post}
                    link={LinkMode.NONE}
                    reply={post.id === Number.parseInt(linkedPost) && Boolean(user?.id)}
                />)}
        </>
    )
}