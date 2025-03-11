import './App.css'
import { Profiles } from './pages/Profiles/Profiles'
import { Feed } from './pages/Feed/Feed'
import { Route, Routes } from 'react-router'
import { ProfilePage } from './pages/Profile/ProfilePage'
import { Thread } from './pages/Thread/Thread'
import { Header } from './components/Header/Header'



function App() {

  return (
    <>
      <Header />
      <div style={{margin: "auto", width: "70%"}}>
        <Routes>
          <Route index element={<Feed />}></Route>
          <Route path="feed/:id" element={<Thread />} />
          <Route path="profiles">
            <Route index element={<Profiles />} />
            <Route path=':id' element={<ProfilePage />} />
          </Route>
          <Route path="*" element={<p>404</p>} />
        </Routes>
      </div>
    </>

  )
}

export default App
