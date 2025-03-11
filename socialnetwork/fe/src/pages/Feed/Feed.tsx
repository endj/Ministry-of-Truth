import { useEffect } from 'react'
import { useLocation } from 'react-router'
import { ThreadCreator } from '../../components/CreateThread/ThreadCreator'
import { LinkMode, Post } from '../../components/Post/Post'
import { Separator } from '../../components/Seperator/Separator'
import { useThreads } from '../../hooks/usePosts'

export const Feed = () => {
    const { threads, loading, error } = useThreads()
    const { hash } = useLocation();

    useEffect(() => {
        if (hash && threads) {
            const id = hash.replace("#", "");
            const element = document.getElementById(id);
            if (element) {
                element.scrollIntoView({ behavior: "smooth" });
            }
        }
    }, [hash, threads]);

    const noPosts = !loading && threads?.length == 0

    return (
        <>
            {error && <p>Error: {error}</p>}
            {loading && <p>Loading...</p>}
            <h3>Feed</h3>
            <ThreadCreator />
            {noPosts && <p>No Posts yet..</p>}
            {threads?.map(post => {
                return <div key={post.id} style={{ width: "100%" }}>
                    <Post key={post.id} post={post} link={LinkMode.LINK_THREAD} />
                    <Separator styles={{ margin: "10px" }} />
                </div>
            })}
        </>
    )

}