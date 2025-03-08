import React from 'react'
import { useProfiles, useUserProfile } from '../../hooks/useProfiles'
import { useParams } from 'react-router';
import { Profile } from '../../components/Profile';
import { useUserPosts } from '../../hooks/usePosts';

export const ProfilePage = () => {
    const {id} = useParams();
    const { profile, error, loading } = useUserProfile(Number.parseInt(id ?? ''))
    
    
    if(error) return <p>Errror {error}</p>
    if(loading || !profile) return <p>Loading..</p>
    
    if(!profile) return <p>No user with id {id}</p>

    return <Profile profile={profile} />
}