import styles from './profiles.module.css'
import { useProfiles } from '../../hooks/useProfiles'
import { NavLink } from 'react-router'
import { Separator } from '../../components/Seperator/Separator'

export const Profiles = () => {
    const { data, error, loading } = useProfiles()

    return (
        <>
            {error && <p>Error: {error}</p>}
            {loading && <p>Loading...</p>}
            {data &&
                <ul className={styles.userList}>
                    {
                        data.map(profile =>
                            <div key={profile.id}>
                                <NavLink to={"/profiles/" + profile.id}><li>{profile.profile.name}</li></NavLink>
                                <Separator />
                            </div>
                        )
                    }
                </ul>
            }
        </>
    )
}
