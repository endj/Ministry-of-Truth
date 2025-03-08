import React, {useMemo} from 'react'
import { useFetch  } from './useFetch'

export interface IProfile {
    id: number,
    traits: Record<string, any>,
    profile: Record<string, any>
}


export interface IRawProfile {
    id: number,
    traits: string,
    profile: string
}

const mapToProfile = (profile: IRawProfile) => {
    return {
        id: profile.id,
        traits: JSON.parse(profile.traits),
        profile: JSON.parse(profile.profile),
    }
}

export const useProfiles = () => {
   const {data, loading, error} = useFetch<IRawProfile[]>("http://localhost:8080/profiles")

    const transformed: IProfile[] | null = data ? data.map(mapToProfile) : data

    return {
        data: transformed,
        loading,
        error
    } 
}

export const useUserProfile = (userId: number) => {
    if(!userId) {
        return {profile: null, loading: false, errro: null}
    }

    const {data, loading, error} = useFetch<IRawProfile[]>("http://localhost:8080/profiles")


    const profile = useMemo(() => {
        const profiles = data ? data : []
        const userProfile = profiles.find(profile => profile.id === userId)
        return userProfile ? mapToProfile(userProfile) : null
    }, [userId, data])

    return {
        profile,
        loading,
        error
    }


}