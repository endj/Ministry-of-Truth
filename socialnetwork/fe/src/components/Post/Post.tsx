import { useState } from 'react'
import { IPost } from '../../hooks/usePosts'
import style from './post.module.css'
import { Separator } from '../Seperator/Separator'
import { Avatar } from '../Avatar/Avatar'
import { NavLink, useNavigate } from 'react-router'
import { useUser } from '../../hooks/useUser'
import { usePosting } from '../../hooks/useReply'

export enum LinkMode {
    LINK_PROFILE,
    LINK_FEED,
    LINK_THREAD,
    NONE
}

const getLink = (post: IPost, link: LinkMode) => {
    switch (link) {
        case LinkMode.LINK_PROFILE:
            return `/profiles/${post.authorId}#${post.id}`
        case LinkMode.LINK_FEED:
            return `/#${post.id}`
        case LinkMode.LINK_THREAD:
            return `/feed/${post.threadId}#${post.id}`
        case LinkMode.NONE:
            return ""
    }
}

export const Reply = ({ post }: { post: IPost }) => {
    const { reply } = usePosting();
    const navigate = useNavigate()
    const [input, setInput] = useState("")
    const { user } = useUser()

    const handleKeyDown = (event: { key: string }) => {
        if (event.key === 'Enter') {
            if (input.trim() && user) {
                reply(user, input, post)
                navigate(0)
                setInput('');
            }
        }
    };

    return (
        <div className={style.reply}>
            <label>{user?.profile.name}</label>
            <input
                value={input}
                onKeyDown={handleKeyDown}
                onChange={e => setInput(e.target.value)}
                disabled={!Boolean(user)}
                autoFocus
                type='text'
                placeholder='Post your reply'></input>
        </div>
    )
}

export const Post = ({ post, link = LinkMode.NONE, reply = false }: { post: IPost, link?: LinkMode, reply?: boolean }) => {


    const postLink = getLink(post, link)

    const body = link == LinkMode.NONE
        ? (
            <div className={style.postContent}>
                <p>{post.content}</p>
            </div>
        )
        : (
            <NavLink to={postLink}>
                <div className={style.postContent}>
                    <p>{post.content}</p>
                </div>
            </NavLink>
        )

    return (
        <div id={String(post.id)} className={style.post}>
            <NavLink to={`/profiles/${post.authorId}`}>
                <div className={style.postHeader}>
                    <Avatar author={post.author} />
                    <p>{post.author}</p>
                </div>
            </NavLink>

            {body}

            <Separator />
            <div className={style.postFooter}>
                <div style={{ display: "flex", justifyContent: "space-between" }}>
                    <p>{new Date(post.createdAt).toUTCString()}</p>
                    <p>{post.threadId}</p>
                </div>
                {reply && <Reply post={post} />
                }
            </div>

        </div>
    )
}