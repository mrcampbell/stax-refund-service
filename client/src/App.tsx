import { useEffect, useState } from 'react';
import './App.css'
import Login from './components/Login';
import Logout from './components/Logout';
import RefundPage from './pages/RefundPage';
import { GetLocalToken } from './services/auth-service'
import DocsPage from './pages/DocsPage';


function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  useEffect(() => {
    const token = GetLocalToken()
    setIsLoggedIn(!!token)
  }, [])

  const onAuthChange = () => {
    console.log('Auth changed')
    setIsLoggedIn(!!GetLocalToken())
  }

  return (
    <>
      <div className="flex flex-row justify-between p-8">
        <h1 className="text-3xl font-bold text-green-500">
          Mike's Refund Service™
        </h1>
        {isLoggedIn && <Logout onAuthChange={onAuthChange} />}
      </div>

      {isLoggedIn ? <RefundPage /> : <Login onAuthChange={onAuthChange} />}
    </>
  )
}

export default App
