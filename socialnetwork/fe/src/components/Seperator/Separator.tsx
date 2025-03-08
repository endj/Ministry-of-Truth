import React from 'react'
import style from "./seperator.module.css"

export const Separator = ({styles = {}}: {styles?: React.CSSProperties}) => {

    return (
        <div style={styles} className={style.Separator}></div>
    )
}