import React from 'react'
import { IProfile } from '../hooks/useProfiles'
import { NavLink } from 'react-router'
import { Avatar } from './Avatar/Avatar'
import { LinkMode, Post } from './Post/Post'
import { useUserPosts } from '../hooks/usePosts'

export const Profile = ({ profile }: { profile: IProfile }) => {
    const { posts, loading, error } = useUserPosts(profile.id)

    return (
        <div>
            <div style={{ margin: "20px" }}>
                <Avatar side={125} author={profile.profile.name} />
                <p><b>{profile.profile.name}</b></p>
                <p>{profile.profile.info}</p>
            </div>
            {loading && <p>Loading...</p>}
            {error && <p>Error...</p>}
            {posts?.map(post => <Post key={post.id} post={post} link={LinkMode.LINK_THREAD} />)}

        </div>

    )
}