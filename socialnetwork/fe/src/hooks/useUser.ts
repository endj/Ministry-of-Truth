import React, {useContext} from 'react'
import { UserContext, UserContextType } from '../context/AuthContext'



export const useUser = () => {
    const {user, setUser} = useContext(UserContext) as UserContextType
    return {
        user, setUser
    }

}