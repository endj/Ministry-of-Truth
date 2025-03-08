import React, { useMemo } from 'react'
import { NavLink } from 'react-router'
import style from "./header.module.css"
import { useUser } from '../../hooks/useUser'
import { IProfile, useProfiles } from '../../hooks/useProfiles'
import { Avatar } from '../Avatar/Avatar'


const LoggedInUser = ({ profile }: { profile: IProfile }) => {
    return (<NavLink to={`/profiles/${profile.id}`}><Avatar author={profile.profile.name}></Avatar></NavLink>)
}

const loggedInProfile = (data: IProfile[], userId: number) => {
    return data?.find(profile => profile.id === userId);
}

export const Header = () => {
    const { data, error, loading } = useProfiles()
    const { user, setUser } = useUser()

    if (loading) return <p>Loading...</p>
    if (error) return <p>Error.. {error}</p>


    const onSelect = (e: React.ChangeEvent<HTMLSelectElement>) => {
        if(e.target.value) {
            const profile = loggedInProfile(data ?? [], Number.parseInt(e.target.value))
            if(profile) setUser(profile)
        }
    }

    return (
        <div className={style.menu}>
            <div>
                <NavLink to={"/"}>Home</NavLink> | < NavLink to="/profiles">Profiles</NavLink>
            </div>
            <div style={{display: "flex", alignContent: 'center', gap: "10px" }}>
                <div style={{alignContent:"center", height: "45px"}}>
                    <label htmlFor='userSelector'>User:{' '}</label>
                    <select
                        id="userSelector"
                        className={style.userselector}
                        value={user?.id ?? ''}
                        onChange={onSelect}
                    >
                        <option value=""></option>
                        {data?.map(user => <option key={user.id} value={user.id}>{user.profile.name}</option>)}
                    </select>
                </div>
                {user && <LoggedInUser profile={user} />}
            </div>
        </div>

    )
}