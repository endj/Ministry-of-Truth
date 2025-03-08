import React, { useState } from 'react'
import { usePosting } from '../../hooks/useReply';
import { useUser } from '../../hooks/useUser';
import style from "./threadcreator.module.css"

export const ThreadCreator = () => {
    const { createThread } = usePosting();
    const [input, setInput] = useState("")
    const { user } = useUser()

    const handleKeyDown = (event: { key: string }) => {
        if (event.key === 'Enter') {
            if (input.trim() && user) {
                createThread(user, input, res => console.log("Created Thread"))
                setInput('');
            }
        }
    };

    if(!user) return null;

    return (
        <div className={style.postbox}>
            <label>{user?.profile.name}</label>
            <input
                value={input}
                onKeyDown={handleKeyDown}
                onChange={e => setInput(e.target.value)}
                disabled={!Boolean(user)}
                autoFocus
                type='text'
                placeholder='What is happening'></input>
        </div>
    )
}