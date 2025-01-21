import { RemoveLocalToken } from "../services/auth-service"

function Logout({onAuthChange}: {onAuthChange: () => void}) {
  const handleLogout = () => {
    // todo: end session on server
    RemoveLocalToken()
    onAuthChange()
  }

  return (
    <div>
      <button className="btn btn-primary" onClick={handleLogout}>Logout</button>
    </div>
  )
}

export default Logout
