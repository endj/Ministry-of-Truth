import React, { createContext, useState } from 'react'
import { IProfile } from '../hooks/useProfiles'



export type UserContextType = {
    user: IProfile | null,
    setUser: (profile: IProfile) => void
}
export const UserContext = createContext<UserContextType | undefined>(undefined)


export const UserProvider = ({ children }: { children: React.ReactNode }) => {
    const [user, setUser] = useState<IProfile | null>(null)

    return (
        <UserContext.Provider value={{
            user: user,
            setUser: setUser
        }}>
            {children}
        </UserContext.Provider>
    )
}