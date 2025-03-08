import React from 'react'
import style from './avatar.module.css'

export const Avatar = ({ author, side = 45 }: { author: string, side?: number }) => {

    const short = author.split("\\s+").map(i => i[0]).map(c => c.toLocaleUpperCase()).join("")

    return (
        <div  style={{height: side, width: side}} className={style.avatar}>
            <span>
                {short}
            </span>
        </div>
    )
}